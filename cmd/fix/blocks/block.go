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
		Use:   "refetch [block heights]",
		Short: "Fix missing block and transaction in database for certain heights",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parse.GetParsingContext(parseConfig)
			if err != nil {
				return err
			}

			workerCtx := parser.NewContext(parseCtx.EncodingConfig.Marshaler, nil, parseCtx.Node, parseCtx.Database, parseCtx.Logger, parseCtx.Modules)
			worker := parser.NewWorker(0, workerCtx)

			for _, k := range args {
				k, err := strconv.ParseInt(k, 10, 64)
				if err != nil {
					return fmt.Errorf("error while converting block height %d to integer: %s", k, err)
				}

				fmt.Printf("Refetching missing blocks and transactions for height %v ... \n", k)
				err = worker.Process(k)
				if err != nil {
					return fmt.Errorf("error while re-fetching block %d: %s", k, err)
				}
			}

			return nil
		},
	}
}
