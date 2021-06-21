package main

import (
	"github.com/dbielecki97/banking-lib/logger"
	"github.com/dbielecki97/banking/app"
)

func main() {
	logger.Info("Starting application...")
	app.Start()
}
