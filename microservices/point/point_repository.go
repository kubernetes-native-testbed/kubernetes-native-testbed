package main

type pointRepository interface {
	findByUUID(string) (*Point, error)
	store(*Point) (string, error)
	update(*Point) error
	deleteByUUID(string) error
}

type pointRepositoryImpl struct {
}

func (pr *pointRepositoryImpl) findByUUID(uuid string) (*Point, error) {
	return nil, nil
}

func (pr *pointRepositoryImpl) store(r *Point) (string, error) {
	return "", nil
}

func (pr *pointRepositoryImpl) update(r *Point) error {
	return nil
}

func (pr *pointRepositoryImpl) deleteByUUID(uuid string) error {
	return nil
}
