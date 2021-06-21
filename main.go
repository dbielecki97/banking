package main

import (
	logger2 "github.com/dbielecki97/banking-lib/logger"
	"github.com/dbielecki97/banking/app"
)

func main() {
	logger2.Info("Starting application...")
	app.Start()
}
