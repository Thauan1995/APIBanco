package apibanco

import (
	"context"
	"site/utils/log"

	"cloud.google.com/go/datastore"
)

var c context.Context
var keys []*datastore.Key

func CalculaCode() int64 {
	code, err := GetMultAPIBanco(c, keys)
	if err != nil {
		log.Warningf(c, "Erro ao api banco")
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
