package uniqueemails

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	uniqueEmails := map[string]string{}

	for _, email := range emails {
		if email == "" || !strings.Contains(email, "@") {
			continue
		}

		local := strings.Split(email, "@")
		name := strings.Split(strings.Replace(local[0], ".", "", -1), "+")[0]

		uniqueEmails[strings.Join([]string{name, local[1]}, "@")] = email
	}

	return len(uniqueEmails)
}
