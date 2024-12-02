package types

type Route struct {
	Name        string                 `json:"name"`
	Path        string                 `json:"path"`
	Method      HTTP_METHOD            `json:"method"`
	Body        map[string]interface{} `json:"body"`
	Description string                 `json:"description"`
}
