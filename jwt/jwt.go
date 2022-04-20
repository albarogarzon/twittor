package jwt

import (
	"time"

	"github.com/albarogarzon/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("MasterdelDesarrollo_grupodeFacebook")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_Nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioWeb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) //Header
	tokenStr, err := token.SignedString(miClave)                //Sign
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
