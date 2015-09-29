package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/rcmgleite/common_io"
	"github.com/rcmgleite/notification/mailer"
)

func TestSendMail(t *testing.T) {
	common_io.Setup(&common_io.Config{
		ModuleName: "testSendMail",
	})
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
