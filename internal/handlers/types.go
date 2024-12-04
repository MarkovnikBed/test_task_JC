package handlers

import "java_code/internal/repository"

type Handler struct {
	rep *repository.Repository
}

type Query struct {
	Walletid      int    `json:"walletid"`
	OperationType string `json:"operationType"`
	Amount        int    `json:"amount"`
}

type Error struct {
	Err string `json:"error"`
}

func GetHandler(rep *repository.Repository) *Handler {
	return &Handler{
		rep: rep,
	}
}
