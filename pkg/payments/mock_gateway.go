package payments

import (
	"errors"
	"math/rand"
	"time"
)

// PaymentStatus representa o status de uma transação
type PaymentStatus string

const (
	StatusApproved PaymentStatus = "approved"
	StatusDeclined PaymentStatus = "declined"
	StatusPending  PaymentStatus = "pending"
)

// PaymentRequest representa uma solicitação de pagamento
type PaymentRequest struct {
	OrderID    string  `json:"order_id"`
	UserID     string  `json:"user_id"`
	Amount     float64 `json:"amount"`
	CardNumber string  `json:"card_number"`
	CardExpiry string  `json:"card_expiry"`
	CardCVV    string  `json:"card_cvv"`
}

// PaymentResponse representa a resposta do mock do gateway de pagamento
type PaymentResponse struct {
	TransactionID string        `json:"transaction_id"`
	Status        PaymentStatus `json:"status"`
	Timestamp     time.Time     `json:"timestamp"`
}

// MockGateway simula um gateway de pagamento
type MockGateway struct{}

// NewMockGateway cria uma nova instância do mock
func NewMockGateway() *MockGateway {
	return &MockGateway{}
}

// ProcessPayment simula o processamento de um pagamento
func (mg *MockGateway) ProcessPayment(req PaymentRequest) (*PaymentResponse, error) {
	if req.Amount <= 0 {
		return nil, errors.New("o valor do pagamento deve ser maior que zero")
	}

	// Simulando um tempo de processamento
	time.Sleep(time.Millisecond * 500)

	// Gerando um ID de transação falso
	transactionID := generateTransactionID()

	// Simulando a resposta com base em lógica aleatória
	status := getRandomStatus()

	return &PaymentResponse{
		TransactionID: transactionID,
		Status:        status,
		Timestamp:     time.Now(),
	}, nil
}

// generateTransactionID gera um identificador aleatório para a transação
func generateTransactionID() string {
	rand.Seed(time.Now().UnixNano())
	return "TX-" + randomString(12)
}

// getRandomStatus retorna um status aleatório para simular a resposta do gateway
func getRandomStatus() PaymentStatus {
	rand.Seed(time.Now().UnixNano())
	statuses := []PaymentStatus{StatusApproved, StatusDeclined, StatusPending}
	return statuses[rand.Intn(len(statuses))]
}

// randomString gera uma string aleatória para simular IDs de transação
func randomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
