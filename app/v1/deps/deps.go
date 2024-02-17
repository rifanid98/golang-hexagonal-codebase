package deps

import (
	"codebase/config"
	"codebase/core/v1/port/account"
	"codebase/core/v1/port/adapter"
	"codebase/core/v1/port/auth"
	"codebase/core/v1/port/broker"
	"codebase/core/v1/port/cache"
	"codebase/core/v1/port/common"
	"codebase/core/v1/port/retrier"
	"codebase/core/v1/port/scheduler"
	"codebase/core/v1/port/subscriber"
	"codebase/infrastructure/v1/broker/gcp"
	"codebase/infrastructure/v1/persistence/mongodb"
	"codebase/infrastructure/v1/persistence/redisdb"
	"codebase/pkg/api"
)

type base struct {
	Cfg   *config.AppConfig
	Mdb   mongodb.Database
	Mdbc  mongodb.Client
	Mdbt  common.Transaction
	Rdbc  redisdb.Client
	Httpc api.HttpDoer
	Rtr   retrier.Retrier
	Schlr scheduler.Scheduler
	Gcpc  gcp.Client
}

type repository struct {
	account.AccountRepository
	cache.CacheRepository
}

type apicall struct {
	adapter.XenditApiCall
}

type usecase struct {
	auth.AuthUsecase
	subscriber.SubscriberUsecase
}

type msgbroker struct {
	broker.Pubsub
}

type dependency struct {
	base    *base
	repo    *repository
	apicall *apicall
	usecase *usecase
	broker  *msgbroker
}

type IDependency interface {
	GetServices() *usecase
	GetRepositories() *repository
	GetBase() *base
	GetBroker() *msgbroker
}

func BuildDependency() *dependency {
	dep := &dependency{
		base:    &base{},
		repo:    &repository{},
		apicall: &apicall{},
		usecase: &usecase{},
		broker:  &msgbroker{},
	}
	dep.initBase()       // execute first
	dep.initRepository() // execute second
	dep.initApiCall()    // execute third
	dep.initBroker()     // execute fourth
	dep.initService()    // execute fifth
	dep.initScheduler()  // execute sixth
	return dep
}

func (d *dependency) GetBase() *base {
	return d.base
}

func (d *dependency) GetServices() *usecase {
	return d.usecase
}

func (d *dependency) GetRepositories() *repository {
	return d.repo
}

func (d *dependency) GetBroker() *msgbroker {
	return d.broker
}
