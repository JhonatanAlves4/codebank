package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JhonatanAlves4/codebank/infrastructure/grpc/server"
	"github.com/JhonatanAlves4/codebank/infrastructure/kafka"
	"github.com/JhonatanAlves4/codebank/infrastructure/repository"
	"github.com/JhonatanAlves4/codebank/usecase"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Server running on Docker")
	db := setupDb()
	defer db.Close()
	producer := setupKafkaProducer()
	processTransactionUseCase := setupTransactionUseCase(db, producer)
	serveGrpc(processTransactionUseCase)
}

func setupTransactionUseCase(db *sql.DB, producer kafka.KafkaProducer) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository) 
	useCase.KafkaProducer = producer
	return useCase
}

func setupKafkaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer("host.docker.internal:9094")
	return producer
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

func serveGrpc(processTransactionUseCase usecase.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessTransactionUseCase = processTransactionUseCase
	fmt.Println("Running gRPC server")
	grpcServer.Serve()
}

// cc := domain.NewCreditCard()
// 	cc.Number = "1234"
// 	cc.Name = "Jhonatan"
// 	cc.ExpirationYear = 2029
// 	cc.ExpirationMonth = 7
// 	cc.CVV = 826
// 	cc.Limit = 750
// 	cc.Balance = 0

// 	repo := repository.NewTransactionRepositoryDb(db)
// 	err := repo.CreateCreditCard(*cc)
// 	if err != nil { fmt.Println(err) }