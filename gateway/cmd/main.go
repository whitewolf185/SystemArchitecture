package main

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/whitewolf185/SystemArchitecture/gateway/api/http_balancer/circuit_breaker"
	"github.com/whitewolf185/SystemArchitecture/gateway/api/http_balancer/domain_balancer"
	"github.com/whitewolf185/SystemArchitecture/gateway/api/http_balancer/people_balancer"
	"github.com/whitewolf185/SystemArchitecture/gateway/api/middleware"
	"github.com/whitewolf185/SystemArchitecture/gateway/api/router"
	"github.com/whitewolf185/SystemArchitecture/gateway/internal/config"
	"github.com/whitewolf185/SystemArchitecture/gateway/internal/config/flags"
)

func main() {
	flags.InitServiceFlags()
	ctx := context.Background()

	// -- connections --
	keydbClient, err := config.NewKeyDBConn(ctx)
	if err != nil {
		logrus.Fatalf("cannot connect to keydb: %s", err.Error())
	}

	// -- depencies --
	circuidBreaker := circuit_breaker.NewCircuiBreaker(3, time.Minute)
	peopleBalancer := people_balancer.New("person")
	domainBalancer := domain_balancer.New("domain", keydbClient, circuidBreaker)

	// Имплементация API

	root := router.NewRouter(middleware.NewErrorHandler(peopleBalancer, domainBalancer))

	logrus.Info("preparation to start successful")
	err = http.ListenAndServe(":"+config.GetValue(config.ListenPort), root)
	if err != nil {
		logrus.Fatalln("cannot start app ", err)
	}
}
