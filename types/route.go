package types

type HTTP_METHOD string
type ROUTE_INFO string

const (
	METHOD ROUTE_INFO = "@method"
	PATH   ROUTE_INFO = "@path"

	METHOD_GET     HTTP_METHOD = "GET"
	METHOD_PUT     HTTP_METHOD = "PUT"
	METHOD_HEAD    HTTP_METHOD = "HEAD"
	METHOD_POST    HTTP_METHOD = "POST"
	METHOD_PATCH   HTTP_METHOD = "PATCH"
	METHOD_TRACE   HTTP_METHOD = "TRACE"
	METHOD_DELETE  HTTP_METHOD = "DELETE"
	METHOD_OPTIONS HTTP_METHOD = "OPTIONS"
)

var (
	HasBodyMethods = map[HTTP_METHOD]bool{
		METHOD_POST:   true,
		METHOD_PUT:    true,
		METHOD_DELETE: true,
		METHOD_PATCH:  true,
	}
)