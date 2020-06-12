package bd

import (
	"EmuTwitter/models"
	"context"
	"time"
)

/*InsertoRelacion graba la relacion en la BD*/
func InsertoRelacion(t models.Realacion) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	db := MongoCN.Database("Twittor")
	col := db.Collection("relacion")
	
	_,err := col.InsertOne(ctx, t)
	if err != nil{
		return false, err
	}

	return true, nil
}
