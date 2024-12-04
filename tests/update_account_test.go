package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"java_code/internal/handlers"
	"java_code/internal/repository"
)

func sendPostRequest(t *testing.T, URL string, body handlers.Query) *http.Response {
	client := &http.Client{}
	requestBody, err := json.Marshal(body)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", URL, bytes.NewReader(requestBody))
	require.NoError(t, err)
	response, err := client.Do(req)
	require.NoError(t, err)
	return response
}

func TestUpdateAccount(t *testing.T) {
	rep := repository.CreateRepository()
	defer rep.DB.Close()
	handler := handlers.GetHandler(rep)
	server := httptest.NewServer(http.HandlerFunc(handler.UpdateAccount))
	defer server.Close()
	body := handlers.Query{
		OperationType: "DEPOSIT",
		Walletid:      123456789,
		Amount:        100,
	}
	resp := sendPostRequest(t, server.URL+"/api/v1/wallet", body)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	body = handlers.Query{
		OperationType: "DEPOerSIT",
		Walletid:      123456789,
		Amount:        100,
	}
	resp = sendPostRequest(t, server.URL, body)
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
	body = handlers.Query{
		OperationType: "DEPOSIT",
		Walletid:      -123456789,
		Amount:        100,
	}
	resp = sendPostRequest(t, server.URL, body)
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
	body = handlers.Query{
		OperationType: "DEPOSIT",
		Walletid:      123456789,
		Amount:        -100,
	}
	resp = sendPostRequest(t, server.URL, body)
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
	body = handlers.Query{
		OperationType: "WITHDRAW",
		Walletid:      123456789,
		Amount:        100,
	}
	resp = sendPostRequest(t, server.URL, body)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	body = handlers.Query{
		OperationType: "WITHDRAW",
		Walletid:      123456789,
		Amount:        -100,
	}
	resp = sendPostRequest(t, server.URL, body)

	require.Equal(t, resp.StatusCode, http.StatusBadRequest)
	body = handlers.Query{
		OperationType: "WITHDRAW",
		Walletid:      -123456789,
		Amount:        100,
	}
	resp = sendPostRequest(t, server.URL, body)
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
	body = handlers.Query{
		OperationType: "WITHDRAW",
		Walletid:      134256789,
		Amount:        100,
	}
	resp = sendPostRequest(t, server.URL, body)
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
	rep.DB.Exec("DELETE FROM wallets WHERE uuid=123456789")
}
