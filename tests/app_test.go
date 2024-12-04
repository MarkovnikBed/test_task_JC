package tests

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"

	"java_code/internal/repository"
)

func init() {
	err := godotenv.Load("../config.env")
	if err != nil {
		log.Fatal(err)
	}
}
func TestDB(t *testing.T) {
	rep := repository.CreateRepository()
	defer rep.DB.Close()
	err := rep.DB.Ping()
	require.NoError(t, err)
	row := rep.DB.QueryRow(`SELECT EXISTS (
		SELECT 1 
		FROM information_schema.tables 
		WHERE table_schema = 'public' 
		AND table_name = 'wallets'
	);`)
	require.NoError(t, err)
	var exist string
	err = row.Scan(&exist)
	require.NoError(t, err)
	require.Equal(t, exist, "true")

}
