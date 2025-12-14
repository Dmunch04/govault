package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"os"
)

var (
	clientID 	 = os.Getenv("OIDC_CLIENT_ID")
	clientSecret = os.Getenv("OIDC_CLIENT_SECRET")
)

func randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func setCallbackCookie(w http.ResponseWriter, r *http.Request, name string, value string) {
	c :=
}