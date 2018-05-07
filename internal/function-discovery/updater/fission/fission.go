package fission

import (
	"github.com/fission/fission"
	"github.com/fission/fission/crd"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/solo-io/gloo/internal/function-discovery/resolver"
	"github.com/solo-io/gloo/pkg/api/types/v1"
	"github.com/solo-io/gloo/pkg/plugins/kubernetes"
	"github.com/solo-io/gloo/pkg/plugins/rest"
)

func GetFuncs(resolve resolver.Resolver, us *v1.Upstream) ([]*v1.Function, error) {
	client, err := getFissionClient()
	if err != nil {
		return nil, err
	}
	fr := NewFissionRetreiver(client)
	return fr.GetFuncs(resolve, us)
}

func IsFissionUpstream(us *v1.Upstream) bool {
	if us.Type != kubernetes.UpstreamTypeKube {
		return false
	}

	spec, err := kubernetes.DecodeUpstreamSpec(us.Spec)
	if err != nil {
		return false
	}

	if spec.ServiceNamespace != "fission" || spec.ServiceName != "router" {
		return false
	}
	return true
}

type retreiver struct {
	fission crd.FissionClient
}

func NewFissionRetreiver(c *crd.FissionClient) *retreiver {
	return &retreiver{fission: c}
}

func (fr *retreiver) listFissionFunctions(ns string) ([]crd.Function, error) {
	l, err := fr.fission.Functions(ns).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return l.Items, nil
}

func (fr *retreiver) GetFuncs(_ resolver.Resolver, us *v1.Upstream) ([]*v1.Function, error) {
	if !IsFissionUpstream(us) {
		return nil, nil
	}

	functions, err := fr.listFissionFunctions("default")
	if err != nil {
		return nil, errors.Wrap(err, "error fetching functions")
	}

	var funcs []*v1.Function

	for _, fn := range functions {
		if fn.Metadata.Name != "" {
			funcs = append(funcs, createFunction(fn))
		}
	}
	return funcs, nil
}

func createFunction(fn crd.Function) *v1.Function {
	headersTemplate := map[string]string{":method": "POST"}

	return &v1.Function{
		Name: fn.Metadata.Name,
		Spec: rest.EncodeFunctionSpec(rest.Template{
			Path:            fission.UrlForFunction(fn.Metadata.Name),
			Header:          headersTemplate,
			PassthroughBody: true,
		}),
	}
}

func getFissionClient() (*crd.FissionClient, error) {
	fissionClient, _, _, err := crd.MakeFissionClient()
	if err != nil {
		return nil, err
	}
	return fissionClient, nil
}
