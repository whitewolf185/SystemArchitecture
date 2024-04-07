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
	db, err := config.ConnectPostgres(ctx)
	if err != nil {
		logrus.Fatalln("cannot connect to postgresql ", err)
	}

	// -- depencies --
	personRepo := repository.NewPersonController(db)

	// Имплементация API
	application, err := app.NewImplementation(personRepo)
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
