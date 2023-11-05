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
`
)

const (
	TmplSkillsByClass = `
        <table class="data-table">
            <thead>
                <tr>
                    <th></th>
                    <th></th>
                    <th class="integer">L1</th>
                    <th class="integer">L2</th>
                    <th class="integer">L3</th>
                    <th class="integer">L4</th>
                    <th class="integer">L5</th>
                    <th class="integer">L6</th>
                    <th class="integer">L7</th>
                    <th class="integer">L8</th>
                    <th class="integer">L9</th>
                    <th class="integer">L10</th>
                    <th class="integer">L11</th>
                    <th class="integer">L12</th>
                    <th class="integer">L13</th>
                    <th class="integer">L14</th>
                    <th class="integer">L15</th>
                    <th class="integer">L16</th>
                    <th class="integer">L17</th>
                    <th class="integer">L18</th>
                    <th class="integer">L19</th>
                    <th class="integer">L20</th>
                </tr>
            </thead>
                <tbody>{{range .}}
                    <tr>
                        <td>{{ .Class }}</td>
                        <td>{{ .Skill }}</td>
                        <td>{{ .Name }}</td>
                        <td class="integer">{{ .Level }}</td>
                    </tr>{{end}}
                </tbody>
        </table>
`
)
