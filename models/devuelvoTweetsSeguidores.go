package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoTweetsSeguidores es la estructura con la que devolveremos los tweets*/
type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitemty"`
	UsuarioID         string             `bson:"usuarioid" json:"userID,omitemty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"userRelationId,omitemty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitemty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitemty"`
		ID      string    `bson:"_id" json:"_id,omitemty"`
	}
}
