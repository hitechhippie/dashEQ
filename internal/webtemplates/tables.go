package webtemplates

const (
	Home = `                    <tr>
                        <th>dashEQ Data Sets</th>
                    </tr>
                    <tr>
                        <th>DataSet</th>
                        <th>Qty</th>
                        <th>LoadTime</th>
                    </tr>
                </thead>
                <tbody>{{range .}}
                    <tr>
                        <td>{{ .Name }}</td>
                        <td>{{ .Count }}</td>
                        <td>{{ .LoadTime }}</td>
                    </tr>{{end}}`
)

const (
	Zones = `                    <tr>
                        <th>Zones</th>
                    </tr>
                    <tr>
                        <th>ID</th>
                        <th>Short Name</th>
                        <th>Long Name</th>
                        <th>Outdoor?</th>
                        <th>Expansion</th>
                    </tr>
                </thead>
                <tbody>{{range .}}
                    <tr>
                        <td>{{ .Id }}</td>
                        <td>{{ .Short_name }}</td>
                        <td>{{ .Long_name }}</td>
                        <td>{{ if eq .Outdoor 1 }}Yes{{ else }}No{{ end }}</td>
                        <td>{{ if eq .Expansion 1 }}Original</td>
                            {{ else if eq .Expansion 2 }}Ruins of Kunark</td>
                            {{ else if eq .Expansion 3 }}Scars of Velious</td>
                            {{ else if eq .Expansion 4 }}Shadows of Luclin</td>
                            {{ else if eq .Expansion 5 }}Planes of Power</td>
                            {{ else }}Unknown</td>{{ end }}
                    </tr>{{end}}`
)

const (
	QuestNPCs = `                <tr>
                        <th>Quest NPCs</th>
                    </tr>
                    <tr>
                        <th>NPC ID</th>
                        <th>Name</th>
                        <th>Zone</th>
                        <th>Quest File</th>
                        <th>Quest Text</th>
                    </tr>
                </thead>
                <tbody>{{range .}}
                    <tr>
                        <td><a href=/?questnpcid={{ .Id }}>{{ .Id }}</a></td>
                        <td>{{ .Name }}</td>
                        <td>{{ .Zone }}</td>
                        <td>{{ .File }}</td>
                    </tr>{{end}}`
)

const (
	QuestNPCdetail = `                <tr>
                        <th>Quest NPC Detail</th>
                    </tr>
                    <tr>
                        <th>Player Says</th>
                        <th>NPC Responds</th>
                    </tr>
                </thead>
                <tbody>{{range .}}
                    <tr>
                        <td>{{ .QuestNPCName }} hears, "{{ .Hears }}"</td>
                        <td>{{ .QuestNPCName }} responds, "{{ .Says }}"</td>
                    </tr>{{end}}
                    <tr>
                        <td><a href="javascript:history.back()">< back ></a></td>`
)
