package routers

import (
	"EmuTwitter/bd"
	"EmuTwitter/models"
	"net/http"
)

/*AltaRelacion realiza el registro de al relacion entre usuarios*/
func AltaRelacion(w http.ResponseWriter, r *http.Request)  {
	ID := r.URL.Query().Get("id")
	if len(ID)<1{
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}
	var t models.Realacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil{
		http.Error(w, "Ocurrio un error al intentar insertar realción"+err.Error(),http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w,"No se ha logrado insertar la relación"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
