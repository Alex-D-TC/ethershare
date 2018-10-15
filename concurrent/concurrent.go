package concurrent

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/golang-collections/go-datastructures/queue"
)

type TransactionQueue struct {
	q             *queue.Queue
	cancelRoutine context.CancelFunc
	ctx           context.Context

	logger *log.Logger
}

func MakeTransactionQueue() *TransactionQueue {

	ctx, cancelFunc := context.WithCancel(context.Background())

	tq := TransactionQueue{
		q:             queue.New(16),
		ctx:           ctx,
		cancelRoutine: cancelFunc,
		logger:        log.New(os.Stdout, "Transaction queue: ", log.Ldate|log.Ltime),
	}

	go tq.startWatcher()

	return &tq
}

func (tq *TransactionQueue) startWatcher() {

	done := false

	for {
		trans, err := tq.q.Get(1)
		if err != nil {
			tq.debug(err)
			return
		}

		tq.debug("Processing transaction")

		// Try processing a transaction
		tran := trans[0]
		switch tran.(type) {
		case func() (error, bool):
			err, cont := tran.(func() (error, bool))()
			if err != nil {
				tq.debug(err)
				if cont {
					tq.q.Put(tran)
				}
			}
			break
		}

		// Check for cancel orders
		select {
		case <-tq.ctx.Done():
			done = true
			break
		default:
			break
		}

		if done {
			break
		}

	}

	tq.debug("Watcher stopping")
}

func (tq *TransactionQueue) Submit(transaction func() (error, bool)) error {
	return tq.q.Put(transaction)
}

func (tq *TransactionQueue) Dispose() {
	tq.cancelRoutine()
	tq.q.Dispose()
}

func (tq *TransactionQueue) SetOutput(writer io.Writer) {
	tq.logger = log.New(writer, tq.logger.Prefix(), tq.logger.Flags())
}

func (tq *TransactionQueue) debug(msg interface{}) {
	tq.logger.Println(msg)
}
