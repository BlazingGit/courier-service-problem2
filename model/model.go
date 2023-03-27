package model

type Coupon struct {
	OfferCode    string
	DiscountPerc int
	MinDistance  int
	MaxDistance  int
	MinWeight    int
	MaxWeight    int
}

type PackageDetail struct {
	PkgId        string
	PkgWeight    int
	Distance     int
	OfferCode    string
	DeliveryCost float64
	DeliveryTime float64
	Discount     float64
}

type PackageCombination struct {
	TotalWeight   int
	TotalDistance int
	PackageIDs    []string
}

type Vehicle struct {
	IsAvailable       bool
	DeliveryStartTime float64
}

type TestData struct {
	Input          []string
	ExpectedOutput []string
}
