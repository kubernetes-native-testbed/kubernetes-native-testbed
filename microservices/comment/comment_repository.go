package main

type commentRepository interface {
	findByUUID(string) (*Comment, error)
	store(*Comment) (string, error)
	update(*Comment) error
	deleteByUUID(string) error
}

type commentRepositoryImpl struct {
}

func (rr *commentRepositoryImpl) findByUUID(uuid string) (*Comment, error) {
	return nil, nil
}

func (rr *commentRepositoryImpl) store(r *Comment) (string, error) {
	return "", nil
}

func (rr *commentRepositoryImpl) update(r *Comment) error {
	return nil
}

func (rr *commentRepositoryImpl) deleteByUUID(uuid string) error {
	return nil
}
