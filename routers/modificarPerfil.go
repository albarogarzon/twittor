package routers

import (
	"encoding/json"
	"net/http"

	"github.com/albarogarzon/twittor/bd"
	"github.com/albarogarzon/twittor/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	status, err := bd.ModificoRegistro(t, IDUsuario) //IDUsuario es la variable global creada en procesoToken.go

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro."+err.Error(), 400)
		return
	}
	if !status{
		http.Error(w, "No se ha logrado modificar el usuario.", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
