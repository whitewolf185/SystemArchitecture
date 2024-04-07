package repository

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/domain-service/internal/config"
	customerrors "github.com/whitewolf185/SystemArchitecture/domain-service/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type domainRepo struct {
	routeCollection     *mongo.Collection
	companionCollection *mongo.Collection
}

func NewDomainRepo(client *mongo.Client) domainRepo {
	return domainRepo{
		routeCollection:     client.Database(config.MongoDbName).Collection(config.RouteCollectionName),
		companionCollection: client.Database(config.MongoDbName).Collection(config.CompanionCollectionName),
	}
}

// InsertRoute - вставка значения в
func (dr domainRepo) InsertRoute(ctx context.Context, dataToInsert interface{}) error {
	insertedData, err := dr.routeCollection.InsertOne(ctx, dataToInsert)
	if err != nil {
		return fmt.Errorf("cannot insert route data %v\nErr: %w", dataToInsert, err)
	}

	logrus.Infof("the data was inserted into route. ID: %v", insertedData)
	return nil
}
func (dr domainRepo) InsertCompanion(ctx context.Context, dataToInsert interface{}) error {
	insertedData, err := dr.companionCollection.InsertOne(ctx, dataToInsert)
	if err != nil {
		return fmt.Errorf("cannot insert companion data %v\nErr: %w", dataToInsert, err)
	}

	logrus.Infof("the data was inserted into companion. ID: %v", insertedData)
	return nil
}

func (dr domainRepo) DeleteRoute(ctx context.Context, idToDelete string) error {
	filter := bson.D{{Key: "_id", Value: idToDelete}}
	result, err := dr.routeCollection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("cannot delete route[%s]: %w", idToDelete, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("cannot delete route[%s]: %w", idToDelete, customerrors.ErrNotFound)
	}

	return nil
}

func (dr domainRepo) DeleteCompanion(ctx context.Context, idToDelete string) error {
	filter := bson.D{{Key: "_id", Value: idToDelete}}
	result, err := dr.companionCollection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("cannot delete companion[%s]: %w", idToDelete, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("cannot delete companion[%s]: %w", idToDelete, customerrors.ErrNotFound)
	}

	return nil
}

func (dr domainRepo) GetCompanions(ctx context.Context, filter domain.Companion) ([]domain.Companion, error) {
	cursor, err := dr.companionCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("cannot get companions: query truble: %w", err)
	}
	defer cursor.Close(ctx)

	var result []domain.Companion
	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("cannot get companions: write result: %w", err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("cannot get companions: %w", customerrors.ErrNotFound)
	}

	return result, nil
}

func (dr domainRepo) GetRoutes(ctx context.Context, filter domain.Route) ([]domain.Route, error) {
	cursor, err := dr.routeCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("cannot get routes: query truble: %w", err)
	}
	defer cursor.Close(ctx)

	var result []domain.Route
	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("cannot get routes: write result: %w", err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("cannot get routes: %w", customerrors.ErrNotFound)
	}

	return result, nil
}
