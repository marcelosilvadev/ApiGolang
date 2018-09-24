package model

//ResponseSuccess representa o response caso retorne sucesso ...
type ResponseSuccess struct {
	Meta    Meta        `json:"meta"`
	Records interface{} `json:"records`
}

//Meta representa o subresponse caso retorne sucesso ...
type Meta struct {
	Limit       int `json:"limit"`
	Offset      int `json:"offset"`
	RecordCount int `json:"recordCount`
}
