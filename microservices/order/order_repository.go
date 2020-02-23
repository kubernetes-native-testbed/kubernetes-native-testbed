package main

type orderRepository interface {
	findByUUID(string) (*Order, error)
	store(*Order) (string, error)
	update(*Order) error
	deleteByUUID(string) error
}

type orderRepositoryImpl struct {
}

func (rr *orderRepositoryImpl) findByUUID(uuid string) (*Order, error) {
	return nil, nil
}

func (rr *orderRepositoryImpl) store(r *Order) (string, error) {
	return "", nil
}

func (rr *orderRepositoryImpl) update(r *Order) error {
	return nil
}

func (rr *orderRepositoryImpl) deleteByUUID(uuid string) error {
	return nil
}
