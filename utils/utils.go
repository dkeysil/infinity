package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 7,
	CheckRedirect: func(r *http.Request, via []*http.Request) error {
		return fmt.Errorf("target '%s' has returned a redirect", r.URL)
	},
}

func RandomHex(nbytes int) string {
	b := make([]byte, nbytes)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func JSONParse(jsonstr string) (interface{}, error) {
	var jsonvalue interface{}
	err := json.Unmarshal([]byte(jsonstr), &jsonvalue)
	return jsonvalue, err
}

func JSONEncode(jsonvalue interface{}) (string, error) {
	jsonb, err := json.Marshal(jsonvalue)
	if err != nil {
		return "", err
	}
	return string(jsonb), nil
}
