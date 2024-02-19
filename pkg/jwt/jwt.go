/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package jwt

import (
	"errors"
	"fmt"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"time"
)

const (
	issuer   = "northbit"
	audience = "unicare"
)

//jwt2.StandardClaims

func GenerateJWT(subject int, dataClaims map[string]interface{}) (string, error) {
	currentTime := time.Now().UTC()

	claims := jwt2.MapClaims{
		"iss": issuer,
		"sub": subject,
		"aud": audience,
		"iat": currentTime.Unix(),
		"jti": uuid.New().String(), // Generate a unique JWT ID using UUID v4.
		"exp": currentTime.Add(time.Hour * 6).Unix(),
	}

	for key, value := range dataClaims {
		claims[key] = value
	}

	token := jwt2.NewWithClaims(jwt2.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(viper.GetString("secret.jwt")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Verify(tokenString string) (bool, jwt2.MapClaims, error) {
	token, err := jwt2.Parse(tokenString, func(token *jwt2.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt2.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(viper.GetString("secret.jwt")), nil
	})

	if err != nil {
		return false, nil, fmt.Errorf("failed to parse JWT token: %v", err)
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt2.MapClaims); ok && token.Valid {
		if !claims.VerifyIssuer(issuer, true) {
			return false, nil, errors.New("invalid JWT token")
		}
		if !claims.VerifyAudience(audience, true) {
			return false, nil, errors.New("invalid JWT token")
		}
		return true, claims, nil
	}

	return false, nil, errors.New("invalid JWT token")
}
