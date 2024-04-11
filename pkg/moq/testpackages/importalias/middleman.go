package importalias

import (
	srcclient "github.com/njupg/moq/pkg/moq/testpackages/importalias/source/client"
	tgtclient "github.com/njupg/moq/pkg/moq/testpackages/importalias/target/client"
)

type MiddleMan interface {
	Connect(src srcclient.Client, tgt tgtclient.Client)
}
