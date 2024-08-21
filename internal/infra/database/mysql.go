package database

import (
	"bytes"
	"fmt"
	"time"

	"github.com/dhuki/go-template/internal/infra/configloader"
	"github.com/dhuki/go-template/internal/infra/database/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func (p *connectionDBClient) NewMySQLRepository() (db repository.IRepository, err error) {
	dbMaster, err := p.openMysqlConnection(&p.conf.Master, &p.conf.DbConnectionInfo)
	if err != nil {
		return nil, err
	}
	dbSlave, err := p.openMysqlConnection(&p.conf.Slave, &p.conf.DbConnectionInfo)
	if err != nil {
		return nil, err
	}
	return repository.NewRepository(dbMaster, dbSlave), nil
}

func (p *connectionDBClient) openMysqlConnection(dbInfo *configloader.DBInfo, dbConnInfo *configloader.DbConnectionInfo) (*sqlx.DB, error) {
	var bufferStr bytes.Buffer
	// the driver will automatically parse these fields into Go's time.Time type.
	// Without parseTime=true, the MySQL driver will return DATE, DATETIME, TIMESTAMP, and TIME columns as strings
	bufferStr.WriteString("parseTime=" + "true")
	bufferStr.WriteString("&loc=" + "Asia%2FJakarta")
	opt := bufferStr.String()

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		dbInfo.User, dbInfo.Password,
		dbInfo.Host, dbInfo.Port,
		dbInfo.DBName,
		opt,
	))
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
