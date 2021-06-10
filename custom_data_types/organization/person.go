package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Handler struct{
	handle string
	name string
}

// TwitterHandler comment
type TwitterHandler Handler

// RedirectURL comment
func (th TwitterHandler) RedirectURL() string {
	
	return ""
}

// Identifiable comment
type Identifiable interface {
	ID() string
}

// Person comment
type Person struct {
	firstName      string
	lastName       string
	twitterHandler TwitterHandler
}

// ID comment
func (p *Person) ID() string {
	return "12345"
}

// NewPerson comment
func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

// FullName comment
func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

// SetTwitterHandler comment
func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("Twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

// TwitterHandler comment
func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
