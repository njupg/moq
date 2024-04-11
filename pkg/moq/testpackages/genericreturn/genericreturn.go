package genericreturn

import "github.com/njupg/moq/pkg/moq/testpackages/genericreturn/otherpackage"

type GenericBar[T any] struct {
	Bar T
}

type IFooBar interface {
	Foobar() GenericBar[otherpackage.Foo]
}
