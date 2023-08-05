package auth

import (
	"github.com/WeebDeveloperz/titsunofficial-server/database"
	"os"
	"gorm.io/gorm"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var secret string = os.Getenv("JWT_SECRET")
var db *gorm.DB
func Init() {
	db = database.DB
	db.AutoMigrate(&User{})
}

type User struct {
	gorm.Model

	Username string `json:"username"`
	Password string `json:"password"`
	Disabled bool   `json:"disabled"`
	Role     string `json:"role"` // admin/peasant
}

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

func newJWT(username, role string) (string, error) {
	claims := Claims {
		jwt.RegisteredClaims {
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
		username, role,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}

func parseJWT(t string) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(t, &Claims{}, func (tk *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return token.Claims, err
}
