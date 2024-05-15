package all

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"unicode"
)

func ValidatePassword(s string) (sevenOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}
	sevenOrMore = letters >= 7
	return
}

func ValidPassword(p string) (ok bool, err error) {

	sevenOrMore, number, upper, special := ValidatePassword(p)

	switch {
	case !sevenOrMore:
		errors.Join(err, ErrPassSevenOrMore)
	case !number:
		errors.Join(err, ErrPassNumber)
	case !upper:
		errors.Join(err, ErrPassUpper)
	case !special:
		errors.Join(err, ErrPassSpecial)
	default:
		ok = true
	}
	return
}

var (
	ErrPassSevenOrMore = errors.New("password must have seven or more letters")
	ErrPassNumber      = errors.New("password must have at least one number")
	ErrPassUpper       = errors.New("password must have at least one upercase letter")
	ErrPassSpecial     = errors.New("password must have at least one special character")
)

// Get public IP using ip-api
func GetPublicIP(c ...http.Client) (string, error) {
	var req *http.Response
	var err error
	if len(c) == 0 {
		req, err = http.Get("http://ip-api.com/json/")
	} else {
		req, err = c[0].Get("http://ip-api.com/json/")
	}

	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(io.Reader(req.Body))
	if err != nil {
		return "", err
	}

	var ip struct{ Query string }
	if err = json.Unmarshal(body, &ip); err != nil {
		return "", err
	}

	return ip.Query, nil
}
