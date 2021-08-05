package rest

import (
	"context"
	"net/http"
	"site/apibanco"
	"site/utils/log"

	"cloud.google.com/go/datastore"
)

func CalculoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		CalculaCode(w, r)
		return
	}
}

var c context.Context
var keys []*datastore.Key

func CalculaCode(w http.ResponseWriter, r *http.Request) int64 {
	code, err := apibanco.GetMultAPIBanco(c, keys)
	if err != nil {
		log.Warningf(c, "Erro ao buscar apis banco")
		return 0
	}

	var soma int64
	soma = 0
	for _, v := range code {
		if v.Code != 0 {
			soma += v.Code
		}
	}
	return soma
}
