package database

import (
	"bytes"
	"strconv"
	"time"

	"github.com/dhuki/go-template/internal/infra/configloader"
	"github.com/dhuki/go-template/internal/infra/database/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func (p *connectionDBClient) NewPgRepository() (db repository.IRepository, err error) {
	dbMaster, err := p.openPostgresConnection(&p.conf.Master, &p.conf.DbConnectionInfo)
	if err != nil {
		return nil, err
	}
	dbSlave, err := p.openPostgresConnection(&p.conf.Slave, &p.conf.DbConnectionInfo)
	if err != nil {
		return nil, err
	}
	return repository.NewRepository(dbMaster, dbSlave), nil
}

func (p *connectionDBClient) openPostgresConnection(dbInfo *configloader.DBInfo, dbConnInfo *configloader.DbConnectionInfo) (*sqlx.DB, error) {
	var bufferStr bytes.Buffer
	bufferStr.WriteString(" host=" + dbInfo.Host)
	bufferStr.WriteString(" port=" + strconv.Itoa(dbInfo.Port))
	bufferStr.WriteString(" user=" + dbInfo.User)
	bufferStr.WriteString(" dbname=" + dbInfo.DBName)
	bufferStr.WriteString(" password=" + dbInfo.Password)
	bufferStr.WriteString(" sslmode=disable fallback_application_name=go-rest-example")
	connectionSource := bufferStr.String()

	db, err := sqlx.Connect("postgres", connectionSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(dbConnInfo.SetMaxIdleCons)
	db.SetMaxOpenConns(dbConnInfo.SetMaxOpenCons)
	db.SetConnMaxIdleTime(time.Duration(dbConnInfo.SetConMaxIdleTime) * time.Minute)
	db.SetConnMaxLifetime(time.Duration(dbConnInfo.SetConMaxLifetime) * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
