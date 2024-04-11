package syncimport

import (
	stdsync "sync"

	"github.com/njupg/moq/pkg/moq/testpackages/syncimport/sync"
)

type Syncer interface {
	Blah(s sync.Thing, wg *stdsync.WaitGroup)
}
