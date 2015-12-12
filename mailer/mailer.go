package mailer

import (
	"encoding/json"
	"fmt"
	"net/smtp"
)

// Mail ...
type Mail struct {
	To      []string
	Subject string
	Body    string
}

func doSendMail(m Mail) error {
	// Set up authentication information. -- FIXME - CONFIG FILE / env variables
	auth := smtp.PlainAuth(
		"",
		"asvins.poli@gmail.com",
		"asvins.poli1010",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"asvins.poli@gmail.com",
		m.To,
		[]byte("MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\"\r\n"+"Subject:"+m.Subject+"\r\n"+m.Body),
	)
	if err != nil {
		return err
	}
	return nil
}

//SendMail ...
func SendMail(msg []byte) {
	fmt.Println("SendMail Called")
	var m Mail
	err := json.Unmarshal(msg, &m)

	if err != nil {
		fmt.Println(">> Unable to Unmarshal json to struct")
		return
	}

	err = doSendMail(m)
	if err != nil {
		fmt.Println(err)
	}
}
