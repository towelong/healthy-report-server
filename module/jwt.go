package module

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	secretKey = "Thisismyapp"
	// hour
	accessTime = 2
)

type Claim struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func GenreateToken(userID int) (string, error) {
	myClaim := Claim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTime * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func VerifyToken(token string) (*Claim, error) {
	pc, err := jwt.ParseWithClaims(token, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !pc.Valid {
		return nil, err
	}
	if cs, ok := pc.Claims.(*Claim); ok {
		return cs, nil
	}
	return nil, nil
}
