package repository

type DbServicer interface {
	Ping() error
}
