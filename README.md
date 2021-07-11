# servicebinding
> **Kubernetes Service Binding Library for Go Applications**

[![Go Reference](https://pkg.go.dev/badge/github.com/baijum/servicebinding.svg)](https://pkg.go.dev/github.com/baijum/servicebinding)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![CI](https://github.com/baijum/servicebinding/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/baijum/servicebinding/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/baijum/servicebinding)](https://goreportcard.com/report/github.com/baijum/servicebinding)

The [Service Binding Specification][spec] for Kubernetes standardizes exposing
backing service secrets to applications.  This project provides a Go package to
consume the bindings projected into the container.  The [Application
Projection][application-projection] section of the spec describes how the
bindings are projected into the application.  The primary mechanism of
projection is through files mounted at a specific directory.  The bindings
directory path can be discovered through `SERVICE_BINDING_ROOT` environment
variable.  The operator must have injected `SERVICE_BINDING_ROOT` environment to
all the containers where bindings are created.

Within this service binding root directory, there could be multiple bindings
projected from different Service Bindings.  For example, suppose an application
requires to connect to a database and cache server. In that case, one Service
Binding can provide the database, and the other Service Binding can offer
bindings to the cache server.

Let's take a look at the example given in the spec:

```
$SERVICE_BINDING_ROOT
├── account-database
│   ├── type
│   ├── provider
│   ├── uri
│   ├── username
│   └── password
└── transaction-event-stream
    ├── type
    ├── connection-count
    ├── uri
    ├── certificates
    └── private-key
```

In the above example, there are two bindings under the `SERVICE_BINDING_ROOT`
directory.  The `account-database` and `transaction-event-system` are the names
of the bindings.  Files within each bindings directory has a special file named
`type`, and you can rely on the value of that file to identify the type of the
binding projected into that directory.  In certain directories, you can also see
another file named `provider`.  The provider is an additional identifier to
narrow down the type further.  This package use the `type` field and,
optionally, `provider` field to look up the bindings.

## Installation

You can download the pacakge using `go get`:

```bash
go get github.com/baijum/servicebinding
```

## Usage

You can import the `binding` package like this:

```go
import "github.com/baijum/servicebinding/binding"
```

You can create `*ServiceBinding` object like this:

```go
sb, err := NewServiceBinding()
```

To get bindings for a specific `type`, say `postgresql`:

```go
bindings, err := sb.Bindings("postgresql")
```

To get bindings for a specific `type`, say `mysql`, and `provider`, say `mariadb`:

```go
bindings, err := sb.BindingsWithProvider("mysql", "mariadb")
```

To get all bindings irrespective of the `type` and `provider`:

```go
bindings, err := sb.AllBindings()
```

This is the complete `ServiceBinding` interface:

```go
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
```

[spec]: https://github.com/k8s-service-bindings/spec
[application-projection]: https://github.com/k8s-service-bindings/spec#application-projection
