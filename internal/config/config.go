package config

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v5" //para parsear claves PEM y trabajar con JWT.
)

// Habilitar la migración HS256 → RS256
var (
	// Centralizar la criptografía de tu app
	HS256Secret = []byte("supersecreto-temporal") // antiguo secreto simetrico para tokens legacy
	RSAPrivate  *rsa.PrivateKey                   //puntero a la clave privada RSA (se usa para firmar RS256).
	RSAPublic   *rsa.PublicKey                    //puntero a la clave pública RSA (se usa para verificar RS256).
)

func LoadKeys() {
	// Inicializar las claves al arrancar
	privBytes, _ := os.ReadFile("private.pem")
	pubBytes, _ := os.ReadFile("public.pem")

	privKey, _ := jwt.ParseRSAPrivateKeyFromPEM(privBytes) //decodifica PEM → DER → *rsa.PrivateKey (para firmar).
	pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(pubBytes)    //decodifica PEM → DER → *rsa.PrivateKey (para verificar).

	RSAPrivate = privKey //parseada
	RSAPublic = pubKey   // parseada , bytes crudos convertidos a estructura de go
}

//PEM significa Privacy-Enhanced Mail.
//Es un formato de texto estandarizado para almacenar y compartir claves, certificados o datos
//criptográficos (como los usados en HTTPS, JWT, SSH, etc.).
