package repository

import "context"

type HealthRepository interface {
	Ping(ctx context.Context) (err error)
}

func (u Repository) Ping(ctx context.Context) (err error) {
	if err = u.dbSlave.PingContext(ctx); err != nil {
		return
	}

	if err = u.dbMaster.PingContext(ctx); err != nil {
		return
	}
	return
}
