package cron

import (
	"time"

	"codebase/core"
	"codebase/pkg/util"
	"github.com/go-co-op/gocron"
)

var log = util.NewLogger()

type schedulerImpl struct {
	S *gocron.Scheduler
}

func New() *schedulerImpl {
	location, _ := time.LoadLocation("Asia/Jakarta")
	s := gocron.NewScheduler(location)
	return &schedulerImpl{
		S: s,
	}
}

func (s *schedulerImpl) Start(ic *core.InternalContext) *core.CustomError {
	log.Info(ic.ToContext(), "cron start...")

	// TODO:
	// - check status qr code
	// - jika sukses maka update field is_verified & verified_at
	// - jika sukses maka update metadata dengan data yang baru
	res, err := s.S.Every(2).Seconds().Do(func() { log.Info(ic.ToContext(), "cron executed...") }, ic)
	if err != nil {
		log.Error(ic.ToContext(), "failed s.S.Every(2).Seconds().Do(func() { log.Info(ic.ToContext(), \"cron executed...\") }, ic)", err)
		return &core.CustomError{
			Code:    core.INTERNAL_SERVER_ERROR,
			Message: err.Error(),
		}
	}
	log.Info(ic.ToContext(), "cron info", res.Error())
	log.Info(ic.ToContext(), "cron info", res.NextRun())
	s.S.StartAsync()

	return nil
}
