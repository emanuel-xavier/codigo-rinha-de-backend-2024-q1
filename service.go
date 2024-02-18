package main

import "context"

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

func (cs *ClientService) clientExists(c context.Context, id int) (bool, error) {
	return cs.clientRepo.checkIfClientExists(c, id)
}

func (cs *ClientService) getClientBalance(c context.Context, id int) (*BalanceDto, error) {
	return cs.clientRepo.getClientBalance(c, id)
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
