package models

/*Realacion modelo para grabar la relacion de un usuario con otro */
type Realacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}
