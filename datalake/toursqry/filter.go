package toursqry

type FromUntil struct {
	From  *string `json:"from"`
	Until *string `json:"until"`
}

type TourFilter struct {
	TenantID                 *string    `json:"tenantID"`
	TourDate                 *FromUntil `json:"tourDate"`
	PickupCountryCode        *string    `json:"pickupCountryCode"`
	DeliveryCountryCode      *string    `json:"deliveryCountryCode"`
	PickupPostalCodePrefix   *string    `json:"pickupPostalCodePrefix"`
	DeliveryPostalCodePrefix *string    `json:"deliveryPostalCodePrefix"`
}

type Filters []TourFilter
