package db

import (
	"errors"
	"time"

	"github.com/sushant102004/zorvex/internal/types"
	"github.com/sushant102004/zorvex/internal/utils"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type DBClient interface {
	AddNewServiceToDB(types.Service) error
	CreateTables() error
	GetServiceInstances(string) ([]types.Service, error)
	GetAllServices() ([]types.Service, error)
	ChangeServiceStatus(id, status string) error
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
		return nil, errors.Join(utils.ErrDBConnection, err)
	}

	return &RethinkClient{
		Session: session,
		DB:      rethinkdb.DB("zorvex"),
	}, nil
}

func (r *RethinkClient) CreateTables() error {
	_, err := r.DB.TableCreate("services").Run(r.Session)
	if err != nil {
		return errors.Join(utils.ErrDBTableCreate, err)
	}
	return nil
}

func (r *RethinkClient) AddNewServiceToDB(data types.Service) error {
	_, err := r.DB.Table("services").Insert(data).RunWrite(r.Session)
	if err != nil {
		return errors.Join(utils.ErrDBInsert, err)
	}

	return nil
}

func (r *RethinkClient) GetServiceInstances(name string) ([]types.Service, error) {
	cursor, err := r.DB.Table("services").Filter(map[string]interface{}{"name": name}).Run(r.Session)
	if err != nil {
		return nil, err
	}

	var result []types.Service
	err = cursor.All(&result)
	if err != nil {
		return nil, errors.Join(utils.ErrDBGet, err)
	}

	return result, nil
}

func (r *RethinkClient) GetAllServices() ([]types.Service, error) {
	cursor, err := r.DB.Table("services").Run(r.Session)
	if err != nil {
		return nil, errors.Join(utils.ErrDBGet, err)

	}

	var result []types.Service
	err = cursor.All(&result)
	if err != nil {
		return nil, errors.Join(utils.ErrDataParse, err)
	}

	return result, nil
}

func (r *RethinkClient) ChangeServiceStatus(id, status string) error {
	_, err := r.DB.Table("services").Update(
		map[string]interface{}{
			"status": status,
		},
		rethinkdb.UpdateOpts{},
	).Run(r.Session)

	if err != nil {
		return errors.Join(utils.ErrServiceStatusChangeError, err)
	}
	return nil
}
