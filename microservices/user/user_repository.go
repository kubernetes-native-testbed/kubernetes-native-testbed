package main

type userRepository interface {
	findByUUID(string) (*User, error)
	store(*User) (string, error)
	update(*User) error
	deleteByUUID(string) error
}

type userRepositoryImpl struct {
}

func (rr *userRepositoryImpl) findByUUID(uuid string) (*User, error) {
	return nil, nil
}

func (rr *userRepositoryImpl) store(r *User) (string, error) {
	return "", nil
}

func (rr *userRepositoryImpl) update(r *User) error {
	return nil
}

func (rr *userRepositoryImpl) deleteByUUID(uuid string) error {
	return nil
}
