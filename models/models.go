package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Address type
type Address struct {
	FullAddress   string `bson:"full_address,omitempty"`
	StreetAddress string `bson:"street_address,omitempty"`
	AddressLine2  string `bson:"addr_line2,omitempty"`
	City          string `bson:"city,omitempty"`
	State         string `bson:"state,omitempty"`
	Zip           string `bson:"zip,omitempty"`
}

// PropertyInfo type
type PropertyInfo struct {
	SquareFeet  int    `bson:"square_feet,omitempty"`
	Bedrooms    int    `bson:"bedrooms,omitempty"`
	Bathrooms   int    `bson:"bathrooms,omitempty"`
	Stories     int    `bson:"stories,omitempty"`
	Garage      string `bson:"garage,omitempty"`
	Attachments int    `bson:"attachments,omitempty"`
	Website     string `bson:"website,omitempty"`
	SchoolDist  string `bson:"school_district,omitempty"`
	Notes       string `bson:"property_notes,omitempty"`
}

// Value type
type Value struct {
	PurchaseDate  string  `bson:"purchase_date,omitempty"`
	PurchaseAmt   float64 `bson:"purchase_amt,omitempty"`
	SoldDate      string  `bson:"sold_date,omitempty"`
	SoldAmt       float64 `bson:"sold_amt,omitempty"`
	PctOwnership  float64 `bson:"pct_ownership,omitmpty"`
	AssessedValue float64 `bson:"assessed_value,omitempty"`
	PropertyTax   float64 `bson:"property_tax,omitempty"`
	CurrentValue  float64 `bson:"current_value,omitempty"`
	OwnedEquity   float64 `bson:"owned_equity,omitempty"`
	Notes         string  `bson:"value_notes,omitempty"`
}

// Mortgage type
type Mortgage struct {
	Bank          string  `bson:"bank,omitempty"`
	// LoanNumber    int64   `bson:"loan_number,omitempty"`
	MonthlyPmt    float64 `bson:"monthly_pmt,omitempty"`
	Interest      float64 `bson:"interest,omitempty"`
	MortgageTotal float64 `bson:"mortgage_total,omitempty"`
	Principal     float64 `bson:"principal,omitempty"`
	Notes         string  `bson:"mortgage_notes,omitempty"`
}

// Insurance type
type Insurance struct {
	Company       string  `bson:"insurance_comany,omitempty"`
	PolicyNumber  string  `bson:"policy_number,omitempty"`
	CoverageAmt   float64 `bson:"coverage_amt,omitempty"`
	EffectiveDate string  `bson:"effective_date,omitempty"`
	Premium       float64 `bson:"monthly_pmt,omitempty"`
	Notes         string  `bson:"insurance_notes,omitempty"`
}

// Lease type
type Lease struct {
	MonthlyPmt  float64 `bson:"lease_monthly,omitempty"`
	StartDate   string  `bson:"lease_start,omitempty"`
	EndDate     string  `bson:"lease_end,omitempty"`
	Leasee      string  `bson:"leasee,omitempty"`
	LeaseePhone string  `bson:"leasee_phone,omitempty"`
}

// Expenses type
type Expenses struct {
	HOADues     float64 `bson:"housing_dues,omitempty"`
	Electricity float64 `bson:"electric,omitempty"`
	Water       float64 `bson:"water,omitempty"`
	Garbage     float64 `bson:"garbage,omitempty"`
	WasteWater  float64 `bson:"wastewater,omitempty"`
	Gas         float64 `bson:"gas,omitempty"`
}

// CustomFields type // change to bool
type CustomFields struct {
	Recorded       string `bson:"recorded,omitempty"`
	HUDPurchase    string `bson:"hud_purchase,omitempty"`
	HUDSold        string `bson:"hud_sold,omitempty"`
	TitleCompanies struct {
		Purchase string `bson:"pur_title_co,omitempty"`
		Sale     string `bson:"sale_title_co,omitempty"`
	}
}

// Property type
type Property struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Nickname        string             `bson:"nickname,omitempty"`
	APN             string             `bson:"APN,omitempty"`
	Type            string             `bson:"type,omitempty"`
	Status          string             `bson:"status,omitempty"`
	Impound         string             `bson:"impound,omitempty"`
	Owner           string             `bson:"owned_by,omitempty"`
	PropertyManager string             `bson:"property_manager,omitempty"`
	Address         Address            `bson:"address,omitempty"`
	PropertyInfo    PropertyInfo       `bson:"property_info,omitempty"`
	Mortgage        Mortgage           `bson:"mortage,omitempty"`
	Insurance       Insurance          `bson:"insurance,omitempty"`
	Value           Value              `bson:"value,omitempty"`
	Lease           Lease              `bson:"lease,omitempty"`
	// Expenses        Expenses           `bson:"expenses,omitempty"`
	CustomFields CustomFields `bson:"custom_fields,omitempty"`
	Notes        string       `bson:"notes,omitempty"`
}
