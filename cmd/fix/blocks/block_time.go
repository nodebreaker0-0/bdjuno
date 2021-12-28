package blocks

import (
	"fmt"
	"strconv"

	"github.com/forbole/bdjuno/v2/database"
	"github.com/forbole/juno/v2/cmd/parse"

	"github.com/spf13/cobra"
)

func blockTimeCmd(parseConfig *parse.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "blocktime",
		Short: "Fix missing blocks and transactions in database from the start height",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parse.GetParsingContext(parseConfig)
			if err != nil {
				return err
			}

			startHeight, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("error while converting string to int on start height: %s", err)
			}

			lastHeight, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("error while converting string to int on last height: %s", err)
			}

			// Prepare DB
			db := database.Cast(parseCtx.Database)

			k := int64(startHeight)
			fmt.Printf("Refetching block time from height %d ... \n", k)
			for ; k <= int64(lastHeight); k++ {
				block, err := parseCtx.Node.Block(k)
				if err != nil {
					return fmt.Errorf("error while getting block result %d: %s", k, err)
				}

				err = db.UpdateBlockTime(k, block.Block.Time)
				if err != nil {
					return fmt.Errorf("error while updating block time: %s", err)
				}

				fmt.Printf("updated block time of height %d ... \n", k)

			}

			return nil
		},
	}
}
