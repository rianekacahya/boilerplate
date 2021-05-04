package repository

import "database/sql"

type authRepository struct {
	dbwrite *sql.DB
	dbread  *sql.DB
}

func NewAuthRepository(dbwrite, dbread *sql.DB) *authRepository {
	return &authRepository{dbwrite, dbread}
}
