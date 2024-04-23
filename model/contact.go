package model

import "fmt"

type Contact struct {
	Name  string
	Email string
}

type ContactList struct {
	Contacts []Contact
}

type CreateContactData struct {
	Values map[string]string
	Errors map[string]string
}

func NewContact(name, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

func NewContactList(n int) ContactList {
	contactsList := make([]Contact, n)
	for i := range n {
		contactsList[i] = Contact{
			Name:  fmt.Sprintf("Person #%v", i),
			Email: fmt.Sprintf("person_%v@gmail.com", i),
		}
	}
	return ContactList{Contacts: contactsList}
}

func ExistsEmail(email string, cl []Contact) bool {
	for _, contact := range cl {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func NewCreateContactData() CreateContactData {
	return CreateContactData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}
