package models

type TerminationInquiry struct {
	AcquirerID string `json:"AcquirerId"`
	Merchant   struct {
		Name                string `json:"Name"`
		DoingBusinessAsName string `json:"DoingBusinessAsName"`
		Address             struct {
			Line1              string `json:"Line1"`
			Line2              string `json:"Line2"`
			City               string `json:"City"`
			CountrySubdivision string `json:"CountrySubdivision"`
			Province           string `json:"Province"`
			PostalCode         string `json:"PostalCode"`
			Country            string `json:"Country"`
		} `json:"Address"`
		PhoneNumber             string   `json:"PhoneNumber"`
		AltPhoneNumber          string   `json:"AltPhoneNumber"`
		NationalTaxID           string   `json:"NationalTaxId"`
		CountrySubdivisionTaxID string   `json:"CountrySubdivisionTaxId"`
		ServiceProvLegal        string   `json:"ServiceProvLegal"`
		ServiceProvDBA          string   `json:"ServiceProvDBA"`
		URL                     []string `json:"Url"`
		Principal               []struct {
			FirstName     string `json:"FirstName"`
			MiddleInitial string `json:"MiddleInitial"`
			LastName      string `json:"LastName"`
			Address       struct {
				Line1              string `json:"Line1"`
				Line2              string `json:"Line2"`
				City               string `json:"City"`
				CountrySubdivision string `json:"CountrySubdivision"`
				Province           string `json:"Province"`
				PostalCode         string `json:"PostalCode"`
				Country            string `json:"Country"`
			} `json:"Address"`
			PhoneNumber    string `json:"PhoneNumber"`
			AltPhoneNumber string `json:"AltPhoneNumber"`
			NationalID     string `json:"NationalId"`
			DriversLicense struct {
				Number             string `json:"Number"`
				CountrySubdivision string `json:"CountrySubdivision"`
				Country            string `json:"Country"`
			} `json:"DriversLicense"`
		} `json:"Principal"`
		SearchCriteria struct {
			SearchAll             string   `json:"SearchAll"`
			Region                []string `json:"Region"`
			Country               []string `json:"Country"`
			MinPossibleMatchCount string   `json:"MinPossibleMatchCount"`
		} `json:"SearchCriteria"`
		AddedOnDate           string `json:"AddedOnDate"`
		TerminationReasonCode string `json:"TerminationReasonCode"`
		AddedByAcquirerID     string `json:"AddedByAcquirerID"`
		URLGroup              []struct {
			ExactMatchUrls struct {
				URL []string `json:"Url"`
			} `json:"ExactMatchUrls"`
			CloseMatchUrls struct {
				URL []string `json:"Url"`
			} `json:"CloseMatchUrls"`
			NoMatchUrls struct {
				URL []string `json:"Url"`
			} `json:"NoMatchUrls"`
		} `json:"UrlGroup"`
	} `json:"Merchant"`
}
