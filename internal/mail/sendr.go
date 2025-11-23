package mail

import (
	"net/smtp"
	"os"

	//"strconv"
	"crypto/tls"
	"log"

	"github.com/jm-menon/Jahnavi-s-Portfolio/internal/mail/template"
)

//SMTP_HOST= smtp.gmail.com
//SMTP_PORT= 587
//SMTP_USERNAME= jahnavi.sends@gmail.com
//SMTP_PASSWORD= gingyissending
//ADMIN_EMAIL= work.jahnavimenon@gmail.com

func SendContact(from, subject, body string) error {
	to := os.Getenv("ADMIN_EMAIL")
	user := os.Getenv("SMTP_USERNAME")
	pass := os.Getenv("SMTP_PASSWORD")
	portStr := os.Getenv("SMTP_PORT")
	host := os.Getenv("SMTP_HOST")

	//port, _ := strconv.Atoi(portStr)
	msg := template.EmailFormat(from, subject, body)

	auth := smtp.PlainAuth("", user, pass, host)
	tlsConfig := &tls.Config{ServerName: host}

	conn, err := tls.Dial("tcp", host+":"+portStr, tlsConfig)
	if err != nil {
		log.Println("Error at connection: ", err)
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Println("Error at smtp client connect: ", err)
		return err
	}
	if err = client.Auth(auth); err != nil {
		return err
	}
	if err = client.Mail(user); err != nil {
		return err
	}
	if err = client.Rcpt(to); err != nil {
		return err
	}
	w, err := client.Data()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return client.Quit()

}
