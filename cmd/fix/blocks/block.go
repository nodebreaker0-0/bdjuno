package blocks

import (
	"fmt"
	"strconv"

	"github.com/forbole/juno/v2/cmd/parse"

	"github.com/forbole/juno/v2/parser"
	"github.com/spf13/cobra"
)

// blocksCmd returns a Cobra command that allows to fix missing blocks in database
func blockCmd(parseConfig *parse.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "heights [start height] [end height]",
		Short: "Fix missing block and transaction in database for a range of heights",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parse.GetParsingContext(parseConfig)
			if err != nil {
				return err
			}

			workerCtx := parser.NewContext(parseCtx.EncodingConfig.Marshaler, nil, parseCtx.Node, parseCtx.Database, parseCtx.Logger, parseCtx.Modules)
			worker := parser.NewWorker(0, workerCtx)

			startHeight, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("error while converting start height %s to integer: %s", args[0], err)
			}
			endHeight, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("error while converting end height %s to integer: %s", args[1], err)
			}

			for k := startHeight; k <= endHeight; k++ {
				fmt.Printf("Refetching missing blocks and transactions for height %v ... \n", k)
				err = worker.ForceProcess(k)
				if err != nil {
					return fmt.Errorf("error while re-fetching block %d: %s", k, err)
				}
			}

			return nil
		},
	}
}
