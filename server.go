package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asvins/common_io"
	"github.com/asvins/notification/mailer"
	"github.com/rcmgleite/router"
)

func main() {
	r := router.NewRouter()
	http.Handle("/", r)

	topics := make(map[string]common_io.CallbackFunc)
	topics["send_mail"] = mailer.SendMail

	cfg, err := common_io.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg.ModuleName = "notification"
	cfg.Topics = topics

	common_io.Setup(cfg)
	defer common_io.TearDown()

	serverConf, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server running on port: ", serverConf.Server.Port)

	http.ListenAndServe(":"+serverConf.Server.Port, nil)
}
