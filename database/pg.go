// Package database provides database support.
package database

import (
	"errors"

	"github.com/dronm/ds"
	"github.com/dronm/ds/pgds"

)

var DB *pgds.PgProvider

type DBStorage interface {
	GetPrimary() string;
	GetSecondaries() map[string]string;
}

func Initialize(storage DBStorage, onNotifFunc pgds.OnDbNotificationProto) error {
	//Db support
	dbProv, err := ds.NewProvider("pg", storage.GetPrimary(), onNotifFunc, storage.GetSecondaries())
	if err != nil {
		return err
	}
	ok := false
	DB, ok = dbProv.(*pgds.PgProvider)
	if !ok {
		return errors.New("dbProv is not of type *pgds.PgProvider")
	}

	if err := DB.Primary.Connect(); err != nil {
		return err
	}

	return nil
}


