package repository

import (
	"github.com/flaviotsx/temperature-zipcode-srv/internal/entity"
	"github.com/stretchr/testify/mock"
)

type WeatherApiRepositoryMock struct {
	mock.Mock
}

func (m *WeatherApiRepositoryMock) GetWeatherData(city string) (*entity.WeatherApiResponse, *entity.ErrorResponse) {
	args := m.Called(city)

	if args.Get(1) != nil {
		return nil, args.Get(1).(*entity.ErrorResponse)
	}

	return args.Get(0).(*entity.WeatherApiResponse), nil
}
