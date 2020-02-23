package main

type deliveryStatusRepository interface {
	findByUUID(string) (*DeliveryStatus, error)
	store(*DeliveryStatus) (string, error)
	update(*DeliveryStatus) error
	deleteByUUID(string) error
}

type deliveryStatusRepositoryImpl struct {
}

func (rr *deliveryStatusRepositoryImpl) findByUUID(uuid string) (*DeliveryStatus, error) {
	return nil, nil
}

func (rr *deliveryStatusRepositoryImpl) store(r *DeliveryStatus) (string, error) {
	return "", nil
}

func (rr *deliveryStatusRepositoryImpl) update(r *DeliveryStatus) error {
	return nil
}

func (rr *deliveryStatusRepositoryImpl) deleteByUUID(uuid string) error {
	return nil
}
