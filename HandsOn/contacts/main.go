package main

import "contacts/model"

func main() {
	contact1 := model.Contact{
		Name:     "Bob Singh",
		Email:    "bob@example.com",
		Age:      35,
		Height:   172.2,
		IsActive: false,
	}

	contact2 := model.Contact{
		Name:     "Charlie Wu",
		Email:    "charlie@example.com",
		Age:      42,
		Height:   180.5,
		IsActive: true,
	}

	model.PrintDetails(contact1)
	model.PrintDetails(contact2)
}
