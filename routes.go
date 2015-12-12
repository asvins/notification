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
			Subject: "Está na hora de tomar sua medicação!",
			Body: `
			<!DOCTYPE html>
			<html lang="pt-br">
				<head>
					<meta charset="utf-8">
					<title>Asvins - Hora do próximo Pack </title>
				</head>
				<body style="color:#909090;font-family:verdana,arial,sans-serif;font-size:10px;text-align:center;text-decoration:none;">
			   <!--
			    <table cellpadding="0" cellspacing="0" align="center" width="650">
			    	<tbody>
			            <tr>
			               <td><p style="border-bottom:1px solid; border-color:#909090; color:#909090;padding-bottom:3px">Hora do Remédio! </p></td>
			            </tr>
			        </tbody>
			    </table>
			    -->
					<table cellpadding="0" cellspacing="0" align="center" width="650">
						<tbody>
							<tr>
								<td>
			                    <table cellpadding="5" cellspacing="0" align="center" width="650">
			                    <tr> <td align="center">
										<img src="https://scontent-gru1-1.xx.fbcdn.net/hphotos-xpt1/v/t1.0-9/12316331_932005923542492_2407541284811252196_n.jpg?oh=5abd97a5ee6cf185dc31538127925644&oe=56E693B6" alt="Asvins" height="100" width="100" style="display:block;">
			                        </td></tr>
			                        </table>
								</td>
							</tr>
			                <tr>
											<td><p style="border-bottom:1px solid; border-color:#909090; color:#909090;padding-bottom:3px; font-size:20px;">Está na hora de tomar seu próximo pack! </p></td>
			            	</tr>
							<tr>
								<td>
			                    <table cellpadding="5" cellspacing="0" align="center" width="650">
			                    <tr> <td align="center">
										<img src="http://medcitynews.com/wp-content/uploads/PillPack_Packet.png" alt="Pack" height="400" width="650" style="display:block;">
			                        </td></tr>
			                        </table>
								</td>
							</tr>


						</tbody>
			</table>
				</body>
			</html>

			`,
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
