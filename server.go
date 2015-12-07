package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asvins/common_io"
	"github.com/asvins/utils/config"
)

var (
	producer *common_io.Producer
	consumer *common_io.Consumer
)

func main() {

	/*
	*	Server config
	 */
	serverConf := Config{}
	err := config.Load("notification_config.gcfg", &serverConf)
	if err != nil {
		log.Fatal(err)
	}

	r := DefRoutes()

	/*
	*	Common io
	 */
	setupCommonIo()

	fmt.Println("Server running on port: ", serverConf.Server.Port)
	http.ListenAndServe(":"+serverConf.Server.Port, r)
}
