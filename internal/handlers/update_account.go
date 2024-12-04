package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := &Query{}
	err := json.NewDecoder(r.Body).Decode(query)
	if err != nil {
		writeErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	if query.Amount <= 0 {
		writeErrorJSON(w, fmt.Errorf("указано некорректное число средств для проведения операции"), http.StatusBadRequest)
		return
	}
	if query.Walletid <= 0 {
		writeErrorJSON(w, fmt.Errorf("некорректный уникальный идентификатор кошелька"), http.StatusBadRequest)
		return
	}
	switch query.OperationType {
	case "DEPOSIT":
		err := h.rep.Deposit(query.Amount, query.Walletid)
		if err != nil {
			writeErrorJSON(w, err, http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode("операция внесения средств прошла успешно")
		return
	case "WITHDRAW":
		err := h.rep.Withdraw(query.Amount, query.Walletid)
		if err != nil {
			writeErrorJSON(w, err, http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode("операция вывода средств прошла успешно")
		return

	default:
		writeErrorJSON(w, fmt.Errorf("несуществующий тип операции"), http.StatusBadRequest)
		return
	}
}

func writeErrorJSON(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Error{Err: err.Error()})
}
