package template

import (
	"fmt"
	"os"
)

func EmailFormat(from, subject, body string) string {
	return fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/plain; charset=\"UTF-8\"\r\n"+
			"\r\n"+
			"Email from: %s\r\n"+
			"%s\r\n",
		from,
		os.Getenv("ADMIN_EMAIL"),
		subject,
		from,
		body,
	)
}
