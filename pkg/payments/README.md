Como usar:

```go
package main

import (
	"fmt"
	"myapp/pkg/payments"
)

func main() {
	gateway := payments.NewMockGateway()

	request := payments.PaymentRequest{
		OrderID:    "12345",
		UserID:     "user-001",
		Amount:     100.50,
		CardNumber: "4111111111111111",
		CardExpiry: "12/25",
		CardCVV:    "123",
	}

	response, err := gateway.ProcessPayment(request)
	if err != nil {
		fmt.Println("Erro ao processar pagamento:", err)
		return
	}

	fmt.Println("Pagamento processado com sucesso!")
	fmt.Printf("Transaction ID: %s\n", response.TransactionID)
	fmt.Printf("Status: %s\n", response.Status)
	fmt.Printf("Timestamp: %s\n", response.Timestamp)
}
```
