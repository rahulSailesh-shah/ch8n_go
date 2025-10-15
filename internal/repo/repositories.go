package repo

import "github.com/rahulSailesh-shah/ch8n_go/internal/db/sqlc"

type Repositories struct {
	User UserRepo
}

func NewRepositories(queries *sqlc.Queries) *Repositories {
	return &Repositories{
		User: NewUserRepo(queries),
	}
}
