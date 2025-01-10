package types

type REQUEST_TYPE string

const (
	GET     REQUEST_TYPE = "get"
	POST    REQUEST_TYPE = "post"
	PATCH   REQUEST_TYPE = "patch"
	PUT     REQUEST_TYPE = "put"
	DELETE  REQUEST_TYPE = "delete"
	OPTIONS REQUEST_TYPE = "options"
)
