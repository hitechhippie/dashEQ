package dashweb

const (
	TmplSkillsByClass = `
        <table class="data-table">
            <thead>
                <tr>
                    <th>Class</th>
                    <th class="integer">ID</th>
                    <th>Name</th>
                    <th class="integer">Level</th>
                    <th></th>
                    <th></th>
                </tr>
            </thead>
                <tbody>{{range .}}
                    <tr>
                        <td>{{ .Class }}</td>
                        <td class="integer">{{ .Id }}</td>
                        <td>{{ .Name }}</td>
                        <td class="integer">{{ .Level }}</td>
                        <td><img src="/icons/{{ .NewIcon }}.gif" /></td>
                        <td><a href="/items.html?itemid={{ .Scroll }}"><img src="/icons/item_504.gif" /></td>
                    </tr>{{end}}
                </tbody>
        </table>
`
)
