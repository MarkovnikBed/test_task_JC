package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *Handler) GetWallet(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	walletId := chi.URLParam(r, "WALLET_UUID")
	uuid, err := strconv.Atoi(walletId)
	log.Printf("Cервер получил запрос баланса по walletid : %d", uuid)
	if err != nil {
		log.Printf("Cервер не смог обработать запрос баланса по walletid : %d", uuid)
		writeErrorJSON(w, fmt.Errorf("некорректный уникальный идентификатор кошелька"), http.StatusBadRequest)
		return
	}
	if uuid <= 0 {
		log.Printf("Cервер не смог обработать запрос баланса по walletid : %d", uuid)
		writeErrorJSON(w, fmt.Errorf("некорректный уникальный идентификатор кошелька"), http.StatusBadRequest)
		return
	}
	balance, err := h.rep.GetBalance(uuid)
	if err != nil {
		log.Printf("Cервер не смог обработать запрос баланса по walletid : %d", uuid)
		writeErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	log.Printf("Сервер отправил значение баланса по walletid : %d", uuid)
	json.NewEncoder(w).Encode("баланс - " + balance)
}
