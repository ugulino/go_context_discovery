package repositories

import (
	"encoding/json"
	"errors"

	"goContextDiscovery/cmd/internal/core/domain"
	"goContextDiscovery/cmd/internal/core/ports"

	uuid "github.com/google/uuid"
)

type memkvc struct {
	kvs map[string][]byte
}

func NewMemKVS() ports.PersonService {
	return &memkvc{
		kvs: map[string][]byte{},
	}
}

func (repo *memkvc) Find(id string) (domain.Person, error) {
	if value, ok := repo.kvs[id]; ok {
		person := domain.Person{}
		err := json.Unmarshal(value, &person)
		if err != nil {
			return domain.Person{}, errors.New("Fail get person from KVS!")
		}

		return person, nil
	}
	return domain.Person{}, errors.New("Person not found in KVS!")
}

func (repo *memkvc) SignUp(name string, age int) error {
	person := domain.Person{
		Id:   uuid.NewString(),
		Name: name,
		Age:  age,
	}

	value, err := json.Marshal(person)
	if err != nil {
		return errors.New("Fail save person from KVS!")
	}
	repo.kvs[person.Id] = value
	return nil
}
