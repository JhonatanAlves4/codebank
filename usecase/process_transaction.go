package usecase

import (
	"encoding/json"
	"time"

	"github.com/JhonatanAlves4/codebank/domain"
	"github.com/JhonatanAlves4/codebank/dto"
	"github.com/JhonatanAlves4/codebank/infrastructure/kafka"
)

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
	KafkaProducer kafka.KafkaProducer
}

func NewUseCaseTransaction(transactionRepository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository}
}

func (u UseCaseTransaction) ProcessTransaction(transactionDTO dto.TransactionDTO) (domain.Transaction, error) {
	creditCard := u.hydrateCreditCard(transactionDTO)
	ccBalanceAndLimit, err := u.TransactionRepository.GetCreditCard(*creditCard)
	if err != nil { return domain.Transaction{}, err }
	
	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Limit = ccBalanceAndLimit.Limit
	creditCard.Balance = ccBalanceAndLimit.Balance

	t := u.NewTransaction(transactionDTO, ccBalanceAndLimit)
	t.ProcessAndValidate(creditCard)

	err = u.TransactionRepository.SaveTransaction(*t, *creditCard)
	if err != nil { return domain.Transaction{}, err }

	transactionDTO.ID = t.ID
	transactionDTO.CreatedAt = t.CreatedAt
	transactionJson, err := json.Marshal(transactionDTO)
	if err != nil {
		return domain.Transaction{}, err
	}
	err = u.KafkaProducer.Publish(string(transactionJson), "payments")
	if err != nil {
		return domain.Transaction{}, err
	}

	return *t, nil
}

func (u UseCaseTransaction) hydrateCreditCard(transactionDTO dto.TransactionDTO) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDTO.Name
	creditCard.Number = transactionDTO.Number
	creditCard.ExpirationMonth = transactionDTO.ExpirationMonth
	creditCard.ExpirationYear = transactionDTO.ExpirationYear
	creditCard.CVV = transactionDTO.CVV

	return creditCard
}

func (u UseCaseTransaction) NewTransaction(transactionDTO dto.TransactionDTO, cc domain.CreditCard) *domain.Transaction {
	t := domain.NewTransaction()
	t.CreditCardId = cc.ID
	t.Amount = transactionDTO.Amount
	t.Store = transactionDTO.Store
	t.Description = transactionDTO.Description
	t.CreatedAt = time.Now()

	return t
}