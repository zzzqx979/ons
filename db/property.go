package db

type Property struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	AccessMode string `json:"access_mode"`
	Required   bool   `json:"required"`
	DataType
}
