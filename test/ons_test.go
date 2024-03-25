package test

import (
	"encoding/json"
	"ons/db"
	"testing"
)

func Test(t *testing.T) {
	services := []db.Service{
		{
			Id:       "service1",
			Name:     "服务1",
			Desc:     "服务测试1",
			Required: true,
			CallType: "sync",
			Input: []db.InputOutput{
				{
					Id:   "input1",
					Name: "入参1",
					DataType: db.DataType{
						Type: "int",
						Spec: db.Spec{
							Min:       0,
							Max:       100,
							Unit:      "meter",
							UnitName:  "米",
							Size:      0,
							Step:      1,
							Length:    0,
							BoolFalse: "",
							BoolTrue:  "",
							ItemType:  "",
						},
					},
				},
			},
			Output: nil,
		},
	}
	properties := []db.Property{
		{
			Id:         "property1",
			Name:       "属性1",
			AccessMode: "rw",
			Required:   true,
			DataType: db.DataType{
				Type: "bool",
				Spec: db.Spec{
					Min:       nil,
					Max:       nil,
					Unit:      "",
					UnitName:  "",
					Size:      0,
					Step:      0,
					Length:    0,
					BoolFalse: "off",
					BoolTrue:  "on",
					ItemType:  "",
				},
			},
		},
	}
	events := []db.Event{
		{
			Id:       "event1",
			Name:     "事件1",
			Desc:     "测试事件1",
			Type:     "info",
			Required: true,
			Output: []db.InputOutput{
				{
					Id:   "output1",
					Name: "出参1",
					DataType: db.DataType{
						Type: "string",
						Spec: db.Spec{
							Min:       nil,
							Max:       nil,
							Unit:      "",
							UnitName:  "",
							Size:      0,
							Step:      0,
							Length:    64,
							BoolFalse: "",
							BoolTrue:  "",
							ItemType:  "",
						},
					},
				},
			},
		},
	}

	servicesByte, _ := json.Marshal(services)
	propertiesByte, _ := json.Marshal(properties)
	eventsByte, _ := json.Marshal(events)
	model := db.ObjectModel{
		Name:        "model1",
		Status:      1,
		Note:        "物模型1",
		ProductCode: "123",
		FactoryCode: "321",
		Properties:  string(propertiesByte),
		Events:      string(eventsByte),
		Services:    string(servicesByte),
	}
	t.Log(model)
}
