package main

import (
	"fmt"
	"net/http"

	"github.com/rcmgleite/common_io"
	"github.com/rcmgleite/labSoft2_Estoque/router"
	"github.com/rcmgleite/notification/mailer"
)

func main() {
	r := router.NewRouter()
	http.Handle("/", r)

	topics := make(map[string]common_io.CallbackFunc)
	topics["send_mail"] = mailer.SendMail

	config := &common_io.Config{
		ModuleName: "notification",
		Topics:     topics,
	}

	common_io.Setup(config)
	defer common_io.TearDown()

	fmt.Println("Server running on port: 8080")

	http.ListenAndServe(":8080", nil)
}
