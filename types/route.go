package types

type REQUEST_TYPE string
type DESCRIPTION string

type IDENTIFIER interface {
	ToString() string
}

type ROUTE_INFO struct {
	Identifier IDENTIFIER
	Value      string
}

const (
	GET     REQUEST_TYPE = "@get"
	POST    REQUEST_TYPE = "@post"
	PATCH   REQUEST_TYPE = "@patch"
	PUT     REQUEST_TYPE = "@put"
	DELETE  REQUEST_TYPE = "@delete"
	OPTIONS REQUEST_TYPE = "@options"
)

type ExtractRouteInfo struct {
	RequestParsingRequired bool
	Route                  string
	PathParams             []string
	Description            string
}

func (r REQUEST_TYPE) ToString() string {
	return string(r)
}

func (d DESCRIPTION) ToString() string {
	return string(d)
}
