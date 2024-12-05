package types

type RequestBody struct {
	Name   string `json:"name"`
	Age    int32  `json:"age"`
	Mobile string `json:"mobile"`
}

type FUNC_TYPE string

const (
	NONE        FUNC_TYPE = "None"
	HANDLE      FUNC_TYPE = "Handle"
	HANDLE_FUNC FUNC_TYPE = "HandleFunc"
)
