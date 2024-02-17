package mocks

import "codebase/core"

type PubsubMock struct {
	Publish   *core.CustomError
	Subscribe *core.CustomError
}
