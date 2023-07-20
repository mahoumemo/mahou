package auth

import (
	"fmt"
	"net/http"

	"github.com/mahoumemo/mahou/misc"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": // stage 1
		w.Header().Add("X-DSi-Auth-Challenge", misc.RandString(8)) // TODO: find out how to validate these
		w.Header().Add("X-DSi-SID", misc.RandString(32))           // TODO: store this
		w.Header().Add("X-DSi-New-Notices", "0")                   // TODO: find out if this is really sent this early
		w.Header().Add("X-DSi-Unread-Notices", "0")                // TODO: same as above
	case "POST": // stage 2
		err := validateStage2AuthRequest(Stage2AuthRequest{
			mac:      r.Header.Get("X-DSi-MAC"),
			id:       r.Header.Get("X-DSi-ID"),
			auth:     r.Header.Get("X-DSi-Auth-Response"),
			sid:      r.Header.Get("X-DSi-SID"),
			version:  r.Header.Get("X-Ugomemo-Version"),
			username: r.Header.Get("X-DSi-User-Name"),
			region:   r.Header.Get("X-DSi-Region"),
			language: r.Header.Get("X-DSi-Lang"),
			country:  r.Header.Get("X-DSi-Country"),
			birthday: r.Header.Get("X-Birthday"),
			datetime: r.Header.Get("X-DSi-DateTime"),
			color:    r.Header.Get("X-DSi-Color"),
		})
		if err != nil {
			w.Header().Add("X-DSi-Dialog-Type", "1")
			w.Write(misc.StringToUTF16LE(fmt.Sprintf("Error while authenticating!\n\n\"%s\"", err)))

			return
		}

		w.Header().Add("X-DSi-SID", r.Header.Get("X-DSi-SID")) // send the same SID back if validated
		w.Header().Add("X-DSi-New-Notices", "0")               // TODO: look up if the user has new notices
		w.Header().Add("X-DSi-Unread-Notices", "0")            // TODO: same as above

		// TODO: friend list stuff
	default:
		http.Error(w, "unsupported method", http.StatusBadRequest)
	}
}

func validateStage2AuthRequest(req Stage2AuthRequest) error {
	if !isValidMAC(req.mac) {
		return errInvalidMAC
	}

	if !isValidFSID(req.id) {
		return errInvalidFSID
	}

	if !isValidAuthResponse(req.auth) {
		return errInvalidAuthResponse
	}

	if !isValidSID(req.sid) {
		return errInvalidSID
	}

	if !isValidVersion(req.version) {
		return errInvalidVersion
	}

	// TODO: handle usernames

	if !isValidRegion(req.region) {
		return errInvalidRegion
	}

	if !isValidLanguage(req.language) {
		return errInvalidLanguage
	}

	if !isValidCountry(req.country) {
		return errInvalidCountry
	}

	if !isValidBirthday(req.birthday) {
		return errInvalidBirthday
	}

	if !isValidDateTime(req.datetime) {
		return errInvalidDateTime
	}

	if !isValidColor(req.color) {
		return errInvalidColor
	}

	// TODO: database stuff

	return nil
}
