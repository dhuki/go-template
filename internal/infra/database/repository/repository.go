package repository

import "github.com/jmoiron/sqlx"

//go:generate mockgen -destination=mocks/mock_repo.go -package=mocks github.com/dhuki/go-template/internal/infra/database/repository IRepository

type IRepository interface {
	Transaction
	HealthRepository
}

type Repository struct {
	dbMaster, dbSlave *sqlx.DB
}

func NewRepository(dbMaster, dbSlave *sqlx.DB) IRepository {
	return Repository{
		dbMaster: dbMaster,
		dbSlave:  dbSlave,
	}
}
