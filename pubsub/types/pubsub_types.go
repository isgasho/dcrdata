package types

import (
	"strings"

	exptypes "github.com/decred/dcrdata/v4/explorer/types"
)

// ErrWsClosed is the error message text used websocket Conn.Close tries to
// close an already closed connection.
var ErrWsClosed = "use of closed network connection"

// IsWSClosedErr checks if the passed error indicates a closed websocket
// connection.
func IsWSClosedErr(err error) (closedErr bool) {
	if err != nil && strings.Contains(err.Error(), ErrWsClosed) {
		closedErr = true
	}
	return
}

// WebSocketMessage represents the JSON object used to send and receive typed
// messages to the web client.
type WebSocketMessage struct {
	EventId string `json:"event"`
	Message string `json:"message"`
}

type TxList []*exptypes.MempoolTx

type HubSignal int

const (
	SigSubscribe HubSignal = iota
	SigUnsubscribe
	SigNewBlock
	SigMempoolUpdate
	SigPingAndUserCount
	SigNewTx
	SigSyncStatus
)

var Subscriptions = map[string]HubSignal{
	"newblock":       SigNewBlock,
	"mempool":        SigMempoolUpdate,
	"ping":           SigPingAndUserCount,
	"newtx":          SigNewTx,
	"blockchainSync": SigSyncStatus,
}

// Event type field for an event.
var eventIDs = map[HubSignal]string{
	SigSubscribe:        "subscribe",
	SigUnsubscribe:      "unsubscribe",
	SigNewBlock:         "newblock",
	SigMempoolUpdate:    "mempool",
	SigPingAndUserCount: "ping",
	SigNewTx:            "newtx",
	SigSyncStatus:       "blockchainSync",
}

func (s HubSignal) String() string {
	str, found := eventIDs[s]
	if !found {
		return "invalid"
	}
	return str
}

func (s HubSignal) IsValid() bool {
	_, found := eventIDs[s]
	return found
}

// var (
// 	SigNewBlock         = "newblock" // exptypes.WebsocketBlock
// 	SigMempoolUpdate    = "mempool"  // exptypes.MempoolShort
// 	SigPingAndUserCount = "ping"     // string (number of connected clients)
// 	SigNewTx            = "newtx"    // TxList a.k.a. []*exptypes.MempoolTx
// 	//SigSyncStatus       = "blockchainSync"
// )
