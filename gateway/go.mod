module github.com/whitewolf185/SystemArchitecture/gateway

go 1.22.1

require (
	github.com/ggicci/httpin v0.18.1
	github.com/go-chi/chi v1.5.5
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/pkg/errors v0.9.1
	github.com/redis/go-redis/v9 v9.5.1
	github.com/sirupsen/logrus v1.9.3
	github.com/swaggo/http-swagger/v2 v2.0.2
	github.com/swaggo/swag v1.16.3
	github.com/whitewolf185/SystemArchitecture/client-service v1.0.1
	github.com/whitewolf185/SystemArchitecture/domain-service v1.0.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/ggicci/owl v0.8.2 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/spec v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/swaggo/files/v2 v2.0.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/tools v0.19.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/whitewolf185/SystemArchitecture/client-service => ../client-service
	github.com/whitewolf185/SystemArchitecture/domain-service => ../domain-service
)
