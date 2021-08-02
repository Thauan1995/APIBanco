package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"site/apibanco"
	"site/utils"
	"site/utils/log"
	"strconv"
)

func APIBancoHandler(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	if r.Method == http.MethodGet {
		BuscaAPIBanco(w, r)
		return
	}

	if r.Method == http.MethodPost {
		InsereAPIBanco(w, r)
		return
	}

	log.Warningf(c, "Método não permitido")
	utils.RespondWithError(w, http.StatusMethodNotAllowed, 0, "Método não permitido")
	return
}

func BuscaAPIBanco(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	var (
		id   int64
		code int64
		err  error
	)
	if r.FormValue("ID") != "" {
		id, err = strconv.ParseInt(r.FormValue("ID"), 10, 64)
		if err != nil {
			log.Warningf(c, "Erro ao converter id pra string : %v", err)
			utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao converter id para string")
			return
		}
	}
	if r.FormValue("Code") != "" {
		code, err = strconv.ParseInt(r.FormValue("Code"), 10, 64)
		if err != nil {
			log.Warningf(c, "Erro ao converter code pra string : %v", err)
			utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao converter code para string")
			return
		}
	}

	filtro := apibanco.ApiBanco{
		ID:   id,
		Code: code,
		//Ispb: r.FormValue("Ispb"),
	}

	apiBanco, err := apibanco.FiltrarAPIBanco(c, filtro)
	if err != nil {
		log.Warningf(c, "Erro ao buscar API Banco %v", err)
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao buscar API Banco")
		return
	}

	log.Debugf(c, "Busca realizada com sucesso")
	utils.RespondWithJSON(w, http.StatusOK, apiBanco)
	return
}

func InsereAPIBanco(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	var apiBanco = []apibanco.ApiBanco{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Warningf(c, "Erro ao receber body de APIBanco %v", err)
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao receber body de APIBanco")
		return
	}

	err = json.Unmarshal(body, &apiBanco)
	if err != nil {
		log.Warningf(c, "Erro ao realizar unmarshal de APIBanco %v", err)
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao realizar unmarshal de APIBanco")
		return
	}

	bancos := apiBanco
	for _, v := range bancos {
		if v.Name != "" {
			err = apibanco.InserirAPIBanco(c, &apibanco.ApiBanco{})
			if err != nil {
				log.Warningf(c, "Falha ao inserir slice de bancos %v", err)
				utils.RespondWithError(w, http.StatusBadRequest, 0, "Falha ao inserir slice de bancos")
				return
			}
		}
	}

	log.Debugf(c, "API Banco inserida com sucesso. %v", apiBanco)
	utils.RespondWithJSON(w, http.StatusOK, apiBanco)
	return
}
