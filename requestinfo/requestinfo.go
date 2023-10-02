package requestinfo

type RequestInfo struct {
	SelectionSetList    []string       `json:"selectionSetList"`
	SelectionSetGraphQL string         `json:"selectionSetGraphQL"`
	ParentTypeName      string         `json:"parentTypeName"`
	FieldName           string         `json:"fieldName"`
	Variables           map[string]any `json:"variables"`
}

type WithRequestInfo struct {
	Info *RequestInfo `json:"requestInfo"`
}
