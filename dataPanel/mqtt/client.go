package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"ons/util"
	"strconv"
	"time"
)

var MqttClient mqtt.Client

func InitMqttClient() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(viper.GetString("mqtt.broker"))
	opts.SetClientID("svc|ons|457765|" + strconv.Itoa(util.Random4Num()))
	opts.SetKeepAlive(time.Second * 5)
	opts.SetUsername("mqtt.username")
	opts.SetPassword("mqtt.password")

	MqttClient = mqtt.NewClient(opts)
	if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	MqttClient.Subscribe(onsShareGroup+requestTopic, 0, omCallback)
}
