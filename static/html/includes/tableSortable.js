document.addEventListener('click', function(e){
	var element = e.target;
	var cellIndex = 0;
	if(element.dataset.nonsortable !== 'true'){
		if(element.parentNode.parentNode.nodeName === "THEAD"){
			if(element.parentNode.parentNode.parentNode.classList.contains("data-table")){
			//if(element.parentNode.parentNode.parentNode.classList.contains("sortable")){
				[].forEach.call(element.parentNode.cells, function(cell, index){
					if(cell !== element){
						cell.classList.remove('order-field', 'order-asc', 'order-desc');
						delete cell.dataset.order;
					}else{
						cellIndex = index;
					}
				});
				element.classList.add('order-field');
				var order = (element.dataset.order === undefined) ? 'desc' : element.dataset.order;
				element.classList.remove('order-' + order);
				order = (order === 'asc') ? 'desc' : 'asc';
				element.classList.add('order-' + order);
				[].slice.call(element.parentNode.parentNode.parentNode.tBodies[0].rows).sort(function(a, b){
					var isNum = element.dataset.isnum;
					var aVal = (order === 'asc') ? getValue(a.cells[cellIndex], isNum) : getValue(b.cells[cellIndex], isNum);
					var bVal = (order === 'asc') ? getValue(b.cells[cellIndex], isNum) : getValue(a.cells[cellIndex], isNum);
					return aVal < bVal ? -1 : aVal > bVal ? 1 : 0;
				}).forEach(function(row, index){
					element.parentNode.parentNode.parentNode.tBodies[0].appendChild(row);
				});
				element.dataset.order = order;
			}
		}
	}

	function getValue(cell, isNum){
		return (cell.dataset.value !== undefined) ? cell.dataset.value : (isNum === 'true') ? parseInt(cell.textContent) : cell.textContent;
	}
});