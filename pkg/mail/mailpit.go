package mail

import (
	"strconv"

	"github.com/DestCom-Student-Projects/GO-PROJECT-1/pkg/environnement"
	"gopkg.in/mail.v2"
)


func NewEmailClient() *mail.Dialer {
    mailIP := environnement.GetEnv("MAIL_SERVER")
    mailPortStr := environnement.GetEnv("MAIL_PORT")
    mailPort, _ := strconv.Atoi(mailPortStr)

    dialer := mail.NewDialer(mailIP, mailPort, "", "")
    return dialer
}

func SendMail(to, subject, body, filepath string) error {
	e := NewEmailClient()

    mailSender:= environnement.GetEnv("MAIL_SENDER")

    msg := mail.NewMessage()
    msg.SetHeader("From", mailSender)
    msg.SetHeader("To", to)
    msg.SetHeader("Subject", subject)
    msg.SetBody("text/plain", body)
    if (len(filepath) > 0) {
        msg.Attach(filepath)
    }

    return e.DialAndSend(msg)
}
	

//AZEVEDO-DA-SILVA, A. (2024) GO-PROJECT-1. [Source code] https://github.com/DestCom-Student-Projects/GO-PROJECT-1