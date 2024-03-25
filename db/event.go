package db

type Event struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	Desc     string        `json:"desc"`
	Type     string        `json:"type"`
	Required bool          `json:"required"`
	Output   []InputOutput `json:"output"`
}
