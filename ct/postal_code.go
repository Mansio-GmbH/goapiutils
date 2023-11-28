package ct

type PostalCode struct {
	PostalCode  string
	CountryCode string
}

func (pc PostalCode) String() string {
	return pc.CountryCode + pc.PostalCode
}
