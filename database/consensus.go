package database

import (
	"fmt"
	"time"

	tmtypes "github.com/tendermint/tendermint/types"

	dbtypes "github.com/forbole/bdjuno/database/types"
	constypes "github.com/forbole/bdjuno/x/consensus/types"
)

// SaveConsensus allows to properly store the given consensus event into the database.
// Note that only one consensus event is allowed inside the database at any time.
func (db *BigDipperDb) SaveConsensus(event *constypes.ConsensusEvent) error {
	// Delete all the existing events
	stmt := `DELETE FROM consensus WHERE true`
	if _, err := db.Sql.Exec(stmt); err != nil {
		return err
	}

	stmt = `INSERT INTO consensus (height, round, step) VALUES ($1, $2, $3)`
	_, err := db.Sql.Exec(stmt, event.Height, event.Round, event.Step)
	return err
}

// GetLastBlock returns the last block stored inside the database based on the heights
func (db *BigDipperDb) GetLastBlock() (*dbtypes.BlockRow, error) {
	stmt := `SELECT * FROM block ORDER BY height DESC LIMIT 1`

	var blocks []dbtypes.BlockRow
	if err := db.Sqlx.Select(&blocks, stmt); err != nil {
		return nil, err
	}

	if len(blocks) == 0 {
		return nil, fmt.Errorf("cannot get block, no blocks saved")
	}

	return &blocks[0], nil
}

// GetLastBlockHeight returns the last block height stored inside the database
func (db *BigDipperDb) GetLastBlockHeight() (int64, error) {
	block, err := db.GetLastBlock()
	if err != nil {
		return 0, err
	}
	return block.Height, nil
}

// getBlockHeightTime retrieves the block at the specific time
func (db *BigDipperDb) getBlockHeightTime(pastTime time.Time) (dbtypes.BlockRow, error) {
	stmt := `SELECT * FROM block WHERE block.timestamp <= $1 ORDER BY block.timestamp DESC LIMIT 1;`

	var val []dbtypes.BlockRow
	if err := db.Sqlx.Select(&val, stmt, pastTime); err != nil {
		return dbtypes.BlockRow{}, err
	}

	if len(val) == 0 {
		return dbtypes.BlockRow{}, fmt.Errorf("cannot get block time, no blocks saved")
	}

	return val[0], nil
}

// GetBlockHeightTimeMinuteAgo return block height and time that a block proposals
// about a minute ago from input date
func (db *BigDipperDb) GetBlockHeightTimeMinuteAgo(now time.Time) (dbtypes.BlockRow, error) {
	pastTime := now.Add(time.Minute * -1)
	return db.getBlockHeightTime(pastTime)
}

// GetBlockHeightTimeHourAgo return block height and time that a block proposals
// about a hour ago from input date
func (db *BigDipperDb) GetBlockHeightTimeHourAgo(now time.Time) (dbtypes.BlockRow, error) {
	pastTime := now.Add(time.Hour * -1)
	return db.getBlockHeightTime(pastTime)
}

// GetBlockHeightTimeDayAgo return block height and time that a block proposals
// about a day (24hour) ago from input date
func (db *BigDipperDb) GetBlockHeightTimeDayAgo(now time.Time) (dbtypes.BlockRow, error) {
	pastTime := now.Add(time.Hour * -24)
	return db.getBlockHeightTime(pastTime)
}

// SaveAverageBlockTimeGenesis save the average block time in average_block_time_from_genesis table
func (db *BigDipperDb) SaveAverageBlockTimeGenesis(averageTime float64, height int64) error {
	stmt := `
INSERT INTO average_block_time_from_genesis(average_time ,height) 
VALUES ($1, $2) ON CONFLICT (height) DO UPDATE SET average_time = excluded.average_time`
	_, err := db.Sqlx.Exec(stmt, averageTime, height)
	return err
}

// SaveAverageBlockTimePerMin save the average block time in average_block_time_per_minute table
func (db *BigDipperDb) SaveAverageBlockTimePerMin(averageTime float64, height int64) error {
	stmt := `
INSERT INTO average_block_time_per_minute(average_time, height) 
VALUES ($1, $2) ON CONFLICT (height) DO UPDATE SET average_time = excluded.average_time`
	_, err := db.Sqlx.Exec(stmt, averageTime, height)
	return err
}

// SaveAverageBlockTimePerHour save the average block time in average_block_time_per_hour table
func (db *BigDipperDb) SaveAverageBlockTimePerHour(averageTime float64, height int64) error {
	stmt := `
INSERT INTO average_block_time_per_hour(average_time, height) 
VALUES ($1, $2) ON CONFLICT (height) DO UPDATE SET average_time = excluded.average_time`
	_, err := db.Sqlx.Exec(stmt, averageTime, height)
	return err
}

// SaveAverageBlockTimePerDay save the average block time in average_block_time_per_day table
func (db *BigDipperDb) SaveAverageBlockTimePerDay(averageTime float64, height int64) error {
	stmt := `
INSERT INTO average_block_time_per_day(average_time, height) 
VALUES ($1, $2) ON CONFLICT (height) DO UPDATE SET average_time = excluded.average_time`
	_, err := db.Sqlx.Exec(stmt, averageTime, height)
	return err
}

// SaveGenesisHeight save the genesis height
func (db *BigDipperDb) SaveGenesisData(genesis *tmtypes.GenesisDoc) error {
	stmt := `DELETE FROM genesis WHERE TRUE`
	_, err := db.Sqlx.Exec(stmt)
	if err != nil {
		return err
	}
	stmt = `INSERT INTO genesis(time, chain_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err = db.Sqlx.Exec(stmt, genesis.GenesisTime, genesis.ChainID)
	return err
}

// GetGenesisTime get genesis time of chain (only work if x/consensus enabled)
func (db *BigDipperDb) GetGenesisTime() (time.Time, error) {
	var rows []*dbtypes.GenesisRow
	err := db.Sqlx.Select(&rows, `SELECT * FROM genesis;`)
	if err != nil || len(rows) == 0 {
		return time.Time{}, err
	}
	return rows[0].Time, nil
}
