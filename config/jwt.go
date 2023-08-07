package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("uYsvfGEYw2PMaGRLL0QNwUsWin7rxbEcRHNRWU2hXv6d_yqTPEi_jpl9sL8ADDRugcrlEu0IBzGewIOPh2fXMqB4F1rMbMaAi0Hlu8b6Wut1QGN0fkIJhb8IBVfqAcB9635nYi-Q_98zv_1yGM7AyxipTD2oGb0xu2fY9Pl6sOU")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
