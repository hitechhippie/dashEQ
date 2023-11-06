package dashweb

const (
	TmplSkills = `
        <table class="data-table">
            <thead>
                <tr>
                    <th class="integer">Level</th>
                    <th>Class</th>
                    <th>Skill</th>
                    <th class="integer">Cap</th>
                    <th>Skill Cap List By Class</th>
                </tr>
            </thead>
                <tbody>{{range .}}
                    <tr>
                        <td class="integer">{{ .Level }}</td>
                        <td>{{ .Class }}</td>
                        <td>{{ .Skill }}</td>
                        <td class="integer">{{ .Cap }}</td>
                    </tr>{{end}}
                </tbody>
        </table>
        <script type="text/javascript" src="/includes/tableSortable.js"></script>
`
)
