package jwt

import (
	"fmt"
	"log"
	"time"

	"Ejercicio_20/internal/config"
	"Ejercicio_20/internal/metrics"

	"github.com/golang-jwt/jwt/v5" // librería para crear/parsear/verificar JWT
)

// ParseToken acepta HS256 y RS256 durante la migración
func ParseToken(tokenStr string) (*jwt.Token, error) {
	parser := jwt.NewParser(jwt.WithValidMethods([]string{"HS256", "RS256"})) // restringe los alg válidos a HS256/RS256

	token, err := parser.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) { // parsea y pide una clave según el alg del token
		switch t.Method.Alg() {
		case "HS256":
			metrics.IncHS256()
			log.Println("[HS256] token validado con clave secreta")
			return config.HS256Secret, nil
		case "RS256":
			metrics.IncRS256()
			log.Println("[RS256] token validado con clave pública")
			return config.RSAPublic, nil
		default:
			return nil, fmt.Errorf("algoritmo no permitido: %s", t.Method.Alg())
		}
	})
	return token, err
}

// EmitToken genera un token en el algoritmo actual (RS256)
func EmitToken(user string) (string, error) {
	claims := jwt.MapClaims{
		"sub": user,
		"iat": time.Now().Unix(),                      //(momento de emisión)
		"exp": time.Now().Add(5 * time.Minute).Unix(), // (momento de expiración)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(config.RSAPrivate) // firma con la clave privada RSA y devuelve el string compacto
}
