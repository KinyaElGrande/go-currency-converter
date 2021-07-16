package models

import (
	"net/url"
	"strings"
)

func (c *CurrencyConvert) Validate() url.Values {
	errs := url.Values{}

	//Check From request
	if !(strings.EqualFold(c.From, "ksh") || strings.EqualFold(c.From, "ghs") || strings.EqualFold(c.From, "ngn")) {
		errs.Add("from", "Sorry! Currencies can only be converted from (NGN), (GHS) and (KSH)")
	}

	//Check From request
	if !(strings.EqualFold(c.To, "ksh") || strings.EqualFold(c.To, "ghs") || strings.EqualFold(c.To, "ngn")) {
		errs.Add("to", "Sorry! Currencies can only be converted to (NGN), (GHS) and (KSH)")
	}

	return errs
}
