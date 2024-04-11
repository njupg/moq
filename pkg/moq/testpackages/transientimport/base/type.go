package base

import (
	four "github.com/njupg/moq/pkg/moq/testpackages/transientimport/four/app/v1"
	one "github.com/njupg/moq/pkg/moq/testpackages/transientimport/one/v1"
	"github.com/njupg/moq/pkg/moq/testpackages/transientimport/onev1"
	three "github.com/njupg/moq/pkg/moq/testpackages/transientimport/three/v1"
	two "github.com/njupg/moq/pkg/moq/testpackages/transientimport/two/app/v1"
)

type Transient interface {
	DoSomething(onev1.Zero, one.One, two.Two, three.Three, four.Four)
}
