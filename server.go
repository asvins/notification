package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asvins/common_io"
	"github.com/asvins/notification/mailer"
	"github.com/asvins/utils/config"
)

var (
	producer *common_io.Producer
	consumer *common_io.Consumer
)

func main() {

	// common_io
	cfg := common_io.Config{}
	err := config.Load("common_io_config.gcfg", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// producer setup
	producer, err = common_io.NewProducer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer producer.TearDown()

	// consumer setup
	consumer = common_io.NewConsumer(cfg)
	consumer.HandleTopic("send_mail", mailer.SendMail)
	if err = consumer.StartListening(); err != nil {
		log.Fatal(err)
	}

	defer consumer.TearDown()

	// Server config
	serverConf := Config{}
	err = config.Load("notification_config.gcfg", &serverConf)
	if err != nil {
		log.Fatal(err)
	}

	r := DefRoutes()

	fmt.Println("Server running on port: ", serverConf.Server.Port)

	http.ListenAndServe(":"+serverConf.Server.Port, r)
}
