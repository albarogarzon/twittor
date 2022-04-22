package routers

import (
	"encoding/json"
	"net/http"

	"github.com/albarogarzon/twittor/bd"
	"github.com/albarogarzon/twittor/models"
)

//Consulta si hay o no relacion entre 2 usuarios
func ConsultoRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	/* 	if len(ID) < 1 {
	   		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
	   		return
	   	}
	*/
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)

	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
