package count

import (
	"fmt"

	"github.com/forbole/bdjuno/v2/database"
	"github.com/forbole/juno/v2/cmd/parse"

	"github.com/spf13/cobra"
)

func NewCountCmd(parseConfig *parse.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "count-missing",
		Short: "get missing blocks and counts",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parse.GetParsingContext(parseConfig)
			if err != nil {
				return err
			}

			// Get all heights from DB
			db := database.Cast(parseCtx.Database)
			var rows []int64
			err = db.Sqlx.Select(&rows, "SELECT height FROM block ORDER BY height")
			if err != nil {
				return err
			}

			// Get latest height
			lastHeight, err := parseCtx.Node.LatestHeight()
			if err != nil {
				return fmt.Errorf("error while getting chain latest block height: %s", err)
			}

			// Get Interval
			var intervals []Inverval
			for index, currHeight := range rows {
				if index == 0 {
					continue
				}
				prevHeight := rows[index-1]
				if currHeight == prevHeight+1 {
					continue
				} else {
					interval := newInterval(prevHeight+1, currHeight-1, currHeight-prevHeight-1)
					intervals = append(intervals, *interval)
				}
			}

			dbLastHeight := rows[len(rows)-1]

			if lastHeight < dbLastHeight {
				interval := newInterval(dbLastHeight+1, lastHeight, lastHeight-dbLastHeight-1)
				intervals = append(intervals, *interval)
			}

			// Print Interval
			for _, interval := range intervals {
				fmt.Printf("Missing Blocks from height %v to %v \n", interval.StartHeight, interval.EndHeight)
				fmt.Println("Missing number of blocks: ", interval.MissingLength)
			}

			return nil
		},
	}
}

type Inverval struct {
	StartHeight   int64
	EndHeight     int64
	MissingLength int64
}

func newInterval(start int64, end int64, length int64) *Inverval {
	return &Inverval{
		StartHeight:   start,
		EndHeight:     end,
		MissingLength: length,
	}
}
