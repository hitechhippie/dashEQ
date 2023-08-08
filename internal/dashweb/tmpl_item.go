package dashweb

const (
	Items = `
		<table class="data-table">
			<thead>
				<tr>
					<th>Items</th>
				</tr>
					<tr>
						<th>ID</th>
						<th>Name</th>
						<th>Item Type</th>
						<th>Magic?</th>
						<th>NODROP?</th>
						<th>NORENT?</th>
						<th>Classes</th>
						<th>Races</th>
					</tr>
			</thead>
			<tbody>{{ range .}}
				<tr>
					<td><a href=/?itemid={{ .Id }}>{{ .Id }}</a></td>
					<td>{{ .Name }}</td>
					<td>{{ .Itemtype }}</td>
					<td>{{ if .Magic }}X{{ else }}{{ end }}</td>
					<td>{{ if .Nodrop }}X{{ else }}{{ end }}</td>
					<td>{{ if .Norent }}X{{ else }}{{ end }}</td>
					<td>{{ .Classes }}</td>
					<td>{{ .Races }}</td>
				</tr>{{ end }}
			</tbody>
		</table
`
)

const (
	ItemDetail = `
		<table class="item-detail-table">
			<tbody>
				<tr>
					<td class="item-detail-title"><img class="item-detail-icon" src="images/dasheq-default.png" alt="Item Icon"></td>
					<td class="item-detail-title">{{ .Name }}</td>
				</tr>
				<tr>
				{{ if .Nodrop }}
					<td class="item-stat"><span class="stat-label">No Drop</td>
				{{ end }}
				{{ if .Magic }}
					<td class="item-stat"><span class="stat-label">Magic</td>
				{{ end }}
				{{ if and (not .Nodrop) (not .Magic) }}
					<td height="20px"></td>
				{{ end }}
				</tr>
				<tr>
					<td class="stat-label-inner">Class:</td>
					<td></td>
					<td></td>
					<td></td>
					<td class="stat-value-inner">{{ .Classes }}</td>
				</tr>
				<tr>
					<td class="stat-label-inner">Race:</td>
					<td></td>
					<td></td>
					<td></td>
					<td class="stat-value-inner">{{ .Races }}</td>
				</tr>
				<tr>
					<td class="stat-label">{{ .Slots }}</td>
				</tr>
				<tr>
					<td class="stat-label">Type:</td>
					<td class="stat-value">{{ .Itemtype }}</td>
				</tr>
				{{ if .Reclevel }}
				<tr>
					<td class="stat-label">Rec Level:</td>
					<td class="stat-value">{{ .Reclevel }}</td>
				</tr>
				{{ end }}
				{{ if .Reqlevel }}
				<tr>
					<td class="stat-label">Req Level:</td>
					<td class="stat-value">{{ .Reqlevel }}</td>
				</tr>
				{{ end }}
				<tr>
					<td height="20px"></td>
				</tr>
				<tr>
					<td class="stat-label-inner">Size:</td>
					<td class="stat-value">{{ .Size }}</td>
					{{ if .AC }}
					<td class="stat-label-inner">AC:</td>
					<td class="stat-value-inner">{{ .AC }}</td>
					{{ else }}
					<td></td>
					<td></td>
					{{ end }}
					{{ if .Damage }}
					<td class="stat-label-inner">Base Damage:</td>
					<td class="stat-value-inner">{{ .Damage }}</td>
					{{ end }}
				</tr>
				<tr>
					<td class="stat-label-inner">Weight:</td>
					<td class="stat-value">{{ .Weight }}</td>
					{{ if .HP }}
					<td class="stat-label-inner">HP:</td>
					<td class="stat-value-inner">{{ .HP }}</td>
					{{ else }}
					<td></td>
					<td></td>
					{{ end }}
					{{ if .Delay }}
					<td class="stat-label-inner">Delay:</td>
					<td class="stat-value-inner">{{ .Delay }}</td>
					{{ else }}
					<td></td>
					<td></td>
					{{ end }}
				</tr>
				<tr>
					<td></td>
					<td></td>
					{{ if .Mana }}
					<td class="stat-label-inner">Mana:</td>
					<td class="stat-value-inner">{{ .Mana }}</td>
					{{ else }}
					<td></td>
					<td></td>
					{{ end }}
					{{ if .Delay }}
					<td class="stat-label-inner">Delay:</td>
					<td class="stat-value-inner">{{ .Delay }}</td>
					{{ else }}
					<td></td>
					<td></td>
					{{ end }}
				</tr>
			</tbody>
		</table>
		<p><a href="javascript:history.back()">< back ></a></p>
`
)
