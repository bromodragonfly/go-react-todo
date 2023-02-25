package middleware

import("context" "github.com/joho/godotenv" "github.com/gorilla/mux" "encoding/json" "fmt" "net/http" "log" "go.mongodb.org/mongo-driver/bson" "go.mongodb.org/mongo-driver/bson/primitive" "go.mongodb.org/mongo-driver/mongo" "go.mongodb.org/mongo-driver/mongo/options"
"os")

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}
func loadTheEnv() {
  err :=	godotenv.Load(".env")
  if err!==nil{
	log.Fatal("Env error loading")
  }
}
func GetAllTasks(w http.ResponceWritter,r *http.Request) {

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



