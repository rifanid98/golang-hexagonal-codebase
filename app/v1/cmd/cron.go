package cmd

import (
	"os"
	"os/signal"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"codebase/app/v1/deps"
	"codebase/core"
)

var cron = &cobra.Command{
	Use:   "cron",
	Short: "start cron",
	Run: func(cmd *cobra.Command, args []string) {
		deps := deps.BuildDependency()
		ic := core.NewInternalContext(uuid.NewString())

		go func() {
			log.Info(ic.ToContext(), "cron running...")

			cerr := deps.GetBase().Schlr.Start(ic)
			if cerr != nil {
				log.Error(ic.ToContext(), "cron failed to start", cerr)
			}

			log.Info(ic.ToContext(), "cron finished...")
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		log.Info(ic.ToContext(), "cron shutting down...")
	},
}
