package member

import "ping.com/src/aggregate"

type Repository interface {
	Register(member aggregate.Member)
	Unregister(member aggregate.Member)
}
