package services

import "server/repositories"

type ItemServiceInterface interface {
}

type ItemServiceImplementation struct {
	ItemRepository       repositories.ItemRepositoryInterface
	TransactorRepository repositories.TransactorRepositoryInterface
}

func NewItemServiceImplementation(
	ItemRepository repositories.ItemRepositoryInterface,
	TransactorRepository repositories.TransactorRepositoryInterface,
) *ItemServiceImplementation {
	return &ItemServiceImplementation{
		ItemRepository:       ItemRepository,
		TransactorRepository: TransactorRepository,
	}
}
