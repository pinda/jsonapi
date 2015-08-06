package jsonapi

type OnePayload struct {
	Links    map[string]string `json:"links,omitempty"`
	Meta     map[string]string `json:"meta,omitempty"`
	Data     *Node             `json:"data"`
	Included []*Node           `json:"included,omitempty"`
	JsonApi  map[string]string `json:"jsonapi,omitempty"`
}

type ManyPayload struct {
	Links    map[string]string `json:"links,omitempty"`
	Meta     map[string]string `json:"meta,omitempty"`
	Data     []*Node           `json:"data"`
	Included []*Node           `json:"included,omitempty"`
}

type Node struct {
	Type          string                 `json:"type"`
	Id            string                 `json:"id"`
	Attributes    map[string]interface{} `json:"attributes,omitempty"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	Links         map[string]string      `json:"links,omitempty"`
}

type RelationshipOneNode struct {
	Data  *Node             `json:"data"`
	Links map[string]string `json:"links,omitempty"`
	Meta  map[string]string `json:"meta,omitempty"`
}

type RelationshipManyNode struct {
	Data  []*Node           `json:"data"`
	Links map[string]string `json:"links,omitempty"`
	Meta  map[string]string `json:"meta,omitempty"`
}
