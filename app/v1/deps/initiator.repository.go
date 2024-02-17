package deps

import (
	mdb "codebase/infrastructure/v1/persistence/mongodb/repository"
	rdb "codebase/infrastructure/v1/persistence/redisdb/repository"
)

func (d *dependency) initRepository() {
	d.initAccountRepository()
	d.initCacheRepository()
}

func (d *dependency) initAccountRepository() {
	d.repo.AccountRepository = mdb.NewAccountRepository(d.base.Mdb, d.base.Cfg)
}

func (d *dependency) initCacheRepository() {
	d.repo.CacheRepository = rdb.NewCacheRepository(d.base.Rdbc)
}
