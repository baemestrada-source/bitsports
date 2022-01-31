package middlew

import (
	"errors"
	"strings"
	"time"
	"net/http"
	"github.com/graphql-go/handler"
	"github.com/baemestrada-source/bitsports/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var JwtSecret []byte = []byte("pruebatecnica")

/*ProcesoToken proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}
	return claims, false, string(""), err
}


func GeneroJWT(t models.Token) (string, error) {
	payload := jwt.MapClaims{
		"username": t.Username,
		"password": t.Password,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(JwtSecret)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			_, _, _, err := ProcesoToken(r.Header.Get("Authorization"))
			if err != nil {
				http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
				return
			}
			next.ServeHTTP(w, r)
		}
}

func HttpHeaderMiddleware(next *handler.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["token"]

		if !ok || len(keys[0]) < 1 {
			http.Error(w, "Error en el Token ! ", http.StatusBadRequest)
			return
		}
	
		token := keys[0]
	
		//log.Println("Token: " + string(token))

		_, _, _, err := ProcesoToken("Bearer "+token)
		if err != nil {
			http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}