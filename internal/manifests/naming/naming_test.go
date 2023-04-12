package naming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	for _, test := range []struct {
		name      string
		component string
		expected  string
	}{
		{
			name:      "foo",
			component: "ingester",
			expected:  "tempo-foo-ingester",
		},
		{
			name:     "bar",
			expected: "tempo-bar",
		},
	} {
		t.Run(test.expected, func(t *testing.T) {
			got := Name(test.component, test.name)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestServiceName(t *testing.T) {
	assert.Equal(t, "tempo-simplest-querier-http", ServiceName("simplest", "querier", "http"))
}

func TestServiceFqdn(t *testing.T) {
	assert.Equal(t, "tempo-simplest-querier-http.default.svc.cluster.local", ServiceFqdn("default", "simplest", "querier", "http"))
}

func TestDefaultServiceAccountName(t *testing.T) {
	serviceAccountName := DefaultServiceAccountName("test")
	assert.Equal(t, "tempo-test", serviceAccountName)
}
