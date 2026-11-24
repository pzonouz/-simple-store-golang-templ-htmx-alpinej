package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	sqlc "github.com/pzonouz/simple-store-golang-templ-htmx-alpinejs/internal/db"
)

func (s *Service) GetParametersGroups(
	ctx context.Context,
	args sqlc.GetParametersGroupsParams,
) ([]sqlc.GetParametersGroupsRow, error) {
	ParametersGroups, err := s.queries.GetParametersGroups(ctx, args)
	if err != nil {
		return []sqlc.GetParametersGroupsRow{}, err
	}
	if ParametersGroups == nil {
		return []sqlc.GetParametersGroupsRow{}, err
	}
	return ParametersGroups, err
}

func (s *Service) GetParametersGroup(
	ctx context.Context,
	id pgtype.UUID,
) (sqlc.GetParametersGroupRow, error) {
	ParameterGroup, err := s.queries.GetParametersGroup(ctx, id)
	if err != nil {
		return sqlc.GetParametersGroupRow{}, err
	}
	return ParameterGroup, err
}

func (s *Service) CreateParametersGroup(
	ctx context.Context,
	args sqlc.CreateParametersGroupParams,
) (sqlc.ParametersGroup, error) {
	parameterGroup, err := s.queries.CreateParametersGroup(ctx, args)
	if err != nil {
		return sqlc.ParametersGroup{}, err
	}
	return parameterGroup, nil
}

func (s *Service) UpdateParametersGroup(
	ctx context.Context,
	args sqlc.UpdateParametersGroupParams,
) (sqlc.ParametersGroup, error) {
	parametersGroup, err := s.queries.UpdateParametersGroup(ctx, args)
	if err != nil {
		return sqlc.ParametersGroup{}, err
	}
	return parametersGroup, nil
}

func (s *Service) DeleteParametersGroup(ctx context.Context, id pgtype.UUID) error {
	return s.queries.DeleteParametersGroup(ctx, id)
}

func (s *Service) GetParameterGroupsByCategoryAndAncestor(
	ctx context.Context,
	id pgtype.UUID,
) ([]sqlc.ParametersGroup, error) {
	return s.queries.GetParametersGroupsByCategoryAndAncestors(ctx, id)
}
