package service

import "mygram-api/database"

type Service struct {
	db database.Database
}

func New(db database.Database) Service {
	return Service{
		db: db,
	}
}
