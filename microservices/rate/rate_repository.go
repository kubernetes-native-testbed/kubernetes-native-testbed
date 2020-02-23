package main

type rateRepository interface {
	findByUUID(string) (*Rate, error)
	store(*Rate) (string, error)
	update(*Rate) error
	deleteByUUID(string) error
}

type rateRepositoryImpl struct {
}

func (rr *rateRepositoryImpl) findByUUID(uuid string) (*Rate, error) {
	return nil, nil
}

func (rr *rateRepositoryImpl) store(r *Rate) (string, error) {
	return "", nil
}

func (rr *rateRepositoryImpl) update(r *Rate) error {
	return nil
}

func (rr *rateRepositoryImpl) deleteByUUID(uuid string) error {
	return nil
}
