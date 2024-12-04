package repository

import "fmt"

func (h *Repository) GetBalance(uuid int) (string, error) {
	query := fmt.Sprintf("SELECT balance FROM wallets WHERE uuid=%d", uuid)
	row := h.DB.QueryRow(query)
	var balance string
	row.Scan(&balance)
	if balance == "" {
		return "", fmt.Errorf("не удалось получить данные о балансе кошелька")
	}
	return balance, nil
}
