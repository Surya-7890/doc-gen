package types

type Route struct {
	Path   string
	Method string
	Body   map[string]interface{}
}
