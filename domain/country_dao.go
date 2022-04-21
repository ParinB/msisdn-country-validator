package domain

import (
	"github.com/parin/msisdn-country-validator/utils"
	"regexp"
)

var  (
	Countries = []Country{
		{
			Name : "Cameroon",
			CountryCode: "+237",
			CountryCodeRegex: regexp.MustCompile(`(?P<Code>\(237\))`),
			MsisdnRegex: regexp.MustCompile(`(?P<Code>\(237\)) ?[2368]\d{7,8}$`),
		},
		{
			Name:    "Ethiopia",
			CountryCode: "+251",
			CountryCodeRegex:  regexp.MustCompile(`(?P<Code>\(251\))`),
			MsisdnRegex : regexp.MustCompile(`(?P<Code>\(251\)) ?([1-59]\d{8})$`),
		},
		{
			Name: "Morocco",
			CountryCode: "+212",
			CountryCodeRegex:  regexp.MustCompile(`(?P<Code>\(212\))`),
			MsisdnRegex : regexp.MustCompile(`(?P<Code>\(212\)) ?([5-9]\d{8})$`),
		},
		{
			Name: "Mozambique",
			CountryCode: "+258",
			CountryCodeRegex :  regexp.MustCompile(`(?P<Code>\(258\))`),
			MsisdnRegex : regexp.MustCompile(`(?P<Code>\(258\)) ?([28]\d{7,8})$`),
		},
		{
			Name: "Uganda",
			CountryCode: "+256",
			CountryCodeRegex:  regexp.MustCompile(`(?P<Code>\(256\))`),
			MsisdnRegex: regexp.MustCompile(`(?P<Code>\(256\)) ?(\d{9})$`),
		},
	}
	CountryDao countryDao
)

type  countryDao struct {}

func (c countryDao) GetCountries() ([]Country,*utils.ApplicationError) {
	return Countries , nil
}