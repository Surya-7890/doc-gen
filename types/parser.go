package types

type METHOD string

const (
	POST METHOD = "post"
	GET  METHOD = "get"
)

type ParsedFunction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Method      METHOD `json:"method"`
}
