package domain_balancer

import (
	"fmt"
	"testing"

	domain_domain "github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
)

func Test_balancer_prepareQuery(t *testing.T) {

	b := balancer{}
	got := b.prepareQuery(&domain_domain.DeleteCompanionRequest{
		ClientID: "stire",
	})
	fmt.Println(got)
}
