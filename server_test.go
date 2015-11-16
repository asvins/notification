package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/asvins/common_io"
	"github.com/asvins/notification/mailer"
	"github.com/asvins/utils/config"
)

func TestSendMail(t *testing.T) {
	cfg := common_io.Config{}
	err := config.Load("common_io_config.gcfg", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	producer, err := common_io.NewProducer(cfg)
	if err != nil {
		t.Error(err)
	}

	defer producer.TearDown()

	m := mailer.Mail{
		To:      []string{"asvins.poli@gmail.com"},
		Subject: "Test from Asvins server",
		Body:    "Test Message from Asvins Servers.\n -- Asvins Team",
	}

	b, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	producer.Publish("send_mail", b)
	fmt.Println(">>Send_mail message 1 published!")
}
