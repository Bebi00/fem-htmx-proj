package model

import "fmt"

var id = 0

type Contact struct {
	Id    int
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
	id++
	return Contact{
		Id:    id,
		Name:  name,
		Email: email,
	}
}

func NewContactList(n int) ContactList {
	contactsList := make([]Contact, n)
	for i := range n {
		id++
		contactsList[i] = Contact{
			Id:    id,
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

func (cl ContactList) IndexOf(id int) (int, error) {
	for i, contact := range cl.Contacts {
		if contact.Id == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Contact with id=%v not found", id)
}

func NewCreateContactData() CreateContactData {
	return CreateContactData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}
