package aggregatorrunner

import (
	"context"

	"k8s.io/apiserver/pkg/authorization/authorizer"

	"github.com/grafana/grafana/pkg/apimachinery/identity"
)

func GetAuthorizer() authorizer.Authorizer {
	return authorizer.AuthorizerFunc(func(ctx context.Context, a authorizer.Attributes) (authorizer.Decision, string, error) {
		_, err := identity.GetRequester(ctx)
		if err != nil {
			return authorizer.DecisionDeny, "", err
		}

		if a.GetResource() == "apiservices" {
			if a.GetVerb() == "get" || a.GetVerb() == "list" {
				return authorizer.DecisionAllow, "", nil
			}
		}

		return authorizer.DecisionDeny, "", nil
	})
}
