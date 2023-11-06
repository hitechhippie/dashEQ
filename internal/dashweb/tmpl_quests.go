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
                    <th>Quest NPCs</th>
                </tr>
            </thead>
                <tbody>{{range .}}
                    <tr>
                        <td><a href=/?questnpcid={{ .Id }}>{{ .Id }}</a></td>
                        <td>{{ .Name }}</td>
                        <td>{{ .ZoneName }}</td>
                        <td>{{ .File }}</td>
                    </tr>{{end}}
                </tbody>
        </table>
        <script type="text/javascript" src="/includes/tableSortable.js"></script>
`
)

const (
	TmplQuestNPCdetail = `
        <table class="data-table">
            <thead>
                <tr>
                    <th>Quest NPC Detail</th>
                </tr>
            </thead>
                <tbody>{{range .}}
                    <tr>
                        <td>{{ .QuestNPCName }} hears, "{{ .Hears }}"</td>
                        <td>{{ .QuestNPCName }} responds, "{{ .Says }}"</td>
                    </tr>{{end}}
                </tbody>
        </table>
        <a href="javascript:history.back()">< back ></a>
        <script type="text/javascript" src="/includes/tableSortable.js"></script>
    `
)
