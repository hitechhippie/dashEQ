package dashweb

const (
	TmplSpells = `
        <table class="data-table">
            <thead>
                <tr>
                    <th class="integer">ID</th>
                    <th>Name</th>
                    <th class="integer">WAR</th>
                    <th class="integer">CLR</th>
                    <th class="integer">PAL</th>
                    <th class="integer">RNG</th>
                    <th class="integer">SHD</th>
                    <th class="integer">DRU</th>
                    <th class="integer">MNK</th>
                    <th class="integer">BRD</th>
                    <th class="integer">ROG</th>
                    <th class="integer">SHM</th>
                    <th class="integer">NEC</th>
                    <th class="integer">WIZ</th>
                    <th class="integer">MAG</th>
                    <th class="integer">ENC</th>
                    <th class="integer">BST</th>
                    <th class="integer">BER</th>
                </tr>
            </thead>
                <tbody>{{range .}}
                    <tr>
                        <td class="integer">{{ .Id }}</td>
                        <td>{{ .Name }}</td>
                        <td class="integer">{{ if eq .Classes1 255 }}-</td>
                            {{ else }}{{ .Classes1 }}</td>{{ end }}
                            <td  class="integer">{{ if eq .Classes2 255 }}-</td>
                            {{ else }}{{ .Classes2 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes3 255 }}-</td>
                            {{ else }}{{ .Classes3 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes4 255 }}-</td>
                            {{ else }}{{ .Classes4 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes5 255 }}-</td>
                            {{ else }}{{ .Classes5 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes6 255 }}-</td>
                            {{ else }}{{ .Classes6 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes7 255 }}-</td>
                            {{ else }}{{ .Classes7 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes8 255 }}-</td>
                            {{ else }}{{ .Classes8 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes9 255 }}-</td>
                            {{ else }}{{ .Classes9 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes10 255 }}-</td>
                            {{ else }}{{ .Classes10 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes11 255 }}-</td>
                            {{ else }}{{ .Classes11 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes12 255 }}-</td>
                            {{ else }}{{ .Classes12 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes13 255 }}-</td>
                            {{ else }}{{ .Classes13 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes14 255 }}-</td>
                            {{ else }}{{ .Classes14 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes15 255 }}-</td>
                            {{ else }}{{ .Classes15 }}</td>{{ end }}
                            <td class="integer">{{ if eq .Classes16 255 }}-</td>
                            {{ else }}{{ .Classes16 }}</td>{{ end }}
                    </tr>{{end}}
                </tbody>
        </table>
`
)

const (
	TmplSpellsByClass = `
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
