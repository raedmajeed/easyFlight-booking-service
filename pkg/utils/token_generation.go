package utils

import (
	"errors"
	"github.com/raedmajeed/booking-service/pkg/DOM"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/raedmajeed/booking-service/config"
)

type Claims struct {
	Email string
	Role  string
	jwt.StandardClaims
}

type SearchClaims struct {
	Adults        int
	Children      int
	Economy       bool
	PassengerType string
	jwt.StandardClaims
}

func GenerateToken(email, role string, cfg *config.ConfigParams) (string, error) {
	expireTime := time.Now().Add(time.Minute * 20).Unix()
	claims := &Claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Subject:   email,
			IssuedAt:  time.Now().Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(cfg.SECRETKEY))
	if err != nil {
		log.Printf("unable to generate jwt token for user %v, err: %v", email, err.Error())
		return "", err
	}

	return signedToken, nil
}

func GeneratePNRToken(pnr, role string, cfg *config.ConfigParams) (string, error) {
	expireTime := time.Now().Add(time.Minute * 20).Unix()
	claims := &Claims{
		Email: pnr,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Subject:   pnr,
			IssuedAt:  time.Now().Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(cfg.SECRETKEY))
	if err != nil {
		log.Printf("unable to generate jwt token for PNR %v, err: %v", pnr, err.Error())
		return "", err
	}

	return signedToken, nil
}

func GenerateSearchToken(info *DOM.AdditionalInfo, cfg *config.ConfigParams) (string, error) {
	expireTime := time.Now().Add(time.Minute * 55).Unix()
	claims := &SearchClaims{
		Adults:        info.AdultsCount,
		Children:      info.ChildrenCount,
		PassengerType: info.PassengerType,
		Economy:       info.Economy,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			IssuedAt:  time.Now().Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(cfg.SECRETKEY))
	if err != nil {
		log.Printf("unable to generate search token err: %v", err.Error())
		return "", err
	}

	return signedToken, nil
}

func ValidateSearchToken(token string, cfg config.ConfigParams) (*SearchClaims, error) {
	if token == "" {
		log.Print("search token missing")
		return &SearchClaims{}, errors.New("search token missing")
	}

	claims := &SearchClaims{}
	parserToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SECRETKEY), nil
	})

	if err != nil {
		log.Printf("error parsing token: %v", err)
		return &SearchClaims{}, errors.New("error parsing search token, token may have expired")
	}
	if !parserToken.Valid {
		log.Print("invalid token")
		return &SearchClaims{}, errors.New("search token is invalid invalid")
	}

	expTime := claims.ExpiresAt
	if expTime < time.Now().Unix() {
		log.Print("token Expired")
		return &SearchClaims{}, errors.New("search token is expired, please initiate a new flight search")
	}
	return claims, nil
}
