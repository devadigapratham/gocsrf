package myJwt

import (
	"io/ioutil"
	"time"

	"github.com/devadigapratham/gocsrf/db/models"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
)

func InitJWT() error {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		return err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return err
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return err
	}

	verifyKey, err := jwt.ParseRSAPrivateKeyFromPEM(verifyBytes)
	if err != nil {
		return err
	}
	return nil
}

func CreateNewTokens(uuid string, role string) (authTokenString, refreshToken string, csrfSecret string, err error) {
	//generate csrf secret
	csrfSecret, err = models.GenerateCSRFSecret()
	if err != nil {
		return
	}
	//generate refresh token
	refreshTokenString, err = createRefreshTokenString(uuid, role, csrfSecret)
	//generating auth token
	authTokenString, err = createAuthTokenString(uuid, role, csrfSecret)
	if err != nil {
		return
	}
	return

}

func CreateAndRefreshTokens() {

}

func createAuthTokenString(uuid string, role string, csrfSecret string) (authTokenString, err error) {
	authTokenExp := time.Now().Add(models.AuthTokenValidTime).Unix()

}

func createRefreshTokenString() {

}

func updateRefreshTokenExpiry() {

}

func updateAuthTokenString() {

}

func RevokeRefreshToken() error {

}

func updateRefreshCsrf() {

}

func GrabUUID() {

}
