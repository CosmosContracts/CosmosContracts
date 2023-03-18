package cmd

// DONTCOVER

// from osmosis
// https://github.com/osmosis-labs/osmosis/blob/main/cmd/junod/cmd/forceprune.go

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	tmstore "github.com/tendermint/tendermint/store"
	tmdb "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/tendermint/tendermint/config"
)

const (
	batchMaxSize      = 1000
	validators        = "validatorsKey:"
	consensusParams   = "consensusParamsKey:"
	ABCIResponses     = "abciResponsesKey:"
	fullHeight        = "full_height"
	minHeight         = "min_height"
	defaultFullHeight = "188000"
	defaultMinHeight  = "1000"
)

func forceprune() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "forceprune",
		Short: "Example junod debug forceprune -f 188000 -m 1000, which would keep blockchain and state data of last 188000 blocks (approximately 2 weeks) and ABCI responses of last 1000 blocks.",
		Long:  "Forceprune options prunes and compacts blockstore.db and state.db. One needs to shut down chain before running forceprune. By default it keeps last 188000 blocks (approximately 2 weeks of data) blockstore and state db (validator and consensus information) and 1000 blocks of abci responses from state.db. Everything beyond these heights in blockstore and state.db is pruned. ABCI Responses are stored in index db and so redundant especially if one is running pruned nodes. As a result we are removing ABCI data from state.db aggressively by default. One can override height for blockstore.db and state.db by using -f option and for abci response by using -m option. Example junod forceprune -f 188000 -m 1000.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fullHeightFlag, err := cmd.Flags().GetString(fullHeight)
			if err != nil {
				return err
			}

			minHeightFlag, err := cmd.Flags().GetString(minHeight)
			if err != nil {
				return err
			}

			clientCtx := client.GetClientContextFromCmd(cmd)
			conf := config.DefaultConfig()
			dbPath := clientCtx.HomeDir + "/" + conf.DBPath

			cmdr := exec.Command("junod", "status")
			err = cmdr.Run()

			if err == nil {
				// continue only if throws error
				return nil
			}

			fullHeight, err := strconv.ParseInt(fullHeightFlag, 10, 64)
			if err != nil {
				return err
			}

			minHeight, err := strconv.ParseInt(minHeightFlag, 10, 64)
			if err != nil {
				return err
			}

			startHeight, currentHeight, err := pruneBlockStoreAndGetHeights(dbPath, fullHeight)
			if err != nil {
				return err
			}

			err = compactBlockStore(dbPath)
			if err != nil {
				return err
			}

			err = forcepruneStateStore(dbPath, startHeight, currentHeight, minHeight, fullHeight)
			if err != nil {
				return err
			}
			fmt.Println("Done ...")

			return nil
		},
	}

	cmd.Flags().StringP(fullHeight, "f", defaultFullHeight, "Full height to chop to")
	cmd.Flags().StringP(minHeight, "m", defaultMinHeight, "Min height for ABCI to chop to")
	return cmd
}

// pruneBlockStoreAndGetHeights prunes blockstore and returns the startHeight and currentHeight.
func pruneBlockStoreAndGetHeights(dbPath string, fullHeight int64) (
	startHeight int64, currentHeight int64, err error,
) {
	opts := opt.Options{
		DisableSeeksCompaction: true,
	}

	dbBs, err := tmdb.NewGoLevelDBWithOpts("blockstore", dbPath, &opts)
	if err != nil {
		return 0, 0, err
	}

	defer dbBs.Close()

	bs := tmstore.NewBlockStore(dbBs)
	startHeight = bs.Base()
	currentHeight = bs.Height()

	fmt.Println("Pruning Block Store ...")
	prunedBlocks, err := bs.PruneBlocks(currentHeight - fullHeight)
	if err != nil {
		return 0, 0, err
	}
	fmt.Println("Pruned Block Store ...", prunedBlocks)

	// N.B: We duplicate the call to dbBs.Close() on top of
	// the call in defer statement above to make sure that the resources
	// are properly released and any potential error from Close()
	// is handled. Close() should be idempotent so this is acceptable.
	if err := dbBs.Close(); err != nil {
		return 0, 0, err
	}

	return startHeight, currentHeight, nil
}

// compactBlockStore compacts block storage.
func compactBlockStore(dbPath string) (err error) {
	compactOpts := opt.Options{
		DisableSeeksCompaction: true,
	}

	fmt.Println("Compacting Block Store ...")

	db, err := leveldb.OpenFile(dbPath+"/blockstore.db", &compactOpts)
	defer func() {
		err = db.Close()
	}()
	if err != nil {
		return err
	}
	err = db.CompactRange(*util.BytesPrefix([]byte{}))
	return err
}

// forcepruneStateStore prunes and compacts state storage.
func forcepruneStateStore(dbPath string, startHeight, currentHeight, minHeight, fullHeight int64) error {
	opts := opt.Options{
		DisableSeeksCompaction: true,
	}

	db, err := leveldb.OpenFile(dbPath+"/state.db", &opts)
	if err != nil {
		return err
	}
	defer db.Close()

	stateDBKeys := []string{validators, consensusParams, ABCIResponses}
	fmt.Println("Pruning State Store ...")
	for i, s := range stateDBKeys {
		fmt.Println(i, s)

		retainHeight := int64(0)
		if s == ABCIResponses {
			retainHeight = currentHeight - minHeight
		} else {
			retainHeight = currentHeight - fullHeight
		}

		batch := new(leveldb.Batch)
		curBatchSize := uint64(0)

		fmt.Println(startHeight, currentHeight, retainHeight)

		for c := startHeight; c < retainHeight; c++ {
			batch.Delete([]byte(s + strconv.FormatInt(c, 10)))
			curBatchSize++

			if curBatchSize%batchMaxSize == 0 && curBatchSize > 0 {
				err := db.Write(batch, nil)
				if err != nil {
					return err
				}
				batch.Reset()
				batch = new(leveldb.Batch)
			}
		}

		err := db.Write(batch, nil)
		if err != nil {
			return err
		}
		batch.Reset()
	}

	fmt.Println("Compacting State Store ...")
	err = db.CompactRange(*util.BytesPrefix([]byte{}))

	return err
}
