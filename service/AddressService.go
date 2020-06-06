package service

import "myapp/graph/model"

func GetAddress() ([]*model.Address, error) {
	var address []*model.Address
	address = append(address, &model.Address{
		ID:    "1",
		Jalan: "asd",
	})

	return address, nil
}
