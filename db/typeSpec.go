package db

type DataType struct {
	Type string `json:"type"`
	Spec
}

type Spec struct {
	Min       interface{} `json:"min"`
	Max       interface{} `json:"max"`
	Unit      string      `json:"unit"`
	UnitName  string      `json:"unitName"`
	Size      int         `json:"size"`
	Step      int         `json:"step"`
	Length    int         `json:"length"`
	BoolFalse string      `json:"bool_false"`
	BoolTrue  string      `json:"bool_true"`
	ItemType  string      `json:"itemType"`
}
