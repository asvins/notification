package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/asvins/common_io"
	"github.com/asvins/notification/mailer"
)

func TestSendMail(t *testing.T) {
	cfg, err := common_io.LoadConfig()
	if err != nil {
		t.Error(err)
	}
	cfg.ModuleName = "testSendMail"

	common_io.Setup(cfg)
	defer common_io.TearDown()

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

	common_io.Publish("send_mail", b)
	fmt.Println(">>Send_mail message 1 published!")

	common_io.Publish("send_mail", b)
	fmt.Println(">>Send_mail message 2 published!")
}
