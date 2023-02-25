package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
);

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}
 
func loadTheEnv() {
  err :=	godotenv.Load(".env")
  if err!=nil{
	log.Fatal("Env error loading")
  }
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collecntionName := os.Getenv("COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err!=nil{
		log.Fatal(err)
	}
	
	err = client.Ping(context.TODO(), nil)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collecntionName)

	fmt.Println("Collection instance created")
	
}

func GetAllTasks(w http.ResponseWriter,r *http.Request) {

}

func Createtask() {

}

func TaskComplete(){
	
}

func UndoTask(){
	
}

func DeleteTask(){
	
}

func DeleteAllTasks(){
	
}



