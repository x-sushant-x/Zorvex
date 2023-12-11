package db

import (
	"fmt"
	"time"

	"github.com/sushant102004/zorvex/internal/types"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type DBClient interface {
	AddNewServiceToDB(types.Service) error
	CreateTables() error
	GetServiceInstances(string) ([]types.Service, error)
}

type RethinkClient struct {
	Session *rethinkdb.Session
	DB      rethinkdb.Term
}

func NewRethinkClient() (*RethinkClient, error) {
	session, err := rethinkdb.Connect(rethinkdb.ConnectOpts{
		Address:  "localhost:28015",
		Database: "zorvex",
		Timeout:  time.Second * 5,
	})

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &RethinkClient{
		Session: session,
		DB:      rethinkdb.DB("zorvex"),
	}, nil
}

func (r *RethinkClient) CreateTables() error {
	_, err := r.DB.TableCreate("services").Run(r.Session)
	if err != nil {
		return fmt.Errorf("unable to create table: %v", err.Error())
	}

	fmt.Println("table created successfully")

	return nil
}

func (r *RethinkClient) AddNewServiceToDB(data types.Service) error {
	_, err := r.DB.Table("services").Insert(data).RunWrite(r.Session)
	if err != nil {
		return fmt.Errorf("unable to add service to db: %v", err.Error())
	}

	return nil
}

func (r *RethinkClient) GetServiceInstances(name string) ([]types.Service, error) {
	cursor, err := r.DB.Table("services").Filter(map[string]interface{}{"name": name}).Run(r.Session)
	if err != nil {
		return nil, fmt.Errorf("unable to add service to db: %v", err.Error())
	}

	var result []types.Service
	err = cursor.All(&result)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal services into json: %v", err.Error())
	}

	return result, nil
}
