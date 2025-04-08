package toursqry

import "github.com/mansio-gmbh/goapiutils/chrono"

type FromUntil struct {
	From  chrono.Date `json:"from"`
	Until chrono.Date `json:"until"`
}

type TourFilter struct {
	TenantID                 string    `json:"tenantID"`
	TourDate                 FromUntil `json:"tourDate"`
	PickupCountryCode        string    `json:"pickupCountryCode"`
	DeliveryCountryCode      string    `json:"deliveryCountryCode"`
	PickupPostalCodePrefix   string    `json:"pickupPostalCodePrefix"`
	DeliveryPostalCodePrefix string    `json:"deliveryPostalCodePrefix"`
	ScheduleID               string    `json:"scheduleID"`
}

type Filters []TourFilter
