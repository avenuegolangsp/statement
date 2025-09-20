package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"statement-validator/model"
)

func SendTransaction(transaction model.TransactionEvent) error {
	jsonData, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/events", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}

	return nil
}

func TestStatementQuery(userID, accountType, currencyType string) error {
	fmt.Printf("🧪 Testando consulta de extrato para user=%s, account=%s, currency=%s...\n",
		userID, accountType, currencyType)

	startDate := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
	endDate := time.Now().Format("2006-01-02")

	url := fmt.Sprintf("http://localhost:8080/statement/%s/%s/%s/%s/%s", userID, accountType, currencyType, startDate, endDate)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("erro na consulta de extrato: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code inesperado: %d", resp.StatusCode)
	}

	fmt.Printf("✅ Consulta de extrato realizada com sucesso\n")
	return nil
}

// Função para consulta de extrato com datas customizadas
func TestStatementQueryWithDates(userID, accountType, currencyType, startDate, endDate string) error {
	fmt.Printf("🧪 Testando consulta de extrato para user=%s, account=%s, currency=%s, período=%s a %s...\n",
		userID, accountType, currencyType, startDate, endDate)

	url := fmt.Sprintf("http://localhost:8080/statement/%s/%s/%s/%s/%s", userID, accountType, currencyType, startDate, endDate)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("erro na consulta de extrato: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code inesperado: %d", resp.StatusCode)
	}

	fmt.Printf("✅ Consulta de extrato realizada com sucesso\n")
	return nil
}
