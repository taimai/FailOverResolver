package resolver

import (
	"config"
	"datastore"
	"errors"
	"log"
	"math/rand"
	"time"
)

type Resolver struct {
	ConfigArray *[]*config.ServiceConfig
	Data        *datastore.DataStore
}

func (r *Resolver) Resolve(serviceName string) ([]string, error) {
	serviceConfExists := false
	for _, conf := range *r.ConfigArray {
		if conf.Servicename == serviceName {
			serviceConfExists = true
		}
	}

	if serviceConfExists {
		sd, err := r.Data.Get(serviceName)
		if sd != nil && err == nil {
			var servers []string
			for key, val := range sd.ServiceDataMap {
				if val.Pos > 0 && val.Queue[val.Pos-1] != nil && val.Queue[val.Pos-1].Serverstatus {
					servers = append(servers, key)
				}
			}
			if len(servers) > 0 {
				// fisher yates shuffle
				rand.Seed(time.Now().UnixNano())
				n := len(servers)
				for i := n - 1; i > 0; i-- {
					j := rand.Intn(i + 1)
					servers[i], servers[j] = servers[j], servers[i]
				}
				return servers, nil
			} else {
				return nil, errors.New("All servers down")
			}

		} else {
			log.Println("Service missing in datastore")
			return nil, errors.New("Service not found in datastore")
		}

	} else {
		log.Println("Error: no service " + serviceName)
		return nil, errors.New("Service not found in config")
	}

}
