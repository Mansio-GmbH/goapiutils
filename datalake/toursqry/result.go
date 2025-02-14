package toursqry

import "github.com/mansio-gmbh/goapiutils/datalake/types"

type Result struct {
	Tours            []Tour                     `json:"tours"`
	Depots           []types.DepotDoc           `json:"depots"`
	HandoverStations []types.HandoverStationDoc `json:"handoverStations"`
}
