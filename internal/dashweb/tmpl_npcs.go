package dashweb

const (
	TmplNpcsByZoneTop = `
		<table class="data-table">
			<thead>
				<tr>
					<th>Expansion</th>
					<th>Zone</th>
					<th>NPCs By Zone</th>
				</tr>
			</thead>
			<tbody>{{ range .}}
				<tr>
					<td>{{ .Expansion }}</td>
					<td><a href="/npcs.html?npcsbyzone={{ .Short_name }}">{{ .Name }}</a></td>
				</tr>{{ end }}
			</tbody>
		</table>
		<script type="text/javascript" src="/includes/tableSortable.js"></script>
	`
)

const (
	TmplNpcsByZone = `
		<table class="data-table">
			<thead>
				<tr>
					<th>ID</th>
					<th>Name</th>
					<th>Level</th>
					<th>Max Level</th>
					<th>Race</th>
					<th>NPCs By Zone</th>
				</tr>
			</thead>
			<tbody>{{ range .}}
				<tr>
					<td>{{ .Id }}</td>
					<td>{{ .Name }}</td>
					<td>{{ .Level }}</td>
					<td>{{ .MaxLevel }}</td>
					<td>{{ .Race }}</td>
				</tr>{{ end }}
			</tbody>
		</table>
		<script type="text/javascript" src="/includes/tableSortable.js"></script>
	`
)
