package l1lcbin

import (
	"github.com/r-ca/l1lc-bin/internal/utils"

	"github.com/samber/do"
)

func main() {
  injector := do.New()

  // Register the logger
  do.Provide(injector, utils.NewLoggerFactory)
}
