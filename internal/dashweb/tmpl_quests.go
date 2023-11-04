package dashweb

const (
	TmplQuestNPCs = `
        <table class="data-table">
            <thead>
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
                    </tr>{{end}}
                </tbody>
        </table>
`
)

const (
	TmplQuestNPCdetail = `                <tr>
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
