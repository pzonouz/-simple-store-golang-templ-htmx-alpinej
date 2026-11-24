package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	sqlc "github.com/pzonouz/simple-store-golang-templ-htmx-alpinejs/internal/db"
)

func (s *Service) GetParameters(
	ctx context.Context,
	args sqlc.GetParametersParams,
) ([]sqlc.GetParametersRow, error) {
	Parameters, err := s.queries.GetParameters(ctx, args)
	if err != nil {
		return []sqlc.GetParametersRow{}, err
	}
	if Parameters == nil {
		return []sqlc.GetParametersRow{}, err
	}
	return Parameters, err
}

func (s *Service) GetParametersByParametersGroup(
	ctx context.Context,
	id pgtype.UUID,
) ([]sqlc.GetParametersByParametersGroupRow, error) {
	Parameters, err := s.queries.GetParametersByParametersGroup(ctx, id)
	if err != nil {
		return []sqlc.GetParametersByParametersGroupRow{}, err
	}
	if Parameters == nil {
		return []sqlc.GetParametersByParametersGroupRow{}, err
	}
	return Parameters, err
}

func (s *Service) GetParameter(ctx context.Context, id pgtype.UUID) (sqlc.GetParameterRow, error) {
	Parameter, err := s.queries.GetParameter(ctx, id)
	if err != nil {
		return sqlc.GetParameterRow{}, err
	}
	return Parameter, err
}

func (s *Service) CreateParameter(
	ctx context.Context,
	args sqlc.CreateParameterParams,
) (sqlc.Parameter, error) {
	parameter, err := s.queries.CreateParameter(ctx, args)
	if err != nil {
		return sqlc.Parameter{}, err
	}
	return parameter, nil
}

func (s *Service) UpdateParameter(
	ctx context.Context,
	args sqlc.UpdateParameterParams,
) (sqlc.Parameter, error) {
	parameter, err := s.queries.UpdateParameter(ctx, args)
	if err != nil {
		return sqlc.Parameter{}, err
	}
	return parameter, nil
}

func (s *Service) DeleteParameter(ctx context.Context, id pgtype.UUID) error {
	return s.queries.DeleteParameter(ctx, id)
}
