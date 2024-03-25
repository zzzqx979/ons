package db

type Service struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	Desc     string        `json:"desc"`
	Required bool          `json:"required"`
	CallType string        `json:"callType"`
	Input    []InputOutput `json:"input_param"`
	Output   []InputOutput `json:"output_param"`
}

type InputOutput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	DataType
}
