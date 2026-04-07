package app

import (
	"os"

	"website-checker/internal/config"
	"website-checker/internal/i18n"
	"website-checker/internal/notification"
	"website-checker/internal/systray"
)

var (
	cfg *config.Config
)

func Run(args []string) {
	cfg, configFilePath, err := config.Load()
	if err != nil {
		notification.Error("Configuration file does not exist: " + err.Error())
		os.Exit(1)
	}

	i18n.Load(cfg.General.Lang)
	config.AppName = i18n.T("app_name")

	notification.Init(cfg)
	notification.SendConfigLoaded()

	systray.Run(cfg, *configFilePath)
}
