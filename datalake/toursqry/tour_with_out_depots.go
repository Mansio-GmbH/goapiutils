package toursqry

import "github.com/mansio-gmbh/goapiutils/datalake/types"

type DepotsAssignment struct {
	OutDepotIDs       []string    `json:"outDepotIDs"`
	InDepotIDs        []string    `json:"inDepotIDs"`
	OwnerDepotIDs     []string    `json:"ownerDepotsIDs"`
	PickupDistances   [][]float64 `json:"pickupDistances"`
	PickupDurations   [][]int     `json:"pickupDurations"`
	DeliveryDistance  [][]float64 `json:"deliveryDistance"`
	DeliveryDurations [][]int     `json:"deliveryDurations"`
}

type Tour struct {
	types.TourDoc
	DepotsAssignment DepotsAssignment `json:"depotsAssignment"`
}
