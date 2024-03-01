package service

import (
	"context"
	"time"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/entity"
	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/repository"
	"github.com/jackc/pgx/v5"
)

type ClientService struct {
	cRepo repository.ClientRepository
	tRepo repository.TransactionRepository
}

func NewClientService(cRepo repository.ClientRepository, tRepo repository.TransactionRepository) *ClientService {
	return &ClientService{
		cRepo: cRepo,
		tRepo: tRepo,
	}
}

func (serv *ClientService) GetClientById(ctx context.Context, id int) (*entity.Client, error) {
	client, err := serv.cRepo.GetClientById(ctx, id)
	return &client, err
}

func (serv *ClientService) GetClientByIdAndLock(ctx context.Context, id int) (*entity.Client, pgx.Tx, error) {
	client, tx, err := serv.cRepo.GetClientByIdAndLock(ctx, id)
	return &client, tx, err
}

func (serv *ClientService) GetClientStatemant(ctx context.Context, id int) (*entity.Statement, error) {
	client, err := serv.cRepo.GetClientById(ctx, id)
	if err != nil {
		return nil, err
	}

	transactions, err := serv.tRepo.GetTenLastTransactionsByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	balance := entity.StatementBalance{
		Amount: client.Balance,
		Date:   time.Now(),
		Limit:  client.Limit,
	}

	return &entity.Statement{
		Balance:             balance,
		LastTenTransactions: transactions,
	}, nil
}
