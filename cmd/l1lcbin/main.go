package main

import (
	"github.com/r-ca/l1lc-bin/internal/utils"

	"github.com/samber/do"
)

func main() {
  injector := do.New()

  // Register the logger
  do.Provide(injector, utils.NewLoggerFactory)

  // Get logger(for Bootloader)

  loggerFactory := do.MustInvoke[utils.LoggerFactory](injector)
  logger := loggerFactory.Create("boot")

  logger.Succ("Hello, world!")
}
