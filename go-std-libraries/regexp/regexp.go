package regexp

import (
	"regexp"
	"fmt"
)

func DemoRegexpPackage() {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	compiled, _ := regexp.Compile(emailPattern)
	testEmails := []string{
		"test@example.com",
		"invalid-email",
		"user@domain.co",
		"user.name@domain.com",
		"user_name@domain.com",
		"user@sub.domain.com",
		"user@domain",
		"user@domain.c",
	}
	var validEmails []string
	var invalidEmails []string
	for _, email := range testEmails {
		if compiled.MatchString(email) {
			validEmails = append(validEmails, email)
		} else {
			invalidEmails = append(invalidEmails, email)
		}
	}
	fmt.Println("Valid emails:", validEmails)
	fmt.Println("Invalid emails:", invalidEmails)


	var namePattern = `^a[a-zA-Z]*[aiueo]$`
	nameCompiled, _ := regexp.Compile(namePattern)
	testNames := []string{
		"ana",
		"budi",
		"cici",
		"amelia",
		"abel",
		"amanda",
		"andi",
	}
	var matcherNames []string
	for _, name := range testNames {
		if nameCompiled.MatchString(name) {
			matcherNames = append(matcherNames, name)
		}
	}
	fmt.Println("Names matching pattern:", matcherNames)
}