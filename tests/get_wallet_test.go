package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"java_code/internal/repository"
)

func TestGetWallet(t *testing.T) {
	rep := repository.CreateRepository()
	defer rep.DB.Close()
	res, err := rep.DB.Exec("INSERT INTO wallets (uuid, balance) VALUES(12345,1234)")
	require.NoError(t, err)
	num, err := res.RowsAffected()
	require.NoError(t, err)
	assert.Equal(t, int(num), 1)
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:8080/api/v1/wallets/12345", nil)
	require.NoError(t, err)
	resp, err := client.Do(request)
	require.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	_, err = rep.DB.Exec("DELETE FROM wallets WHERE uuid=12345")
	require.NoError(t, err)
	request, err = http.NewRequest("GET", "http://localhost:8080/api/v1/wallets/12345", nil)
	require.NoError(t, err)
	resp, err = client.Do(request)
	require.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
}
