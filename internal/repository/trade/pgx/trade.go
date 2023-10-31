package pgx

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"time"

	"challenge/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

const workerCount int = 8
const batchSize int = 1000
const queueSize int = 8000
const batchInsertDuration = time.Second * 5 //5 seconds for wait pool to be filled

type Trade interface {
	Create(trade *entity.Trade) error
}
type trade struct {
	db *pgxpool.Pool
}

var queue = make(chan *entity.Trade, queueSize)

func NewTradeRepository(db *pgxpool.Pool) Trade {
	t := &trade{
		db: db,
	}
	for i := 1; i < workerCount; i++ {
		go t.QueueWorker()
	}
	return t
}

// QueueWorker 5 seconds fill a pool for batch inserting
func (r *trade) QueueWorker() {
	for {
		trades := make([]*entity.Trade, 1, batchSize)
		trades[0] = <-queue
		timer := time.NewTimer(batchInsertDuration)
		go func() {
			for {
				if len(trades) == batchSize {
					return
				}
				t := <-queue
				trades = append(trades, t)
			}
		}()
		<-timer.C
		r.BatchInsert(trades)
		return
	}
}

// BatchInsert do batch insert , this function block run time until it batch insert be succesful
func (r *trade) BatchInsert(trades []*entity.Trade) {
	for {
		// Start a new transaction.
		tx, err := r.db.Begin(context.Background())
		if err != nil {
			fmt.Println(err)
			time.Sleep(40 * time.Millisecond)
			continue
		}
		defer tx.Rollback(context.Background())

		// Prepare the bulk INSERT statement.
		stmt := `
		INSERT INTO trade (id, instrumentId, dateEn, open, high, low, close)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

		// Create a batch to store the SQL statements and values.
		batch := &pgx.Batch{}

		// Build the batch of SQL statements.
		for _, trade := range trades {
			batch.Queue(stmt, trade.ID, trade.InstrumentID, trade.DateEn, trade.Open, trade.High, trade.Low, trade.Close)
		}

		// Send the batch of SQL statements to the database.
		results := tx.SendBatch(context.Background(), batch)
		if err := results.Close(); err != nil {
			fmt.Println(err)
			time.Sleep(40 * time.Millisecond)
			continue
		}

		// Commit the transaction to save changes.
		if err := tx.Commit(context.Background()); err != nil {
			fmt.Println(err)
			time.Sleep(40 * time.Millisecond)
			continue
		}
		return
	}
}

func (r *trade) Create(trd *entity.Trade) error {
	queue <- trd
	return nil
}
