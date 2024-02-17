package deps

import (
	"codebase/core/v1/usecase/auth"
	"codebase/core/v1/usecase/subscriber"
)

func (d *dependency) initService() {
	d.initAuthUsecase()
	d.initSubscriberUsecase()
}

func (d *dependency) initAuthUsecase() {
	d.usecase.AuthUsecase = auth.NewAuthUsecase(
		d.repo.AccountRepository,
		d.repo.CacheRepository,
		d.base.Cfg,
	)
}

func (d *dependency) initSubscriberUsecase() {
	d.usecase.SubscriberUsecase = subscriber.NewSubscriberUsecase()
}
