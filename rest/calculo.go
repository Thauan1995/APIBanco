package rest

import (
	"net/http"
	"site/apibanco"
	"site/utils"
	"site/utils/log"
)

func CalculoHandler(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	log.Warningf(c, "Inicializando soma dos codigos bancarios")
	calculo := apibanco.CalculaCode()
	if calculo == 0 {
		log.Warningf(c, "Erro ao fazer o calculo: %v", calculo)
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao fazer a soma dos codigos bancarios")
	}
}
