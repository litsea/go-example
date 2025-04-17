package complete

import (
	"github.com/spf13/cobra"

	"github.com/litsea/go-example/cmd/complete/service"
	"github.com/litsea/go-example/config"
)

var cmd *cobra.Command

func New() *cobra.Command {
	cmd = &cobra.Command{
		Use:   "complete",
		Short: "complete go example",
		Run: func(cmd *cobra.Command, args []string) {
			service.New(config.V())
		},
	}

	return cmd
}
