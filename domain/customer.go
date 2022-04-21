package domain

type Customer struct {
	Msisdn string `json:"msisdn"`
	Name string `json:"name"`
	Country string `json:"country"`
	CountryCode string `json:"country_code"`
	State string `json:"state"`
}


func  (c *Customer) Validate()  {
	for _, country := range Countries {
		ValidCode := country.CountryCodeRegex.MatchString(c.Msisdn)
		if ValidCode {
			c.CountryCode = country.CountryCode
			c.Country= country.Name
			ValidMsisdn := country.MsisdnRegex.MatchString(c.Msisdn)
			if ValidMsisdn == true{
				c.State = "VALID"
			}else {
				c.State = "NOT VALID"
			}
		}
	}

}