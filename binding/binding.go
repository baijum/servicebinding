package binding

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	// serviceBindingRoot is the environment variable pointing to the root of all bindings
	// Ref. https://github.com/k8s-service-bindings/spec#application-projection
	serviceBindingRoot = "SERVICE_BINDING_ROOT"
)

// ServiceBinding represents the bindings projected into an application
// through an implementation of Service Binding Specification for Kubernetes
type ServiceBinding struct {
	root string
}

// NewServiceBinding returns the ServiceBinding object
// Error is returned when SERVICE_BINDING_ROOT
// environment variable is not set
func NewServiceBinding() (*ServiceBinding, error) {
	root, exists := os.LookupEnv(serviceBindingRoot)
	if !exists {
		return nil, errors.New("environment variable not set: SERVICE_BINDING_ROOT")
	}
	sb := &ServiceBinding{root: root}
	return sb, nil
}

// AllBindings get all bndings as a sice of map[string]string.
// Return empty slice if no bindings found.
func (sb *ServiceBinding) AllBindings() ([]map[string]string, error) {
	result := []map[string]string{}
	dirs, err := os.ReadDir(sb.root)
	if err != nil {
		return nil, err
	}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		m := map[string]string{}
		files, err := os.ReadDir(filepath.Join(sb.root, d.Name()))
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			if f.IsDir() {
				continue
			}
			fc, err := os.ReadFile(filepath.Join(sb.root, d.Name(), f.Name()))
			if err != nil {
				return nil, err
			}
			m[f.Name()] = string(fc)
		}
		result = append(result, m)
	}
	return result, nil
}

// Bindings get the bindings for a given type as a slice of
// map[string]string.
// Return empty slice if no bindings found
func (sb *ServiceBinding) Bindings(_type string) ([]map[string]string, error) {
	result := []map[string]string{}
	dirs, err := os.ReadDir(sb.root)
	if err != nil {
		return nil, err
	}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}

		if _, err := os.Stat(filepath.Join(sb.root, d.Name(), "type")); err == nil {

			tp, err := os.ReadFile(filepath.Join(sb.root, d.Name(), "type"))
			if err != nil {
				continue
			}
			if string(tp) != _type {
				continue
			}

			m := map[string]string{}
			files, err := os.ReadDir(filepath.Join(sb.root, d.Name()))
			if err != nil {
				return nil, err
			}
			for _, f := range files {
				if f.IsDir() {
					continue
				}
				fc, err := os.ReadFile(filepath.Join(sb.root, d.Name(), f.Name()))
				if err != nil {
					return nil, err
				}
				m[f.Name()] = string(fc)
			}
			result = append(result, m)
		}
	}
	return result, nil
}

// BindingsWithProvider get the bindings for a given type and provider
// as a slice of map[string]string.
// Return empty slice if no bindings found.
func (sb *ServiceBinding) BindingsWithProvider(_type, provider string) ([]map[string]string, error) {
	result := []map[string]string{}
	dirs, err := os.ReadDir(sb.root)
	if err != nil {
		return nil, err
	}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}

		if _, err := os.Stat(filepath.Join(sb.root, d.Name(), "type")); err == nil {

			tp, err := os.ReadFile(filepath.Join(sb.root, d.Name(), "type"))
			if err != nil {
				continue
			}
			if string(tp) != _type {
				continue
			}

			if _, err := os.Stat(filepath.Join(sb.root, d.Name(), "provider")); err == nil {

				pp, err := os.ReadFile(filepath.Join(sb.root, d.Name(), "provider"))
				if err != nil {
					continue
				}
				if string(pp) != provider {
					continue
				}

				m := map[string]string{}
				files, err := os.ReadDir(filepath.Join(sb.root, d.Name()))
				if err != nil {
					return nil, err
				}
				for _, f := range files {
					if f.IsDir() {
						continue
					}
					fc, err := os.ReadFile(filepath.Join(sb.root, d.Name(), f.Name()))
					if err != nil {
						return nil, err
					}
					m[f.Name()] = string(fc)
				}
				result = append(result, m)
			}
		}
	}
	return result, nil
}
