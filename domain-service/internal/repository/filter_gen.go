package repository

import (
	"reflect"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (dr domainRepo) filterConstructorRoute(filter interface{}) bson.D {
	r := reflect.ValueOf(filter)
	newFilter := make(bson.D, 0, r.NumField())

	for i := 0; i < r.NumField(); i++ {
		tagName := r.Type().Field(i).Tag.Get("json")
		if tagName == "client_id" {
			tagName = "_id"
		}

		switch value := r.Field(i).Interface().(type) {
		case string:
			if value != "" {
				newFilter = append(newFilter, bson.E{Key: tagName, Value: r.Field(i).Interface()})
			}

		case time.Time:
		case []string:
			if len(value) != 0 {
				newFilter = append(newFilter, bson.E{Key: tagName, Value: bson.M{"$in": r.Field(i).Interface()}})
			}
		}
	}

	logrus.Infoln("new filter:\n", newFilter)
	return newFilter
}
