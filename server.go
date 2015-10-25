package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/asvins/common_io"
	"github.com/asvins/notification/mailer"
	"github.com/asvins/router"
	"github.com/asvins/utils/config"
)

func main() {

	topics := make(map[string]common_io.CallbackFunc)
	topics["send_mail"] = mailer.SendMail

	cfg := common_io.Config{}
	err := config.Load("common_io_config.gcfg", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	cfg.ModuleName = "notification"
	cfg.Topics = topics

	common_io.Setup(&cfg)
	defer common_io.TearDown()

	serverConf := Config{}
	err = config.Load("notification_config.gcfg", &serverConf)
	if err != nil {
		log.Fatal(err)
	}

	r := router.NewRouter()
	r.AddRoute("/api/discovery", router.GET, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "NOTIFICATION MODULE - TODO")
	})

	// TODO remover
	r.AddRoute("/api/mailTest", router.GET, func(w http.ResponseWriter, r *http.Request) {
		m := mailer.Mail{
			To:      []string{"asvins.poli@gmail.com"},
			Subject: "Test from Asvins server",
			Body:    "Test Message from Asvins Servers.\n -- Asvins Team",
		}

		b, err := json.Marshal(&m)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

		common_io.Publish("send_mail", b)
		fmt.Fprintf(w, ">>Send_mail message 1 published!")

	})

	http.Handle("/", r)

	fmt.Println("Server running on port: ", serverConf.Server.Port)

	http.ListenAndServe(":"+serverConf.Server.Port, nil)
}
