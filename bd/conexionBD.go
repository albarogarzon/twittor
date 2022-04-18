package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN  */
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.ij2em.mongodb.net/twittor?retryWrites=true&w=majority")

/* ConectarBD conecta con db  */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("Conexion exitosa con la BD")
	return client
}

/*ChequeoConnection ping  */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
