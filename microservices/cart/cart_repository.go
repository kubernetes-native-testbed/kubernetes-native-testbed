package main

type cartRepository interface {
	findByUUID(string) (*Cart, error)
	store(*Cart) (string, error)
	update(*Cart) error
	deleteByUUID(string) error
}

type cartRepositoryImpl struct {
}

func (rr *cartRepositoryImpl) findByUUID(uuid string) (*Cart, error) {
	return nil, nil
}

func (rr *cartRepositoryImpl) store(r *Cart) (string, error) {
	return "", nil
}

func (rr *cartRepositoryImpl) update(r *Cart) error {
	return nil
}

func (rr *cartRepositoryImpl) deleteByUUID(uuid string) error {
	return nil
}
