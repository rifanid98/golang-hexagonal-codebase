package xendit

import (
	"codebase/config"
	"codebase/pkg/api"
	"codebase/pkg/util"
)

var log = util.NewLogger()

type xenditApiCallImpl struct {
	client api.HttpDoer
	cfg    config.XenditApiCall
}

func New(client api.HttpDoer, cfg config.XenditApiCall) *xenditApiCallImpl {
	return &xenditApiCallImpl{
		client: client,
		cfg:    cfg,
	}
}
