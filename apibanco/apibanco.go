package apibanco

import (
	"context"
	"fmt"
	"site/utils/consts"
	"site/utils/log"

	"cloud.google.com/go/datastore"
)

const (
	KindBanco = "Banco"
)

type ApiBanco struct {
	ID       int64 `datastore:"-"`
	Ispb     string
	Name     string
	Code     int64
	FullName string
}

func GetAPIBanco(c context.Context, id int64) *ApiBanco {
	datastoreClient, err := datastore.NewClient(c, consts.IDProjeto)
	if err != nil {
		log.Warningf(c, "Erro ao conectar-se com o Datastore: %v", err)
		return nil
	}
	defer datastoreClient.Close()

	key := datastore.IDKey(KindBanco, id, nil)

	var apibanco ApiBanco
	if err = datastoreClient.Get(c, key, &apibanco); err != nil {
		log.Warningf(c, "Erro ao buscar API Banco: %v", err)
		return nil
	}

	apibanco.ID = id
	return &apibanco
}

func GetMultAPIBanco(c context.Context, keys []*datastore.Key) ([]ApiBanco, error) {
	datastoreClient, err := datastore.NewClient(c, consts.IDProjeto)
	if err != nil {
		log.Warningf(c, "Erro ao conectar-se com o Datastore: %v", err)
		return []ApiBanco{}, err
	}
	defer datastoreClient.Close()

	apiBanco := make([]ApiBanco, len(keys))
	if err := datastoreClient.GetMulti(c, keys, apiBanco); err != nil {
		if errs, ok := err.(datastore.MultiError); ok {
			for _, e := range errs {
				if e == datastore.ErrNoSuchEntity {
					return []ApiBanco{}, nil
				}
			}
		}
		log.Warningf(c, "Erro ao buscar Multi API Bancos: %v", err)
		return []ApiBanco{}, err
	}
	for i := range keys {
		apiBanco[i].ID = keys[i].ID
	}
	return apiBanco, nil
}

func PutAPIBanco(c context.Context, apiBanco *ApiBanco) error {
	datastoreClient, err := datastore.NewClient(c, consts.IDProjeto)
	if err != nil {
		log.Warningf(c, "Erro ao conectar-se com o Datastore: %v", err)
		return err
	}
	defer datastoreClient.Close()

	key := datastore.IDKey(KindBanco, apiBanco.ID, nil)
	key, err = datastoreClient.Put(c, key, apiBanco)
	if err != nil {
		log.Warningf(c, "Erro ao inserir API Banco: %v", err)
		return err
	}
	apiBanco.ID = key.ID
	return nil
}
func PutMultAPIBanco(c context.Context, apiBanco []ApiBanco) error {
	if len(apiBanco) == 0 {
		return nil
	}
	datastoreClient, err := datastore.NewClient(c, consts.IDProjeto)
	if err != nil {
		log.Warningf(c, "Erro ao conectar-se com o Datastore: %v", err)
		return err
	}
	defer datastoreClient.Close()

	keys := make([]*datastore.Key, 0, len(apiBanco))
	for i := range apiBanco {
		keys = append(keys, datastore.IDKey(KindBanco, apiBanco[i].ID, nil))
	}
	keys, err = datastoreClient.PutMulti(c, keys, apiBanco)
	if err != nil {
		log.Warningf(c, "Erro ao inserir Multi API Banco: %v", err)
		return err
	}
	return nil
}
func InserirAPIBanco(c context.Context, apiBanco *ApiBanco) error {
	log.Debugf(c, "Inserindo API Banco: %v", apiBanco)

	if apiBanco.Code == 0 {
		return fmt.Errorf("Codigo do banco inválido: %v", apiBanco.Code)
	}

	if apiBanco.Ispb == "" {
		return fmt.Errorf("ISPB do banco inválida: %v", apiBanco.Ispb)
	}

	if apiBanco.Name == "" {
		return fmt.Errorf("Nome do banco inválido: %v", apiBanco.Name)
	}

	if apiBanco.FullName == "" {
		return fmt.Errorf("Nome completo do banco inválido: %v", apiBanco.FullName)
	}
	return PutMultAPIBanco(c, []ApiBanco{})
}

func FiltrarAPIBanco(c context.Context, apiBanco ApiBanco) ([]ApiBanco, error) {
	datastoreClient, err := datastore.NewClient(c, consts.IDProjeto)
	if err != nil {
		log.Warningf(c, "Erro ao conectar-se com o Datastore: %v", err)
		return nil, err
	}
	defer datastoreClient.Close()

	q := datastore.NewQuery(KindBanco)

	if apiBanco.Code != 0 {
		q = q.Filter("Code =", apiBanco.Code)
	}
	if apiBanco.Ispb != "" {
		q = q.Filter("Ispb =", apiBanco)
	}
	if apiBanco.Name != "" {
		q = q.Filter("Name =", apiBanco.Name)
	}
	if apiBanco.FullName != "" {
		q = q.Filter("FullName =", apiBanco.FullName)
	}
	if apiBanco.ID != 0 {
		key := datastore.IDKey(KindBanco, apiBanco.ID, nil)
		q = q.Filter("__key__ =", key)
	}

	q = q.KeysOnly()
	keys, err := datastoreClient.GetAll(c, q, nil)
	if err != nil {
		log.Warningf(c, "Erro ao buscar API Banco: %v", err)
		return nil, err
	}
	return GetMultAPIBanco(c, keys)
}
