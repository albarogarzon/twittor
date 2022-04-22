package routers

import (
	"net/http"

	"github.com/albarogarzon/twittor/bd"
	"github.com/albarogarzon/twittor/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
