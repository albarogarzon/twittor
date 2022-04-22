package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/albarogarzon/twittor/bd"
)

func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/avatars/" + perfil.Avartar)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)

	if err != nil {
		http.Error(w, "Error al copiar imagen", http.StatusBadRequest)
		return
	}

}
