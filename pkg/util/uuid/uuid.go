package uuid

import "github.com/google/uuid"

type UUID = uuid.UUID

func NewUUID() (UUID, error) {
	return uuid.NewRandom()
}

func MustUUID() UUID {
	v, err := NewUUID()
	if err != nil {
		panic(err)
	}
	return v
}

func MustString() string {
	return MustUUID().String()
}
