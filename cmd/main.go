package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	"java_code/internal/handlers"
	"java_code/internal/repository"
)

func init() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal(err)
	}

}

var ch = make(chan struct{}, 4)

func MiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ch <- struct{}{}
		next(w, r)
		<-ch
	}
}

func main() {
	rep := repository.CreateRepository()
	rep.PrepareTable()
	router := chi.NewRouter()
	handler := handlers.GetHandler(rep)
	router.Post("/api/v1/wallet", MiddleWare(handler.UpdateAccount))
	router.Get("/api/v1/wallets/{WALLET_UUID}", MiddleWare(handler.GetWallet))
	log.Fatal(http.ListenAndServe(":8080", router))
}
