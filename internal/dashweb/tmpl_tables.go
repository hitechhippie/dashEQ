package dashweb

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
                        <td>{{ if .Outdoor }}Yes{{ else }}No{{ end }}</td>
                        <td>{{ if eq .Expansion 1 }}Original</td>
                            {{ else if eq .Expansion 2 }}Ruins of Kunark</td>
                            {{ else if eq .Expansion 3 }}Scars of Velious</td>
                            {{ else if eq .Expansion 4 }}Shadows of Luclin</td>
                            {{ else if eq .Expansion 5 }}Planes of Power</td>
                            {{ else }}Unknown</td>{{ end }}
                    </tr>{{end}}`
)

const (
	Items = `                    <tr>
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
                    </tr>{{ end }}`
)

const (
	ItemDetail = `
        <div class="item-window">
            <div class="item-icon">
                <img src="favicon.png" alt="Item Icon">
            </div>
            <div class="item-details">
                <h1 class="item-name">{{ .Name }}</h1>
                <div class="item-stats">
                    <div class="item-stat">
                        {{ if .Nodrop }}<span class="stat-label">No Trade</span><br><br>{{ end }}
                    </div>
                    <div class="item-stat">
                        <span class="stat-label">Type: </span>
                        <span class="stat-value">{{ .Itemtype }}</span>
                    </div>
                    <div class="item-stat">
                        <span class="stat-label">Class: </span>
                        <span class="stat-value">{{ .Classes }}</span>
                    </div>
                    <div class="item-stat">
                        <span class="stat-label">Race: </span>
                        <span class="stat-value">{{ .Races }}</span>
                    </div>
                    <div class="item-stat">
                        <span class="stat-label">{{ .Slots }}</span>
                    </div>
                    <br><br>
                    <div class="item-stat">
                        <span class="stat-label">Size:</span>
                        <span class="stat-value">{{ .Size }}</span>{{ if .AC }}
                        <span class="stat-label">AC:</span>
                        <span class="stat-value">{{ .AC }}</span>{{ else }}
                        <span class="stat-label"> </span>
                        <span class="stat-value"> </span>{{ end }}
                    </div>
                    <div class="item-stat">
                        <span class="stat-label">Weight:</span>
                        <span class="stat-value">{{ .Weight }}</span>{{ if .HP }}
                        <span class="stat-label">HP:</span>
                        <span class="stat-value">{{ .HP }}</span>{{ else }}
                        <span class="stat-label"> </span>
                        <span class="stat-value"> </span>{{ end }}
                    </div>{{ if .Reclevel }}
                    <div class="item-stat">
                        <span class="stat-label">Rec Level:</span>
                        <span class="stat-value">{{ .Reclevel }}</span>{{ if .Mana }}
                        <span class="stat-label">Mana:</span>
                        <span class="stat-value">{{ .Mana }}</span>{{ end }}
                    </div>{{ else }}<div class="item-stat">
                    <span class="stat-label"> </span>
                    <span class="stat-value"> </span>{{ if .Mana }}
                    <span class="stat-label">Mana:</span>
                    <span class="stat-value">{{ .Mana }}</span>{{ end }}
                </div>{{ end }}
                    </div>{{ if .Reqlevel }}
                    <div class="item-stat">
                        <span class="stat-label">Req Level:</span>
                        <span class="stat-value">{{ .Reqlevel }}</span>
                    </div>{{ else }}<div class="item-stat">
                    <span class="stat-label"> </span>
                    <span class="stat-value"> </span>
                </div>{{ end }}
                    <!-- Add more item stats as needed -->
                </div>
            </div>
        </div>
        <p><a href="javascript:history.back()">< back ></a></p>
`
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
