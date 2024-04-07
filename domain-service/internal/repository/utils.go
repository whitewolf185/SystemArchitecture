package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"github.com/whitewolf185/SystemArchitecture/domain-service/internal/gen/client_service/public/model"
)

// возвращается pgx.ErrNoRows, если пользователя не получилось получить
func (r personRepo) selectPerson(ctx context.Context, sql string, args []interface{}) (*model.Persons, error) {
	var result model.Persons
	row := r.db.QueryRow(ctx, sql, args...)
	err := row.Scan(
		&result.ID,
		&result.Username,
		&result.IsDriver,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		logrus.Info("selectPerson: person not found")
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("cannot select person: %w", err)
	}
	return &result, nil
}
