module github.com/whitewolf185/SystemArchitecture/gateway

go 1.22.1

require (
	github.com/ggicci/httpin v0.18.1
	github.com/go-chi/chi v1.5.5
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.9.3
	github.com/swaggo/http-swagger/v2 v2.0.2
	github.com/swaggo/swag v1.16.3
	github.com/whitewolf185/SystemArchitecture/client-service v1.0.0
	github.com/whitewolf185/SystemArchitecture/domain-service v1.0.0
	go.mongodb.org/mongo-driver v1.15.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/ggicci/owl v0.8.2 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/spec v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/swaggo/files/v2 v2.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.19.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/whitewolf185/SystemArchitecture/client-service => ../client-service
	github.com/whitewolf185/SystemArchitecture/domain-service => ../domain-service
)
