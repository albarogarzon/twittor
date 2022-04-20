package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/albarogarzon/twittor/bd"
	"github.com/albarogarzon/twittor/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	regitro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(regitro)
	if err != nil {
		http.Error(w, "Error al insertar tweet"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se logro insertar tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
