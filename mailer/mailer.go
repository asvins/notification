package mailer

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"strings"
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

func Parse(opts map[string]string, raw string) string {
	result := raw
	for k, v := range opts {
		matching := fmt.Sprintf("{{%s}}", k)
		result = strings.Replace(result, matching, v, -1)
	}
	return result
}

//TEMPLATES
// 1) TOMAR PACK
const TemplatePackTime = `
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

`

// 2) COMPLETE SEUS DADOS
const TemplateFinishProfile = `
<!DOCTYPE html>
<html lang="pt-br">
	<head>
			<meta charset="utf-8">
					<title>Asvins - Complete Seu Cadastro</title>
						</head>
							<body style="color:#909090;font-family:verdana,arial,sans-serif;font-size:18px;text-align:center;text-decoration:none;">

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
																																																																																									               		<td><p style="border-bottom:1px solid; border-color:#909090; color:#909090;padding-bottom:3px;font-size:20px">Oi}! 
																																																																																																		               		<br><br> Falta alguns dados para sua assinatura ficar completa e você receber os medicamentos. 
																																																																																																											               		<br><br> Para isso basta seguir os seguintes passos:
																																																																																																																				               		<br>1) Acesse o sistema
																																																																																																																																		<br>2) Abra a aba "Contas"
																																																																																																																																							<br>3) Preencha os dados que faltam
																																																																																																																																												<br>4) Acesse o sistema
																																																																																																																																												               		</p></td>
																																																																																																																																																					            	</tr>
																																																																																																																																																												            				
																																																																																																																																																															</tbody>
																																																																																																																																																															</table>
																																																																																																																																																																</body>
																																																																																																																																																																</html>
`

// 3) BEM VINDO
const TemplateWelcome = `
<!DOCTYPE html>
<html lang="pt-br">
	<head>
			<meta charset="utf-8">
					<title>Asvins - Seja Bem-Vindo </title>
						</head>
							<body style="color:#909090;font-family:verdana,arial,sans-serif;font-size:18px;text-align:center;text-decoration:none;">

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
																																																																																									               		<td><p style="border-bottom:1px solid; border-color:#909090; color:#909090;padding-bottom:3px;font-size:20px">Oi! <br><br> Seja bem-vindo ao Sistema Asvins. <br> </p></td>
																																																																																																		            	</tr>
																																																																																																									            				
																																																																																																												</tbody>
																																																																																																												</table>
																																																																																																													</body>
																																																																																																													</html>

`
