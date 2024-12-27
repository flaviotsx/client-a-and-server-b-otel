package web

import (
	"context"

	"github.com/flaviotsx/temperature-zipcode-srv/internal/entity"
	"go.opentelemetry.io/otel/trace"
)

type GetTemperatureByZipCodeUseCaseInterface interface {
	Execute(ctx context.Context, zipcode string, tracer trace.Tracer) (*entity.GetTemperatureByZipcodeResponse, *entity.ErrorResponse)
}
