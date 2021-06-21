package main

import (
	"github.com/dbielecki97/banking/app"
	"github.com/dbielecki97/banking/logger"
)

func main() {
	logger.Info("Starting application...")
	app.Start()
}
