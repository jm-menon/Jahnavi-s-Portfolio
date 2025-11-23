package template

import (
	"fmt"
	"os"
)

func EmailFormat(from, subject, message string) string {
	user := os.Getenv("SMTP_USERNAME")
	return fmt.Sprintf(`From: %s
To: %s
Subject: Portfolio Contact: %s

Email: %s
Message:
%s`, from, user, subject, from, message)
}
