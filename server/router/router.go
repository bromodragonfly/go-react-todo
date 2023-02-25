package router

import (
	"go-react-todo/middleware"

	"github.com/gorilla/mux"
)

// ?Это запись типо мы наследуемся от класса?
func Router() *mux.Router{

	router := mux.NewRouter()
	router.HandleFunc("api/tasks", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.Createtask).Methods("POST" , "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT" , "OPTIONS")
						// !в ундотаск метод PUT в видео, но я думаю подошел бы лучше UPDATE!
	router.HandleFunc("/api/undoTusk/{id}", middleware.UndoTask).Methods("PUT" , "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE" , "OPTIONS")
	router.HandleFunc("/api/deleteAllTasks}", middleware.DeleteAllTasks).Methods("DELETE" , "OPTIONS")

	return router
}