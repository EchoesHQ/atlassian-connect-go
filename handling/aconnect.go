package handling

// AtlassianConnect is auto generated by github.com/perrito666/LAC from a json file
type AtlassianConnect struct {
	Authentication Authentication         `json:"authentication"`
	BaseURL        string                 `json:"baseUrl"`
	Description    string                 `json:"description"`
	Key            string                 `json:"key"`
	Lifecycle      Lifecycle              `json:"lifecycle"`
	Modules        map[string]interface{} `json:"modules"`
	Name           string                 `json:"name"`
	Scopes         []string               `json:"scopes"`
	Vendor         Vendor                 `json:"vendor"`
}

// Authentication is auto generated by github.com/perrito666/LAC from a json file
type Authentication struct {
	Type string `json:"type"`
}

// Conditions is auto generated by github.com/perrito666/LAC from a json file
type Conditions struct {
	Condition string `json:"condition"`
	Or        []Or   `json:"or"`
}

// Description is auto generated by github.com/perrito666/LAC from a json file
type Description struct {
	Value string `json:"value"`
}

// JiraIssueFields is auto generated by github.com/perrito666/LAC from a json file
type JiraIssueFields struct {
	Description Description `json:"description"`
	Key         string      `json:"key"`
	Name        Name        `json:"name"`
	Type        string      `json:"type"`
}

// WebPanel is auto generated by github.com/perrito666/LAC from a json file
type WebPanel struct {
	Conditions []Conditions `json:"conditions"`
	Context    string       `json:"context"`
	Key        string       `json:"key"`
	Location   string       `json:"location"`
	Name       Name         `json:"name"`
	URL        string       `json:"url"`
	Weight     float64      `json:"weight"`
}

// Lifecycle is auto generated by github.com/perrito666/LAC from a json file
type Lifecycle struct {
	Installed   string `json:"installed"`
	UnInstalled string `json:"uninstalled"`
	Enabled     string `json:"enabled"`
	Disabled    string `json:"disabled"`
}

// Name is auto generated by github.com/perrito666/LAC from a json file
type Name struct {
	Value string `json:"value"`
}

// Or is auto generated by github.com/perrito666/LAC from a json file
type Or struct {
	Condition string `json:"condition"`
}

// Vendor is auto generated by github.com/perrito666/LAC from a json file
type Vendor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Webhooks is auto generated by github.com/perrito666/LAC from a json file
type Webhooks struct {
	Event string `json:"event"`
	URL   string `json:"url"`
}
