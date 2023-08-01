package templates

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
                        <td>{{ .Outdoor }}</td>
                        <td>{{ .Expansion }}</td>
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
                        <td>{{ .Id }}</td>
                        <td>{{ .Name }}</td>
                        <td>{{ .Zone }}</td>
                        <td>{{ .File }}</td>
                    </tr>{{end}}`
)

const (
	QuestNPCdetail = `                <tr>
                        <th>{{ .QuestNPCName }}</th>
                    </tr>
                    <tr>
                        <th>Player Says</th>
                        <th>NPC Responds</th>
                    </tr>
                </thead>
                <tbody>{{range .}}
                    <tr>
                        <td>{{ .Hears }}</td>
                        <td>{{ .Says }}</td>
                    </tr>{{end}}
    `
)
