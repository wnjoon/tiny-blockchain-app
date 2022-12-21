package config

type Config struct {
	Endpoint string
}

// func (c Config) BlockChain() blockchain.Config {
// 	path := "blockchain"

// 	return blockchain.Config{
// 		EndPoint:                   c.viper.GetString(path + ".endPoint"),
// 		WebSocket:                  c.viper.GetString(path + ".websocket"),
// 		TxTimeoutSec:               c.viper.GetUint64(path + ".txTimeoutSec"),
// 		CheckTxReceiptTimeMilliSec: c.viper.GetUint64(path + ".checkTxReceiptTimeMilliSec"),
// 		NonceErrRetryCnt:           uint8(c.viper.GetUint(path + ".nonceErrRetryCnt")),
// 		DebugMode:                  c.viper.GetBool(path + ".debugMode"),
// 		UserLockEnable:             c.viper.GetBool(path + ".userLockEnable"),
// 	}
// }
