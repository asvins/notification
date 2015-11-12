package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asvins/notification/mailer"
	"github.com/asvins/router"
	"github.com/asvins/router/errors"
)

func DefRoutes() *router.Router {
	r := router.NewRouter()
	r.Handle("/api/discovery", router.GET, func(w http.ResponseWriter, r *http.Request) errors.Http {
		fmt.Fprint(w, "NOTIFICATION MODULE - TODO")
		return nil
	}, []router.Interceptor{})

	r.Handle("/api/mailTest", router.GET, func(w http.ResponseWriter, r *http.Request) errors.Http {
		m := mailer.Mail{
			To:      []string{"asvins.poli@gmail.com"},
			Subject: "Test from Asvins server",
			Body:    "Test Message from Asvins Servers.\n -- Asvins Team",
		}

		b, err := json.Marshal(&m)
		if err != nil {
			return errors.BadRequest(err.Error())
		}

		producer.Publish("send_mail", b)
		fmt.Fprintf(w, "[INFO] Email sent")
		return nil
	}, []router.Interceptor{})

	return r
}
