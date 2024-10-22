package services

type ItemServiceInterface interface {
}

type ItemServiceImplementation struct {
}

func NewItemServiceImplementation() *ItemServiceImplementation {
	return &ItemServiceImplementation{}
}
