package tokyovacapi

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

func verifyJWT(token string) error {
	decodedToken, err := DecodeJWT(token)
	if err != nil {
		return err
	}
	if decodedToken.Exp < int(time.Now().Unix()) {
		return ErrTokenExpired
	}
	return nil
}

func trimMap(data map[string]string) map[string]string {
	result := map[string]string{}
	for k, v := range data {
		if v != "" {
			result[k] = v
		}
	}
	return result
}

func DecodeJWT(token string) (*Token, error) {
	splited := strings.Split(token, ".")
	if len(splited) != 3 {
		return nil, ErrTokenInvalid
	}
	data, err := base64.RawStdEncoding.DecodeString(splited[1])
	if err != nil {
		return nil, ErrTokenInvalid
	}
	decodedToken := &Token{}
	if err := json.Unmarshal(data, decodedToken); err != nil {
		return nil, ErrTokenInvalid
	}
	return decodedToken, nil
}

func toError(in *Error) (error, bool) {
	if in == nil {
		return nil, true
	}
	if in.Code == "token_not_valid" {
		return ErrTokenInvalid, true
	}
	if in.NonFieldErrors == "No more reservation available" {
		return ErrReservationUnavailable, true
	}
	if in.Detail != "" {
		return errors.New(in.Detail), true
	}
	if in.NonFieldErrors != "" {
		return errors.New(in.NonFieldErrors), true
	}
	return nil, false
}
