package services

import (
	"errors"

	"goContextDiscovery/cmd/internal/core/domain"
	"goContextDiscovery/cmd/internal/core/ports"
)

type Port struct {
	personService ports.PersonService
}

//Function injection
func NewInstance(p ports.PersonService) *Port {
	return &Port{
		personService: p,
	}
}

func (s *Port) Find(id string) (domain.Person, error) {
	person, err := s.personService.Find(id)
	if err != nil {
		return person, errors.New("Fail get Person!")
	}
	return person, nil
}

func (s *Port) SignUp(name string, age int) error {
	err := s.personService.SignUp(name, age)
	if err != nil {
		return errors.New("Fail register Person!")
	}
	return nil
}
