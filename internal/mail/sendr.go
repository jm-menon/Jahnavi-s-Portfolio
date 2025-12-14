//replacing the old smtip code with the more practical oauth solution, as render blocks port 587
/*package mail

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
}*/

package mail

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	gmail "google.golang.org/api/gmail/v1"
)

func SendContact(from, subject, body string) error {
	clientID := os.Getenv("GMAIL_CLIENT_ID")
	clientSecret := os.Getenv("GMAIL_CLIENT_SECRET")
	refreshToken := os.Getenv("GMAIL_REFRESH_TOKEN")
	admin := os.Getenv("ADMIN_EMAIL")

	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{gmail.GmailSendScope},
		Endpoint:     google.Endpoint,
	}

	token := &oauth2.Token{RefreshToken: refreshToken}
	client := conf.Client(context.Background(), token)

	srv, err := gmail.New(client)
	if err != nil {
		return fmt.Errorf("gmail service error: %w", err)
	}

	msg := []byte(
		"From: " + from + "\r\n" +
			"To: " + admin + "\r\n" +
			"Subject: " + subject + "\r\n\r\n" +
			body,
	)

	var gmailMsg gmail.Message
	gmailMsg.Raw = base64.URLEncoding.EncodeToString(msg)

	_, err = srv.Users.Messages.Send("me", &gmailMsg).Do()
	return err
}
