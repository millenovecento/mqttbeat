package config

import "time"

// Config represents every needed configuration fields
type Config struct {
	BrokerURL       string        `config:"broker_url"`
	BrokerUsername  string        `config:"broker_username"`
	BrokerPassword  string        `config:"broker_password"`
	BrokerClientId  string        `config:"broker_client_id"`
	TopicsSubscribe []string      `config:"topics_subscribe"`
	DecodePaylod    bool          `config:"decode_payload"`
	Period          time.Duration `config:"period"`
}

// DefaultConfig will be used if no config file is founded
var DefaultConfig = Config{
	BrokerURL:       "tcp://localhost:1883",
	BrokerUsername:	 "",
	BrokerPassword:	 "",
	BrokerClientId:	 "",
	TopicsSubscribe: []string{"/test/mqttbeat/#?1"},
	DecodePaylod:    true,
}