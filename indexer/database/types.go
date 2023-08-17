package database

type ERC20TransferEvent struct {
	BlockNumber int64
	TxHash      string
	FromAddress string
	ToAddress   string
	Value       string
}
