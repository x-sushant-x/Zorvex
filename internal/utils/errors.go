package utils

import "errors"

var (
	ErrDBConnection             = errors.New("failed to connect to database")
	ErrDBTableCreate            = errors.New("failed to create database table")
	ErrDBInsert                 = errors.New("failed to insert to database")
	ErrDBGet                    = errors.New("failed to get data from database")
	ErrDataParse                = errors.New("failed to parse data from into struct")
	ErrServiceStatusChangeError = errors.New("failed to change service status in database")
	ErrDBTablesGet              = errors.New("failed to get all tables from database")

	ErrServiceNotFound     = errors.New("failed to find service with given name")
	ErrNoServiceAlive      = errors.New("failed to find any alive service")
	ErrUnableToLoadBalance = errors.New("failed to load balance")
)
