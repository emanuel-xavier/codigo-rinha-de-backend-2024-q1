package main

type ClientService struct {
	transactionRepo TransactionRepository
	clientRepo      ClientRepository
}

func NewClientService(transactionRepo TransactionRepository, clientRepo ClientRepository) *ClientService {
	return &ClientService{
		transactionRepo: transactionRepo,
		clientRepo:      clientRepo,
	}
}

type TransactionService struct {
	transactionRepo TransactionRepository
	clientRepo      ClientRepository
}

func NewTransactionService(transactionRepo TransactionRepository, clientRepo ClientRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		clientRepo:      clientRepo,
	}
}
