package memory

import (
	"github.com/google/uuid"
	"ping.com/src/aggregate"
)

type InMemoryMemberRepository struct {
	members map[uuid.UUID]aggregate.Member
}

func New() *InMemoryMemberRepository {
	return &InMemoryMemberRepository{
		members: make(map[uuid.UUID]aggregate.Member),
	}
}

func (m *InMemoryMemberRepository) Register(member aggregate.Member) {
	if m.members == nil {
		m.members = make(map[uuid.UUID]aggregate.Member)
	}
	if _, ok := m.members[member.GetID()]; ok {
		return
	}
	m.members[member.GetID()] = member
}
