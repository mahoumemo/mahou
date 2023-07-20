package auth

import (
	"regexp"

	"github.com/mahoumemo/mahou/misc"
)

var (
	isValidMAC          = regexp.MustCompile("^[A-F0-9]{16}$").MatchString
	isValidFSID         = regexp.MustCompile("^[0159]{1}[0-9A-F]{15}$").MatchString
	isValidAuthResponse = regexp.MustCompile("^[a-f0-9]{8}$").MatchString
	isValidSID          = regexp.MustCompile("^[" + misc.RandStringChars + "]{32}$").MatchString
	isValidVersion      = regexp.MustCompile("^[0-2]$").MatchString
	// TODO: isValidUsername
	isValidRegion   = regexp.MustCompile("^[0-2]$").MatchString
	isValidLanguage = regexp.MustCompile("^[a-z]{2}$").MatchString
	isValidCountry  = regexp.MustCompile("^[A-Z]{2}$").MatchString
	isValidBirthday = regexp.MustCompile("^[0-9]{8}$").MatchString
	isValidDateTime = regexp.MustCompile("^[0-9]{4}-[0-9]{2}-[0-9]{2}_[0-9]{2}:[0-9]{2}:[0-9]{2}$").MatchString
	isValidColor    = regexp.MustCompile("^[a-f]{6}$").MatchString
)
