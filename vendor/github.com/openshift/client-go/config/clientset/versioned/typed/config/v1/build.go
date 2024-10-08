// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/openshift/api/config/v1"
	configv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	scheme "github.com/openshift/client-go/config/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// BuildsGetter has a method to return a BuildInterface.
// A group's client should implement this interface.
type BuildsGetter interface {
	Builds() BuildInterface
}

// BuildInterface has methods to work with Build resources.
type BuildInterface interface {
	Create(ctx context.Context, build *v1.Build, opts metav1.CreateOptions) (*v1.Build, error)
	Update(ctx context.Context, build *v1.Build, opts metav1.UpdateOptions) (*v1.Build, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Build, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.BuildList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Build, err error)
	Apply(ctx context.Context, build *configv1.BuildApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Build, err error)
	BuildExpansion
}

// builds implements BuildInterface
type builds struct {
	*gentype.ClientWithListAndApply[*v1.Build, *v1.BuildList, *configv1.BuildApplyConfiguration]
}

// newBuilds returns a Builds
func newBuilds(c *ConfigV1Client) *builds {
	return &builds{
		gentype.NewClientWithListAndApply[*v1.Build, *v1.BuildList, *configv1.BuildApplyConfiguration](
			"builds",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.Build { return &v1.Build{} },
			func() *v1.BuildList { return &v1.BuildList{} }),
	}
}
