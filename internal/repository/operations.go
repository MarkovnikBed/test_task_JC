package repository

import "fmt"

func (rep *Repository) Deposit(amount int, walletid int) error {
	query := fmt.Sprintf(`
	INSERT INTO wallets (uuid, balance) VALUES(%d, %d) 
	ON CONFLICT (uuid) 
	DO UPDATE SET balance=wallets.balance+%d`,
		walletid, amount, amount)
	_, err := rep.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (rep *Repository) Withdraw(amount int, walletid int) error {
	query := fmt.Sprintf(`
	UPDATE wallets 
	SET balance=balance-%d 
	WHERE uuid=%d`,
		amount, walletid)
	res, err := rep.DB.Exec(query)
	if err != nil {
		return err
	}
	numrows, _ := res.RowsAffected()
	if numrows == 0 {
		return fmt.Errorf("операция вывода средств не может осуществлять с несуществующего аккаунта")
	}
	return nil
}
