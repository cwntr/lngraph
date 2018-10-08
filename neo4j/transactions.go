package neo4j

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/xsb/lngraph/ln"
)

const (
	createTransactionQuery = `CREATE (t:Transaction {
		TxHash: {txHash},
		Amount: {amount},
		NumConfirmations: {numConfirmations},
		BlockHash: {blockHash},
		BlockHeight: {blockHeight},
		TimeStamp: {timeStamp},
		TotalFees: {totalFees}
	} )`

	relTransactionChannelQuery = `MATCH (t:Transaction),(c:Channel)
	WHERE c.ChanPoint STARTS WITH {txHash} AND t.TxHash = {txHash}
	CREATE (t)-[r:FUNDED]->(c)`
)

// CreateTransaction writes a blockchain transaction resource into neo4j.
func CreateTransaction(conn bolt.Conn, tx ln.Transaction) (bolt.Result, error) {
	return conn.ExecNeo(createTransactionQuery, map[string]interface{}{
		"txHash":           tx.TxHash,
		"amount":           tx.Amount,
		"numConfirmations": tx.NumConfirmations,
		"blockHash":        tx.BlockHash,
		"blockHeight":      tx.BlockHeight,
		"timeStamp":        tx.TimeStamp,
		"totalFees":        tx.TotalFees,
	})
}

// CreateTransactionChannelRelationship creates relationship between a
// blockchain transaction and the channel it's part of.
func CreateTransactionChannelRelationship(conn bolt.Conn, tx ln.Transaction) (bolt.Result, error) {
	return conn.ExecNeo(relTransactionChannelQuery, map[string]interface{}{
		"txHash": tx.TxHash,
	})
}