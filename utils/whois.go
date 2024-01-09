package utils

import (
	"github.com/likexian/whois"
)




func ConsultDomain(host string) string{
	result, _:= whois.Whois(host)
	
	return result
}