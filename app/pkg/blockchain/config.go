package blockchain

type Config struct {
	EndPoint                   string
	WebSocket                  string
	TxTimeoutSec               uint64
	CheckTxReceiptTimeMilliSec uint64
	NonceErrRetryCnt           uint8
	DebugMode                  bool
	UserLockEnable             bool
}
