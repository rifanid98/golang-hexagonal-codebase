package subscriber

import (
	"codebase/core"
	"codebase/pkg/util"
)

var log = util.NewLogger()

type subscriberUsecaseImpl struct{}

func NewSubscriberUsecase() *subscriberUsecaseImpl {
	return &subscriberUsecaseImpl{}
}

func (s *subscriberUsecaseImpl) ProcessMessage(ic *core.InternalContext, msg any) *core.CustomError {
	panic("not yet implemented")
}
