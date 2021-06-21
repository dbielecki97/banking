package domain

import (
	"encoding/json"
	"github.com/dbielecki97/banking-lib/logger"
	"github.com/dbielecki97/banking/dto"
	"net/http"
	"net/url"
)

type AuthRepository interface {
	IsAuthorized(string, string, map[string]string) bool
}

type RemoteAuthRepository struct {
}

func NewRemoteAuthRepository() *RemoteAuthRepository {
	return &RemoteAuthRepository{}
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	u := buildVerifyUrl(token, routeName, vars)

	if response, err := http.Get(u); err != nil {
		logger.Error("Error while sending..." + err.Error())
		return false
	} else {
		var ar dto.VerifyResponse
		if err := json.NewDecoder(response.Body).Decode(&ar); err != nil {
			logger.Error("Error while decoding response from auth server: " + err.Error())
			return false
		}
		return ar.IsAuthorized
	}
}

func buildVerifyUrl(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: "localhost:8033", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}

	u.RawQuery = q.Encode()
	return u.String()
}
