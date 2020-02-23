package main

type paymentInfoRepository interface {
	findByUUID(string) (*PaymentInfo, error)
	store(*PaymentInfo) (string, error)
	update(*PaymentInfo) error
	deleteByUUID(string) error
}

type paymentInfoRepositoryImpl struct {
}

func (rr *paymentInfoRepositoryImpl) findByUUID(uuid string) (*PaymentInfo, error) {
	return nil, nil
}

func (rr *paymentInfoRepositoryImpl) store(r *PaymentInfo) (string, error) {
	return "", nil
}

func (rr *paymentInfoRepositoryImpl) update(r *PaymentInfo) error {
	return nil
}

func (rr *paymentInfoRepositoryImpl) deleteByUUID(uuid string) error {
	return nil
}
