package types

type SampleMap map[string]uint8

type SampleStruct struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	IsAvailable bool    `json:"is_available"`
}

type SampleRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      uint8  `json:"age"`
}

type SampleResponse struct {
	Status  string       `json:"status"`
	Product SampleStruct `json:"product"`
	Map     SampleMap    `json:"map"`
}
