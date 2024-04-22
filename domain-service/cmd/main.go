package main

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/middleware"
	"github.com/whitewolf185/SystemArchitecture/domain-service/api/router"
	"github.com/whitewolf185/SystemArchitecture/domain-service/internal/app"
	"github.com/whitewolf185/SystemArchitecture/domain-service/internal/config"
	"github.com/whitewolf185/SystemArchitecture/domain-service/internal/config/flags"
	"github.com/whitewolf185/SystemArchitecture/domain-service/internal/repository"
)

func main() {
	ctx := context.Background()
	flags.InitServiceFlags()

	// -- depencies --
	mongoClient, err := config.NewMongoDbConnection(ctx)
	if err != nil {
		logrus.Fatalln("cannot configure DB ", err)
	}
	defer mongoClient.Disconnect(ctx)

	domainRepo := repository.NewDomainRepo(mongoClient)

	// Имплементация API
	application, err := app.NewImplementation(domainRepo)
	if err != nil {
		logrus.Fatalln("cannot configure implementation")
	}

	root := router.NewRouter(middleware.NewErrorHandler(application))

	logrus.Info("preparation to start successful")
	err = http.ListenAndServe(":"+config.GetValue(config.ListenPort), root)
	if err != nil {
		logrus.Fatalln("cannot start app ", err)
	}
}
