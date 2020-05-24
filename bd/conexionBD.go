package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la base de datos */
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://edtap14:Americas1487@cluster0-umwsj.mongodb.net/test?retryWrites=true&w=majority")

/*ConectarBD es la funcion que me permite conectar la BD */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("conexion éxitosa con la BD")

	return client
}

/*ChequeoConnection es el ping a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
