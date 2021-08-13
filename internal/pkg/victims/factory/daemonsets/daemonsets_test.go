package daemonsets

import (
	"testing"

	"kube-monkey/internal/pkg/config"

	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	IDENTIFIER = "kube-monkey-id"
	NAME       = "daemonset_name"
	NAMESPACE  = metav1.NamespaceDefault
)

func newDaemonSet(name string, labels map[string]string) appsv1.DaemonSet {

	return appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: NAMESPACE,
			Labels:    labels,
		},
	}
}

func TestNew(t *testing.T) {

	v1ds := newDaemonSet(
		NAME,
		map[string]string{
			config.IdentLabelKey: IDENTIFIER,
			config.MinbfLabelKey: "1",
		},
	)
	ds, err := New(&v1ds)

	assert.NoError(t, err)
	assert.Equal(t, "v1.DaemonSet", ds.Kind())
	assert.Equal(t, NAME, ds.Name())
	assert.Equal(t, NAMESPACE, ds.Namespace())
	assert.Equal(t, IDENTIFIER, ds.Identifier())
	assert.Equal(t, 1, ds.Minbf())
}

func TestInvalidIdentifier(t *testing.T) {
	v1ds := newDaemonSet(
		NAME,
		map[string]string{
			config.MinbfLabelKey: "1",
		},
	)
	_, err := New(&v1ds)

	assert.Errorf(t, err, "Expected an error if "+config.IdentLabelKey+" label doesn't exist")
}

func TestInvalidMinbf(t *testing.T) {
	v1ds := newDaemonSet(
		NAME,
		map[string]string{
			config.IdentLabelKey: IDENTIFIER,
		},
	)
	_, err := New(&v1ds)

	assert.Errorf(t, err, "Expected an error if "+config.MinbfLabelKey+" label doesn't exist")

	v1ds = newDaemonSet(
		NAME,
		map[string]string{
			config.IdentLabelKey: IDENTIFIER,
			config.MinbfLabelKey: "string",
		},
	)
	_, err = New(&v1ds)

	assert.Errorf(t, err, "Expected an error if "+config.MinbfLabelKey+" label can't be converted a Int type")

	v1ds = newDaemonSet(
		NAME,
		map[string]string{
			config.IdentLabelKey: IDENTIFIER,
			config.MinbfLabelKey: "0",
		},
	)
	_, err = New(&v1ds)

	assert.Errorf(t, err, "Expected an error if "+config.MinbfLabelKey+" label is lower than 1")
}
