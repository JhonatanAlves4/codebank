package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JhonatanAlves4/codebank/domain"
	"github.com/JhonatanAlves4/codebank/infrastructure/repository"
	"github.com/JhonatanAlves4/codebank/usecase"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Server running on Docker")
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Jhonatan"
	cc.ExpirationYear = 2029
	cc.ExpirationMonth = 7
	cc.CVV = 826
	cc.Limit = 750
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil { fmt.Println(err) }
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository) 
	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db", "5432", "postgres", "root", "codebank",	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connection to database")
	}

	return db
}