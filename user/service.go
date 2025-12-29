package user

import "main/domain"

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{usrRepo: usrRepo}
}

func (svc *service) CreateUser(u domain.User) (*domain.User, error) {
	usr, err := svc.usrRepo.CreateUser(u)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}

func (svc *service) GetUserByEmail(email string, password string) (*domain.User, error) {
	usr, err := svc.usrRepo.GetUserByEmail(email, password)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}
