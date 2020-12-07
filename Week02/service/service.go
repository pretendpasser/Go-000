package service

import (
	"week02/dao"
	"github.com/pkg/errors"
)

type Service struct {
	dao *dao.Dao
}

func NewService() *Service {
	return &Service{dao.NewDao()}
}

func (s *Service) GetUsernameByUserById(id int) (u dao.User, err error) {
	s = NewService()
	u, err = s.dao.FindUserById(id);
	return u, errors.Wrapf(err, "service -%d- miss match!", id)
}
