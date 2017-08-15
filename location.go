package outlook

type Location struct {
	DisplayName  string
	Address      PhysicalAddress
	Coordinates  GeoCoordinates
	EmailAddress string `json:"LocationEmailAddress"`
}
