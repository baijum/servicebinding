package v1beta1

// ServiceBinding retrieve bindings from a file-system created through an
// implementation of Service Binding Specification for Kubernetes
// https://github.com/k8s-service-bindings/spec
type ServiceBinding interface {

	// AllBindings get all bndings as a sice of map[string]string.
	// Return empty slice if no bindings found.
	AllBindings() ([]map[string]string, error)

	// Bindings get the bindings for a given type as a slice of
	// map[string]string.
	// Return empty slice if no bindings found
	Bindings(_type string) ([]map[string]string, error)

	// BindingsWithProvider get the bindings for a given type and provider
	// as a slice of map[string]string.
	// Return empty slice if no bindings found.
	BindingsWithProvider(_type, provider string) ([]map[string]string, error)
}
