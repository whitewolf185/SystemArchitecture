package http_balancer

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"

	customerrors "github.com/whitewolf185/SystemArchitecture/gateway/pkg/custom_errors"
)

func PrepareQuery(request any) url.Values {
	result := make(url.Values)
	if request == nil {
		return result
	}
	fields := reflect.TypeOf(request).Elem()
	// fmt.Println(fields.NumField())
	for i := 0; i < fields.NumField(); i++ {
		currField := fields.Field(i)
		// fmt.Println(currField)
		// fmt.Println(currField.Tag.Get("in"))
		if strings.HasPrefix(currField.Tag.Get("in"), "query=") {
			queryKey := strings.Split(currField.Tag.Get("in"), "=")[1]
			queryValue := reflect.ValueOf(request).Elem().Field(i).String()
			// fmt.Printf("queryKey: %s, queryValue: %s\n", queryKey, queryValue)
			result.Set(queryKey, queryValue)
		}
	}
	return result
}

func RequestSender(request *http.Request) ([]byte, error) {
	client := &http.Client{}

	request.Header.Add("Content-Type", "application/json")
	logrus.Info(request)

	res, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("gateway: unkwodn error: %w", err)
	}
	if res == nil {
		return nil, fmt.Errorf("empty response")
	}
	logrus.Infof("status code: %d, type: %T", res.StatusCode, res.StatusCode)
	switch res.StatusCode {
	case http.StatusBadRequest:
		return nil, customerrors.CodesBadRequest(err)
	case http.StatusInternalServerError:
		return nil, err
	case http.StatusUnauthorized:
		return nil, customerrors.CodesUnauthorized(err)
	case http.StatusNotFound:
		return nil, customerrors.CodesNotFound(err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("gateway: cannot ready body: %w", err)
	}
	if resBody == nil {
		return nil, nil
	}
	logrus.Info(string(resBody))

	return resBody, nil
}
