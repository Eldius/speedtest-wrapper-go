package mqttclient

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Eldius/speedtest-wrapper-go/config"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

/*
SendTestResult sends the ping to MQTT broker
*/
func SendTestResult(p interface{}, cfg config.MQTTConfig) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", cfg.Host, cfg.Port))
	opts.SetClientID(cfg.ClientName)

	if cfg.User != "" {
		fmt.Println("Using user:", cfg.User)
		opts.SetUsername(cfg.User)
	}

	if cfg.Pass != "" {
		fmt.Println("Using pass: ***")
		opts.SetPassword(cfg.Pass)
	}
	//opts.SetCleanSession(*cleansess)

	fmt.Println("Connecting to:", fmt.Sprintf("tcp://%s:%s", cfg.Host, cfg.Port))
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Failed to conect to broker")
		panic(token.Error())
	}
	log.Println("Sample Publisher Started")
	log.Println("---- doing publish ----")
	token := client.Publish(cfg.Topic, cfg.Qos, false, serialize(p))
	token.Wait()

	client.Disconnect(250)
	log.Println("Sample Publisher Disconnected")
}

func serialize(obj interface{}) []byte {
	if data, err := json.Marshal(obj); err != nil {
		panic(err.Error())
	} else {
		log.Println(fmt.Sprintf("serialized:\n%s", string(data)))
		return data
	}
}
