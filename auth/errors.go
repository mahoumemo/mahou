package auth

import "errors"

var (
	errInvalidMAC          = errors.New("invalid mac")
	errInvalidFSID         = errors.New("invalid fsid")
	errInvalidAuthResponse = errors.New("invalid auth response")
	errInvalidSID          = errors.New("invalid sid")
	errInvalidVersion      = errors.New("invalid version")
	errInvalidUsername     = errors.New("invalid username")
	errInvalidRegion       = errors.New("invalid region")
	errInvalidLanguage     = errors.New("invalid language")
	errInvalidCountry      = errors.New("invalid country")
	errInvalidBirthday     = errors.New("invalid birthday")
	errInvalidDateTime     = errors.New("invalid datetime")
	errInvalidColor        = errors.New("invalid color")
)
