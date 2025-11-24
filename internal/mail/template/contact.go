package template

import "fmt"

func EmailFormat(from, subject, body string) string {
	return fmt.Sprintf(
		`From: %s
To: %s
Subject: %s
MIME-Version: 1.0
Content-Type: text/plain; charset="UTF-8"

Email from: %s

%s
`, from, "work.jahnavimenon@gmail.com", subject, from, body)
}
