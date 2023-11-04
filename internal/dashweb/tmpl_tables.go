package dashweb

const (
	TmplHome = `
        <p>A simple, fast, EQEmu data browser written in Go.</p>
`
)

const (
	TmplDataSets = `
        <table class="data-table">
            <thead>
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
                    </tr>{{end}}
                </tbody>
        </table>
`
)

const (
	TmplZones = `
        <table class="data-table">
            <thead>
                <tr>
                    <th class="integer">ID</th>
                    <th>Short Name</th>
                    <th>Long Name</th>
                    <th>Outdoor?</th>
                    <th>Expansion</th>
                </tr>
            </thead>
                <tbody>{{range .}}
                    <tr>
                        <td class="integer">{{ .Id }}</td>
                        <td>{{ .Short_name }}</td>
                        <td>{{ .Long_name }}</td>
                        <td>{{ if .Outdoor }}Yes{{ else }}No{{ end }}</td>
                        <td>{{ if eq .Expansion 1 }}Original</td>
                            {{ else if eq .Expansion 2 }}Ruins of Kunark</td>
                            {{ else if eq .Expansion 3 }}Scars of Velious</td>
                            {{ else if eq .Expansion 4 }}Shadows of Luclin</td>
                            {{ else if eq .Expansion 5 }}Planes of Power</td>
                            {{ else }}Unknown</td>{{ end }}
                    </tr>{{end}}
                </tbody>
        </table>
`
)
