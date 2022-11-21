package aggregate

import (
	"github.com/google/uuid"
	"ping.com/src/entity"
)

type Member struct {
	user *entity.User
}

func (m *Member) GetID() uuid.UUID {
	return m.user.Id
}

func (m *Member) SetID(id uuid.UUID) {
	if m.user == nil {
		m.user = &entity.User{}
	}
	m.user.Id = id
}
