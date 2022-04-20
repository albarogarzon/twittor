package bd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN  */
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI(goDotEnvVariable("MONGO_URI"))

/* ConectarBD conecta con db  */
func ConectarBD() *mongo.Client {
	fmt.Println("Conexion: ",clientOptions.GetURI())
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

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	if os.Getenv("APP_ENV") == "heroku" {
		return os.Getenv(key)
	}

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}
	return os.Getenv(key)

}
