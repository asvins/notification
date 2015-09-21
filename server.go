package main

import (
	"fmt"
	"net/http"

	"github.com/rcmgleite/common_io"
	"github.com/rcmgleite/labSoft2_Estoque/router"
	"github.com/rcmgleite/notification/mailer"
)

// Stupid example
func printStdout(msg []byte) {
	fmt.Println(string(msg))
}

func main() {
	r := router.NewRouter()
	http.Handle("/", r)

	// Here we can subscribe to tags like send_sms, send_mail, send_push_notification, etc ...
	common_io.Setup()
	defer common_io.TearDown()

	common_io.Subscribe("send_mail", mailer.SendMail)
	common_io.Subscribe("print_stdout", printStdout)

	fmt.Println("Server running on port: 8080")

	http.ListenAndServe(":8080", nil)
}
