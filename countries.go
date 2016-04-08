// Package countries implements a simple country name/code parser.
package countries

import (
	"fmt"
	"strings"
)

//go:generate go run gen_countries.go

type Country struct {
	// The country name, in English
	Name string
	// ISO 3166-1 ALPHA-2 code
	ISO2 string
	// ISO 3166-1 ALPHA-3 code
	ISO3 string
}

// Parse parses a country code or name. Only English names are supported
// right now. Supported codes are: ISO2, ISO3 and FIPS.
func Parse(nameOrCode string) (*Country, error) {
	upper := strings.ToUpper(strings.TrimSpace(nameOrCode))
	if _, ok := iso2Codes[upper]; ok {
		return fromISOCode(upper)
	}
	if code, ok := iso2ByName[upper]; ok {
		return fromISOCode(code)
	}
	if code, ok := iso2ByIso3[upper]; ok {
		return fromISOCode(code)
	}
	if code, ok := iso2ByFips[upper]; ok {
		return fromISOCode(code)
	}
	return nil, fmt.Errorf("country %q not found", nameOrCode)
}

func fromISOCode(code string) (*Country, error) {
	if c, ok := countriesByIso2[code]; ok {
		return &c, nil
	}
	return nil, fmt.Errorf("country with ISO2 code %q not found", code)
}
