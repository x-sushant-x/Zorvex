package observer

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/sushant102004/zorvex/internal/db"
	"github.com/sushant102004/zorvex/internal/types"
)

type Observer struct {
	db                *db.RethinkClient
	ServicesInstances map[string][]types.Service
	// This map will store pointers for Round Robin Load Balancing
	ServicesPointers map[string]int
}

func NewObserver(db *db.RethinkClient) *Observer {
	return &Observer{
		db:                db,
		ServicesInstances: make(map[string][]types.Service),
		ServicesPointers:  make(map[string]int),
	}
}

func (o *Observer) SetupAllServicesOnStart() {
	// This function will get all the already registered services from database
	// and store them to o.ServicesInstances
	services, err := o.db.GetAllServices()
	if err != nil {
		log.Error().Msgf("unable to fetch all services from database: %v", err.Error())
	}

	for _, service := range services {
		o.ServicesInstances[service.Name] = append(o.ServicesInstances[service.Name], service)
		log.Info().Msgf("Appended Service: %s", service.Name)
	}
}

func (o *Observer) StreamInstances() {
	cursor, err := o.db.DB.Table("services").Changes().Run(o.db.Session)
	if err != nil {
		log.Error().Msgf("Error while streaming instances: %v", err)
		return
	}

	var change map[string]interface{}
	var data types.Service

	for cursor.Next(&change) {
		// Checking if a new service is added to database.
		if change["new_val"] != nil {
			newDoc := change["new_val"].(map[string]interface{})
			encoded, err := json.Marshal(newDoc)
			if err != nil {
				log.Info().Msgf("unable to marshal service data: %v", err.Error())
			}
			err = json.Unmarshal(encoded, &data)
			if err != nil {
				log.Info().Msgf("unable to unmarshal service data: %v", err.Error())
			}
			log.Info().Msgf("Appended Service: %s", data.Name)
			o.ServicesInstances[data.Name] = append(o.ServicesInstances[data.Name], data)
		} else if change["new_val"] == nil && change["old_val"] != nil {
			tempInstances := []types.Service{}
			oldDoc := change["old_val"].(map[string]interface{})

			encoded, err := json.Marshal(oldDoc)
			if err != nil {
				log.Info().Msgf("unable to marshal service data: %v", err.Error())
			}
			err = json.Unmarshal(encoded, &data)
			if err != nil {
				log.Info().Msgf("unable to unmarshal service data: %v", err.Error())
			}

			for _, services := range o.ServicesInstances {
				for _, service := range services {
					if service.ID != data.ID {
						tempInstances = append(tempInstances, service)
					}
				}
			}
			o.ServicesInstances[data.Name] = tempInstances
			log.Info().Msgf("Removed Service: %s", data.Name)
		}
	}
}

func (o *Observer) GetServiceFromObserver(name string) ([]types.Service, error) {
	return o.ServicesInstances[name], nil
}
