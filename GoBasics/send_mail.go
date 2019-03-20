package main

import (
	"net/mail"
	"fmt"
	"strconv"
	"net/smtp"
)

func main(){
	Send("Hello")
}
func Send(body string) error {

	//logger:= tempCtx.Loggers
	//mailCreds, respCode, tempError := reademailconf1(tempCtx)
	mailCreds := struct {
		Username string
		Password string
		SMTPHost string
		Port int
		To string
	}{
		Username: "mahesh.jadhav073@gmail.com",
		Password: "MJ@001801678803",
		SMTPHost: "smtp.gmail.com",
		Port: 587,
		To: "siddhantdoshi280592@gmail.com",
	}

	//// Check if server listens on that port.
	//if len(mailCreds.Username) == 0 && len(mailCreds.Password) == 0 {
	//	isAuthorized = false
	//	conn, err := smtp.Dial(mailCreds.SMTPHost + ":" + strconv.Itoa(mailCreds.Port))
	//	if err != nil {
	//		logger.Error("Could not setup smtp connection!")
	//		return err
	//	}
	//	client = conn
	//}

	//} else {
	//	isAuthorized = true
	//	conn, err := net.DialTimeout("tcp", mailCreds.SMTPHost + ":" + strconv.Itoa(mailCreds.Port), 3 * time.Second)
	//	if err != nil {
	//		return err
	//	}
	//	if conn != nil {
	//		defer conn.Close()
	//	}
	//}
	//// Validate sender and recipient
	_, err := mail.ParseAddress(mailCreds.Username)
	if err != nil {
		fmt.Println("Error: Sender could not be validated!")
		return err
	}
	_, err = mail.ParseAddress(mailCreds.To)
	if err != nil {
		fmt.Println("Error: Receiver could not be validated!")
		return err
	}

	msg := "From: " + mailCreds.Username + "\n" +
		"To: " + mailCreds.To + "\n" +
		"Subject: CLASSEC Server Error\n\n" +
		body
	host := mailCreds.SMTPHost + ":" + strconv.Itoa(mailCreds.Port)
	Err := smtp.SendMail(host,
		smtp.PlainAuth("", mailCreds.Username, mailCreds.Password, mailCreds.SMTPHost),
		mailCreds.Username, []string{mailCreds.To}, []byte(msg))

	if Err != nil {
		fmt.Errorf("smtp error: %s", Err)
		return Err
	}
	return nil
}


