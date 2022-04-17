package util

import (
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
}
