package ports

import "goContextDiscovery/cmd/internal/core/domain"

type PersonService interface {
	Find(id string) (domain.Person, error)
	SignUp(name string, age int) error
}
