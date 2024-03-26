package repository

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"github.com/whitewolf185/SystemArchitecture/client-service/internal/gen/client_service/public/model"
	"github.com/whitewolf185/SystemArchitecture/client-service/internal/gen/client_service/public/table"
	customerrors "github.com/whitewolf185/SystemArchitecture/client-service/pkg/custom_errors"
)

type personRepo struct {
	db *pgx.Conn
}

func NewPersonController(db *pgx.Conn) personRepo {
	return personRepo{
		db: db,
	}
}

func (r personRepo) GetUserByID(ctx context.Context, clientID uuid.UUID) (*model.Persons, error) {
	sql, args := table.Persons.SELECT(
		table.Persons.ID,
		table.Persons.Username,
		table.Persons.IsDriver,
	).WHERE(
		table.Persons.ID.EQ(postgres.UUID(clientID)),
	).Sql()

	return r.selectPerson(ctx, sql, args)
}

func (r personRepo) GetUserByUserNameIn(ctx context.Context, username string) ([]model.Persons, error) {
	username = "%" + username + "%"
	sql, args := table.Persons.SELECT(
		table.Persons.ID,
		table.Persons.Username,
		table.Persons.IsDriver,
	).WHERE(
		table.Persons.Username.LIKE(postgres.String(username)),
	).Sql()

	var result []model.Persons
	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, customerrors.ErrUserNotFound
		}
		return nil, fmt.Errorf("cannot select user by his username: sql error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var person model.Persons
		err = rows.Scan(
			&person.ID,
			&person.Username,
			&person.IsDriver,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot select user by his username: %w", err)
		}

		result = append(result, person)
	}

	// no persons found
	if len(result) == 0 {
		return nil, customerrors.ErrUserNotFound
	}

	return result, nil
}

func (r personRepo) DeleteUserByID(ctx context.Context, clientID uuid.UUID) (*model.Persons, error) {
	sql, args := table.Persons.DELETE().
		WHERE(table.Persons.ID.EQ(postgres.UUID(clientID))).
		RETURNING(
			table.Persons.ID,
			table.Persons.Username,
			table.Persons.IsDriver,
		).Sql()

	return r.selectPerson(ctx, sql, args)
}

func (r personRepo) CreateUser(ctx context.Context, req model.Persons) (*model.Persons, error) {
	sql, args := table.Persons.SELECT(
		table.Persons.ID,
		table.Persons.Username,
		table.Persons.IsDriver,
	).WHERE(
		table.Persons.Username.EQ(postgres.String(req.Username)),
	).Sql()

	person, err := r.selectPerson(ctx, sql, args)
	if err != nil {
		logrus.Errorf("cannot validate user before insert: %v", err)
	}
	if person != nil {
		return nil, customerrors.ErrUserAlreadyExists
	}

	req.UserPassword = base64.StdEncoding.EncodeToString([]byte(req.UserPassword))
	sql, args = table.Persons.INSERT().MODEL(req).
		RETURNING(
			table.Persons.ID,
			table.Persons.Username,
			table.Persons.IsDriver,
		).Sql()

	return r.selectPerson(ctx, sql, args)
}
