package subscriber

import "codebase/core"

//go:generate mockery --name SubscriberUsecase --filename subscriber_usecase.go --output ./mocks
type SubscriberUsecase interface {
	ProcessMessage(ic *core.InternalContext, message any) *core.CustomError
}
