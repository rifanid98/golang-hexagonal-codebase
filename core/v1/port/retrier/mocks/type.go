package mocks

import "codebase/core/v1/port/retrier"

type RetrierMock struct {
	Retry retrier.Effector
}
