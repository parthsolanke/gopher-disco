package main

import (
	"fmt"
	"slices"
)

type Activity int

const (
	sprint Activity = iota
	jog
	gym
	cardio
)

func (a Activity) String() string {
	switch a {
	case sprint:
		return "sprint"
	case jog:
		return "jog"
	case gym:
		return "gym"
	case cardio:
		return "cardio"
	default:
		return "unknown"
	}
}

type user struct {
	name       string
	email      string
	activities []Activity
}

func newUser(name string, email string) *user {
	p := user{name: name, email: email}
	return &p
}

func (u *user) trackActivity(act Activity) {
	u.activities = append(u.activities, act)
	fmt.Println()
	fmt.Printf("%v added for user: %s\n", act, u.name)
}

type allUsers struct {
	users []*user
}

func (au *allUsers) getAllEmails() []string {
	allEmails := make([]string, 0, 10)
	for _, usr := range au.users {
		allEmails = append(allEmails, usr.email)
	}
	return allEmails
}

func (au *allUsers) addUser(u *user) {
	if emails := au.getAllEmails(); slices.Contains(emails, u.email) {
		fmt.Println("User already exists")
		return
	}
	au.users = append(au.users, u)
	fmt.Println("User Added Successfully")
}

func (au *allUsers) getSummary() {
	fmt.Println("\n---Summary---")
	for _, usr := range au.users {
		fmt.Println("name: ", usr.name)
		fmt.Println("email: ", usr.email)
		fmt.Println("Activities: ", usr.activities)
		fmt.Println()
	}
}

func main() {
	var users allUsers

	user1 := newUser("parth", "parth@example.com")
	user2 := newUser("Bob", "bob@example.com")

	users.addUser(user1)
	users.addUser(user2)

	duplicateUser := newUser("parth2", "parth@example.com")
	users.addUser(duplicateUser)

	user1.trackActivity(sprint)
	user1.trackActivity(gym)

	user2.trackActivity(jog)
	user2.trackActivity(cardio)

	users.getSummary()
}
