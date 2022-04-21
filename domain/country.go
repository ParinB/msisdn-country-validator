package domain

import "regexp"

type Country struct {
	Name string
	CountryCode string
	CountryCodeRegex *regexp.Regexp
	MsisdnRegex *regexp.Regexp
}
