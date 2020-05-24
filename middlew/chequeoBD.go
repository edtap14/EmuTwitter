package middlew

import (
	"EmuTwitter/bd"
	"net/http"
)

/*ChequeoBD funcion middelware para conocer el estado de la base de datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdidad con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
