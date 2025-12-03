package repository

type ServerStore interface {
	GetHost() string
	GetPort() int
}
