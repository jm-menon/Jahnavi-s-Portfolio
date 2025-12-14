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
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func SendContact(userEmail, subject, body string) error {
	ctx := context.Background()

	clientID := os.Getenv("GMAIL_CLIENT_ID")
	clientSecret := os.Getenv("GMAIL_CLIENT_SECRET")
	refreshToken := os.Getenv("GMAIL_REFRESH_TOKEN")
	admin := os.Getenv("ADMIN_EMAIL")

	if clientID == "" || clientSecret == "" || refreshToken == "" || admin == "" {
		return fmt.Errorf("missing gmail oauth environment variables")
	}

	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/gmail.send"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/v2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	client := conf.Client(ctx, token)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("gmail service error: %w", err)
	}

	rawMessage := strings.Join([]string{
		"From: " + admin,
		"To: " + admin,
		"Subject: " + subject,
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=\"UTF-8\"",
		"",
		"New contact form submission",
		"",
		"User email: " + userEmail,
		"",
		"Message:",
		body,
	}, "\r\n")

	msg := &gmail.Message{
		Raw: base64.RawURLEncoding.EncodeToString([]byte(rawMessage)),
	}

	resp, err := srv.Users.Messages.Send("me", msg).Do()
	if err != nil {
		return fmt.Errorf("gmail send failed: %w", err)
	}

	log.Println("Email sent successfully, message ID:", resp.Id)
	return nil
}
