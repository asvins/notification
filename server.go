package main

import (
	"fmt"
	"net/http"

	"github.com/rcmgleite/asvinsKafka"
	"github.com/rcmgleite/labSoft2_Estoque/router"
	"github.com/rcmgleite/notification_module/mailer"
)

// Stupid example
func printStdout(msg []byte) {
	fmt.Println(string(msg))
}

func main() {
	r := router.NewRouter()
	http.Handle("/", r)

	// Here we can subscribe to tags like send_sms, send_mail, send_push_notification, etc ...
	asvinsKafka.Setup()
	defer asvinsKafka.TearDown()

	asvinsKafka.Subscribe("send_mail", mailer.SendMail)
	asvinsKafka.Subscribe("print_stdout", printStdout)

	fmt.Println("Server running on port: 8080")

	http.ListenAndServe(":8080", nil)
}
