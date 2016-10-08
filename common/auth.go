package common

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

//Ruta de los archivos crypto/RSA keys
const (
	//Key privada
	privKeyPath = "keys/app.rsa"
	//Key pública
	pubKeyPath = "key/app.rsa.pub"
)

//Key privada para firmar y key pública para verificar.
var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

//Claim Estructura que se tomará como base para crear el Token.
type Claims struct {
	*jwt.StandardClaims
	UserID uint
	RolID  uint
}

func initKeys() {
	var err error

	//Leemos la clave privada del archivo en disco
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys - privKeyPath]: %s\n", err)
	}
	//Parseamos signBytes de []byte a rsa.PrivateKey
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[ParseRSAPrivateKeyFromPEM]: %s\n", err)
	}

	//Leemos la clave pública del archivo en disco
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys - pubKeyPath]: %s\n", err)
		panic(err)
	}
	//Parseamos verifyBytes de []byte a rsa.PublicKey
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[ParseRSAPublicKeyFromPEM]: %s\n", err)
	}
}

//GenerateJWT Genera un toke encriptado con RS256
func GenerateJWT(claims *Claims) (string, error) {
	// Crear un signer para rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims = claims

	// Crear token string
	token, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return token, nil

}

//Authorize Middleware para validar tokens JWT
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequestWithClaims(r, request.AuthorizationHeaderExtractor, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //El tiempo del token expiró
				DisplayAppError(
					w,
					err,
					"Expiró en tiempo de vigencia del Token de Acceso!",
					401,
				)
				return
			default:
				DisplayAppError(
					w,
					err,
					"Error procesando su Token de Acceso!",
					500,
				)
				return
			}
		default:
			DisplayAppError(
				w,
				err,
				"Error procesando su Token de Acceso!",
				500,
			)
			return
		}
	}

	if token.Valid {
		next(w, r)
	} else {
		DisplayAppError(
			w,
			err,
			"Token de Acceso Invalido!",
			401,
		)
	}

}
