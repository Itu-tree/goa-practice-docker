package calcapi

import (
	calc "calc/gen/calc"
	"context"
	"log"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	return p.A + p.B, nil
}

// Add implements sub.
func (s *calcsrvc) Sub(ctx context.Context, p *calc.SubPayload) (res int, err error) {
	return p.A - p.B, nil
}

// Add implements mul.
func (s *calcsrvc) Mul(ctx context.Context, p *calc.MulPayload) (res int, err error) {
	return p.A * p.B, nil
}

// Divide returns the integral division of two integers.
func (s *calcsrvc) Divide(ctx context.Context, p *calc.DividePayload) (res int, err error) {
	return p.A / p.B, nil
}
