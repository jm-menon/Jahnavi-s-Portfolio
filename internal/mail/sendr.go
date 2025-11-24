package mail

import (
	"net/smtp"
	"os"

	//"strconv"
	"crypto/tls"
	"log"

	"github.com/jm-menon/Jahnavi-s-Portfolio/internal/mail/template"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

//SMTP_HOST= smtp.gmail.com
//SMTP_PORT= 587
//SMTP_USERNAME= jahnavi.sends@gmail.com
//SMTP_PASSWORD= xxxxxxx
//ADMIN_EMAIL= work.jahnavimenon@gmail.com

func SendContact(from, subject, body string) error {
	to := os.Getenv("ADMIN_EMAIL")
	user := os.Getenv("SMTP_USERNAME")
	pass := os.Getenv("SMTP_PASSWORD")
	port := os.Getenv("SMTP_PORT")
	host := os.Getenv("SMTP_HOST")

	log.Println(to, " ,", user, " ,", pass, " ,", port, " ,", host)

	msg := []byte(template.EmailFormat(from, subject, body))

	addr := host + ":" + port

	// 1. Plain TCP connection
	conn, err := smtp.Dial(addr)
	if err != nil {
		log.Println("Dial error:", err)
		return err
	}
	defer conn.Close()

	// 2. STARTTLS Upgrade
	tlsConfig := &tls.Config{ServerName: host}
	if err = conn.StartTLS(tlsConfig); err != nil {
		log.Println("STARTTLS error:", err)
		return err
	}

	// 3. Authentication
	auth := smtp.PlainAuth("", user, pass, host)
	if err = conn.Auth(auth); err != nil {
		log.Println("Auth error:", err)
		return err
	}

	// 4. Envelope
	if err = conn.Mail(user); err != nil {
		log.Println("MAIL FROM error:", err)
		return err
	}
	if err = conn.Rcpt(to); err != nil {
		log.Println("RCPT TO error:", err)
		return err
	}

	// 5. Data
	w, err := conn.Data()
	if err != nil {
		log.Println("DATA begin error:", err)
		return err
	}

	if _, err = w.Write(msg); err != nil {
		log.Println("DATA write error:", err)
		return err
	}

	if err = w.Close(); err != nil {
		log.Println("DATA close error:", err)
		return err
	}

	return conn.Quit()
}
