package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	visited := make(map[string]struct{})

	for _, email := range emails {
		parts := strings.Split(email, "@")
		local, domain := parts[0], parts[1]

		// ignore everything after "+"
		if plusIndex := strings.Index(local, "+"); plusIndex != -1 {
			local = local[:plusIndex]
		}

		// remove all "."
		local = strings.ReplaceAll(local, ".", "")

		normalizedEmail := local + "@" + domain
		visited[normalizedEmail] = struct{}{}
	}

	return len(visited)
}

func main() {
	emails := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}
	fmt.Println(numUniqueEmails(emails))
}
