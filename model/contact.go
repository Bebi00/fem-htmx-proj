package model

import "fmt"

type Contact struct {
	Name  string
	Email string
}

type ContactsList struct {
	Contacts []Contact
}

func NewContactList(n int) ContactsList {
	contactsList := ContactsList{
		make([]Contact, n),
	}
	for i := range n {
		contactsList.Contacts[i] = Contact{
			Name:  fmt.Sprintf("Person #%v", i),
			Email: fmt.Sprintf("person_%v@gmail.com", i),
		}
	}
	return contactsList
}
