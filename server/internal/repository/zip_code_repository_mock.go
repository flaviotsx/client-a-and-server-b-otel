package repository

import (
	"github.com/flaviotsx/temperature-zipcode-srv/internal/entity"
	"github.com/stretchr/testify/mock"
)

type ZipCodeRepositoryMock struct {
	mock.Mock
}

func (m *ZipCodeRepositoryMock) GetZipcodeData(zipcode string) (*entity.ZipcodeDataResponse, *entity.ErrorResponse) {
	args := m.Called(zipcode)

	if args.Get(1) != nil {
		return nil, args.Get(1).(*entity.ErrorResponse)
	}

	return args.Get(0).(*entity.ZipcodeDataResponse), nil
}
