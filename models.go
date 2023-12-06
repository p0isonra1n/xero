package xero

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type GetEmployeesResponse struct {
	ID           string     `json:"Id"`
	Status       string     `json:"Status"`
	ProviderName string     `json:"ProviderName"`
	DateTimeUTC  string     `json:"DateTimeUTC"`
	Employees    []Employee `json:"Employees"`
}
type Employee struct {
	EmployeeID             string `json:"EmployeeID"`
	FirstName              string `json:"FirstName"`
	LastName               string `json:"LastName"`
	Status                 string `json:"Status"`
	Email                  string `json:"Email"`
	DateOfBirth            string `json:"DateOfBirth"`
	Gender                 string `json:"Gender"`
	Phone                  string `json:"Phone"`
	Mobile                 string `json:"Mobile,omitempty"`
	StartDate              string `json:"StartDate,omitempty"`
	OrdinaryEarningsRateID string `json:"OrdinaryEarningsRateID,omitempty"`
	PayrollCalendarID      string `json:"PayrollCalendarID,omitempty"`
	UpdatedDateUTC         string `json:"UpdatedDateUTC"`
	IsSTP2Qualified        any    `json:"IsSTP2Qualified"`
}
