// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	account.proto
	blockchain.proto
	common.proto
	config.proto
	db.proto
	executor.proto
	p2p.proto
	pbft.proto
	rpc.proto
	transaction.proto
	wallet.proto

It has these top-level messages:
	Account
	ReceiptExecAccountTransfer
	ReceiptAccountTransfer
	ReqBalance
	Accounts
	ReqTokenBalance
	TotalFee
	ReqGetTotalCoins
	ReplyGetTotalCoins
	IterateRangeByStateHash
	ReqAccountTokenAssets
	TokenAsset
	ReplyAccountTokenAssets
	Header
	Block
	Blocks
	BlockPid
	BlockDetails
	Headers
	HeadersPid
	BlockOverview
	BlockDetail
	Receipts
	ReceiptCheckTxList
	ChainStatus
	ReqBlocks
	MempoolSize
	ReplyBlockHeight
	BlockBody
	IsCaughtUp
	IsNtpClockSync
	BlockChainQuery
	Reply
	ReqString
	ReplyString
	ReplyStrings
	ReqInt
	Int64
	ReqHash
	ReplyHash
	ReqNil
	ReqHashes
	ReplyHashes
	KeyValue
	Config
	Log
	MemPool
	Consensus
	Wallet
	Store
	BlockChain
	P2P
	Rpc
	Exec
	Authority
	LeafNode
	InnerNode
	MAVLProof
	StoreNode
	LocalDBSet
	LocalDBList
	LocalDBGet
	LocalReplyValue
	StoreSet
	StoreSetWithSync
	StoreGet
	StoreReplyValue
	Genesis
	CoinsAction
	CoinsGenesis
	CoinsTransfer
	CoinsWithdraw
	Hashlock
	HashlockAction
	HashlockLock
	HashlockUnlock
	HashlockSend
	HashRecv
	Hashlockquery
	Ticket
	TicketAction
	TicketMiner
	TicketBind
	TicketOpen
	TicketGenesis
	TicketClose
	TicketList
	TicketInfos
	ReplyTicketList
	ReplyWalletTickets
	ReceiptTicket
	ReceiptTicketBind
	ExecTxList
	Query
	Norm
	NormAction
	NormPut
	RetrievePara
	Retrieve
	RetrieveAction
	BackupRetrieve
	PreRetrieve
	PerformRetrieve
	CancelRetrieve
	ReqRetrieveInfo
	RetrieveQuery
	TokenAction
	TokenPreCreate
	TokenFinishCreate
	TokenRevokeCreate
	Token
	ReqTokens
	ReplyTokens
	ReceiptToken
	Trade
	TradeForSell
	SellOrder
	SellOrderReceipt
	ReqAddrTokens
	ReplySellOrders
	TokenRecv
	ReplyAddrRecvForTokens
	TradeForBuy
	TradeForRevokeSell
	ReceiptTradeBase
	ReceiptTradeSell
	ReceiptTradeBuy
	ReceiptTradeRevoke
	TradeBuyDone
	ReplyTradeBuyOrders
	ArrayConfig
	StringConfig
	Int32Config
	ConfigItem
	ModifyConfig
	ManageAction
	ReceiptConfig
	ReplyConfig
	ReqTokenSellOrder
	P2PGetPeerInfo
	P2PPeerInfo
	P2PVersion
	P2PVerAck
	P2PPing
	P2PPong
	P2PGetAddr
	P2PAddr
	P2PExternalInfo
	P2PGetBlocks
	P2PGetMempool
	P2PInv
	Inventory
	P2PGetData
	P2PTx
	P2PBlock
	BroadCastData
	P2PGetHeaders
	P2PHeaders
	InvData
	InvDatas
	Peer
	PeerList
	NodeNetInfo
	Operation
	Checkpoint
	Entry
	ViewChange
	Summary
	Result
	Request
	RequestClient
	RequestPrePrepare
	RequestPrepare
	RequestCommit
	RequestCheckpoint
	RequestViewChange
	RequestAck
	RequestNewView
	ClientReply
	CreateTx
	UnsignTx
	SignedTx
	Transaction
	Signature
	AddrOverview
	ReqAddr
	HexTx
	ReplyTxInfo
	ReqTxList
	ReplyTxList
	TxHashList
	ReplyTxInfos
	ReceiptLog
	Receipt
	ReceiptData
	TxResult
	TransactionDetail
	TransactionDetails
	ReqAddrs
	WalletTxDetail
	WalletTxDetails
	WalletAccountStore
	WalletPwHash
	WalletStatus
	WalletAccounts
	WalletAccount
	WalletUnLock
	GenSeedLang
	GetSeedByPw
	SaveSeedByPw
	ReplySeed
	ReqWalletSetPasswd
	ReqNewAccount
	MinerFlag
	ReqWalletTransactionList
	ReqWalletImportPrivKey
	ReqWalletSendToAddress
	ReqWalletSetFee
	ReqWalletSetLabel
	ReqWalletMergeBalance
	ReqStr
	ReplyStr
	ReqTokenPreCreate
	ReqTokenFinishCreate
	ReqTokenRevokeCreate
	ReqSellToken
	ReqBuyToken
	ReqRevokeSell
	ReqModifyConfig
	ReqSignRawTx
	ReplySignRawTx
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Account 的信息
type Account struct {
	// coins标识，目前只有0 一个值
	Currency int32 `protobuf:"varint,1,opt,name=currency" json:"currency,omitempty"`
	// 账户可用余额
	Balance int64 `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
	// 账户冻结余额
	Frozen int64 `protobuf:"varint,3,opt,name=frozen" json:"frozen,omitempty"`
	// 账户的地址
	Addr string `protobuf:"bytes,4,opt,name=addr" json:"addr,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetCurrency() int32 {
	if m != nil {
		return m.Currency
	}
	return 0
}

func (m *Account) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetFrozen() int64 {
	if m != nil {
		return m.Frozen
	}
	return 0
}

func (m *Account) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

// 账户余额改变的一个交易回报（合约内）
type ReceiptExecAccountTransfer struct {
	// 合约地址
	ExecAddr string `protobuf:"bytes,1,opt,name=execAddr" json:"execAddr,omitempty"`
	// 转移前
	Prev *Account `protobuf:"bytes,2,opt,name=prev" json:"prev,omitempty"`
	// 转移后
	Current *Account `protobuf:"bytes,3,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptExecAccountTransfer) Reset()                    { *m = ReceiptExecAccountTransfer{} }
func (m *ReceiptExecAccountTransfer) String() string            { return proto.CompactTextString(m) }
func (*ReceiptExecAccountTransfer) ProtoMessage()               {}
func (*ReceiptExecAccountTransfer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReceiptExecAccountTransfer) GetExecAddr() string {
	if m != nil {
		return m.ExecAddr
	}
	return ""
}

func (m *ReceiptExecAccountTransfer) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptExecAccountTransfer) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

// 账户余额改变的一个交易回报（coins内）
type ReceiptAccountTransfer struct {
	// 转移前
	Prev *Account `protobuf:"bytes,1,opt,name=prev" json:"prev,omitempty"`
	// 转移后
	Current *Account `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptAccountTransfer) Reset()                    { *m = ReceiptAccountTransfer{} }
func (m *ReceiptAccountTransfer) String() string            { return proto.CompactTextString(m) }
func (*ReceiptAccountTransfer) ProtoMessage()               {}
func (*ReceiptAccountTransfer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReceiptAccountTransfer) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptAccountTransfer) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

// 查询一个地址列表在某个执行器中余额
type ReqBalance struct {
	// 地址列表
	Addresses []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	// 执行器名称
	Execer string `protobuf:"bytes,2,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqBalance) Reset()                    { *m = ReqBalance{} }
func (m *ReqBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqBalance) ProtoMessage()               {}
func (*ReqBalance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReqBalance) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ReqBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

// Account 的列表
type Accounts struct {
	Acc []*Account `protobuf:"bytes,1,rep,name=acc" json:"acc,omitempty"`
}

func (m *Accounts) Reset()                    { *m = Accounts{} }
func (m *Accounts) String() string            { return proto.CompactTextString(m) }
func (*Accounts) ProtoMessage()               {}
func (*Accounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Accounts) GetAcc() []*Account {
	if m != nil {
		return m.Acc
	}
	return nil
}

type ReqTokenBalance struct {
	Addresses   []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	TokenSymbol string   `protobuf:"bytes,2,opt,name=tokenSymbol" json:"tokenSymbol,omitempty"`
	Execer      string   `protobuf:"bytes,3,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqTokenBalance) Reset()                    { *m = ReqTokenBalance{} }
func (m *ReqTokenBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenBalance) ProtoMessage()               {}
func (*ReqTokenBalance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReqTokenBalance) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ReqTokenBalance) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

func (m *ReqTokenBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

// 手续费
type TotalFee struct {
	Fee     int64 `protobuf:"varint,1,opt,name=fee" json:"fee,omitempty"`
	TxCount int64 `protobuf:"varint,2,opt,name=txCount" json:"txCount,omitempty"`
}

func (m *TotalFee) Reset()                    { *m = TotalFee{} }
func (m *TotalFee) String() string            { return proto.CompactTextString(m) }
func (*TotalFee) ProtoMessage()               {}
func (*TotalFee) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TotalFee) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *TotalFee) GetTxCount() int64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

// 查询symbol代币总额
type ReqGetTotalCoins struct {
	Symbol    string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	StateHash []byte `protobuf:"bytes,2,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	StartKey  []byte `protobuf:"bytes,3,opt,name=startKey,proto3" json:"startKey,omitempty"`
	Count     int64  `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
}

func (m *ReqGetTotalCoins) Reset()                    { *m = ReqGetTotalCoins{} }
func (m *ReqGetTotalCoins) String() string            { return proto.CompactTextString(m) }
func (*ReqGetTotalCoins) ProtoMessage()               {}
func (*ReqGetTotalCoins) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReqGetTotalCoins) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqGetTotalCoins) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *ReqGetTotalCoins) GetStartKey() []byte {
	if m != nil {
		return m.StartKey
	}
	return nil
}

func (m *ReqGetTotalCoins) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// 查询symbol代币总额应答
type ReplyGetTotalCoins struct {
	Count   int64  `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Num     int64  `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
	Amount  int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	NextKey []byte `protobuf:"bytes,4,opt,name=nextKey,proto3" json:"nextKey,omitempty"`
}

func (m *ReplyGetTotalCoins) Reset()                    { *m = ReplyGetTotalCoins{} }
func (m *ReplyGetTotalCoins) String() string            { return proto.CompactTextString(m) }
func (*ReplyGetTotalCoins) ProtoMessage()               {}
func (*ReplyGetTotalCoins) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReplyGetTotalCoins) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReplyGetTotalCoins) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *ReplyGetTotalCoins) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReplyGetTotalCoins) GetNextKey() []byte {
	if m != nil {
		return m.NextKey
	}
	return nil
}

// 迭代查询symbol代币总额
type IterateRangeByStateHash struct {
	StateHash []byte `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Start     []byte `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End       []byte `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	Count     int64  `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
}

func (m *IterateRangeByStateHash) Reset()                    { *m = IterateRangeByStateHash{} }
func (m *IterateRangeByStateHash) String() string            { return proto.CompactTextString(m) }
func (*IterateRangeByStateHash) ProtoMessage()               {}
func (*IterateRangeByStateHash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *IterateRangeByStateHash) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *IterateRangeByStateHash) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *IterateRangeByStateHash) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *IterateRangeByStateHash) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReqAccountTokenAssets struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Execer  string `protobuf:"bytes,2,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqAccountTokenAssets) Reset()                    { *m = ReqAccountTokenAssets{} }
func (m *ReqAccountTokenAssets) String() string            { return proto.CompactTextString(m) }
func (*ReqAccountTokenAssets) ProtoMessage()               {}
func (*ReqAccountTokenAssets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ReqAccountTokenAssets) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ReqAccountTokenAssets) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

type TokenAsset struct {
	Symbol  string   `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Account *Account `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
}

func (m *TokenAsset) Reset()                    { *m = TokenAsset{} }
func (m *TokenAsset) String() string            { return proto.CompactTextString(m) }
func (*TokenAsset) ProtoMessage()               {}
func (*TokenAsset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *TokenAsset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *TokenAsset) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type ReplyAccountTokenAssets struct {
	TokenAssets []*TokenAsset `protobuf:"bytes,1,rep,name=tokenAssets" json:"tokenAssets,omitempty"`
}

func (m *ReplyAccountTokenAssets) Reset()                    { *m = ReplyAccountTokenAssets{} }
func (m *ReplyAccountTokenAssets) String() string            { return proto.CompactTextString(m) }
func (*ReplyAccountTokenAssets) ProtoMessage()               {}
func (*ReplyAccountTokenAssets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ReplyAccountTokenAssets) GetTokenAssets() []*TokenAsset {
	if m != nil {
		return m.TokenAssets
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "types.Account")
	proto.RegisterType((*ReceiptExecAccountTransfer)(nil), "types.ReceiptExecAccountTransfer")
	proto.RegisterType((*ReceiptAccountTransfer)(nil), "types.ReceiptAccountTransfer")
	proto.RegisterType((*ReqBalance)(nil), "types.ReqBalance")
	proto.RegisterType((*Accounts)(nil), "types.Accounts")
	proto.RegisterType((*ReqTokenBalance)(nil), "types.ReqTokenBalance")
	proto.RegisterType((*TotalFee)(nil), "types.TotalFee")
	proto.RegisterType((*ReqGetTotalCoins)(nil), "types.ReqGetTotalCoins")
	proto.RegisterType((*ReplyGetTotalCoins)(nil), "types.ReplyGetTotalCoins")
	proto.RegisterType((*IterateRangeByStateHash)(nil), "types.IterateRangeByStateHash")
	proto.RegisterType((*ReqAccountTokenAssets)(nil), "types.ReqAccountTokenAssets")
	proto.RegisterType((*TokenAsset)(nil), "types.TokenAsset")
	proto.RegisterType((*ReplyAccountTokenAssets)(nil), "types.ReplyAccountTokenAssets")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x4d, 0x6f, 0xdb, 0x3c,
	0x0c, 0xc7, 0xe1, 0x3a, 0x69, 0x12, 0xb6, 0xcf, 0xb3, 0x4e, 0xe8, 0x5a, 0xa3, 0xd8, 0xc1, 0xf0,
	0xc9, 0x87, 0x21, 0x87, 0x15, 0xd8, 0xbd, 0x29, 0xf6, 0x52, 0x0c, 0xe8, 0x41, 0xcd, 0x17, 0x50,
	0x1c, 0x66, 0x0b, 0x9a, 0xc8, 0x89, 0xa4, 0x0c, 0x71, 0x3f, 0xc0, 0x3e, 0xf7, 0x40, 0x9a, 0xce,
	0x4b, 0xbb, 0x0c, 0xbd, 0xe9, 0x2f, 0x51, 0xe4, 0xef, 0x4f, 0xd1, 0x86, 0xff, 0x4c, 0x51, 0x94,
	0x2b, 0x1b, 0xfa, 0x0b, 0x57, 0x86, 0x52, 0xb5, 0x43, 0xb5, 0x40, 0x9f, 0x3d, 0x42, 0xe7, 0xa6,
	0xde, 0x57, 0x57, 0xd0, 0x2d, 0x56, 0xce, 0xa1, 0x2d, 0xaa, 0x24, 0x4a, 0xa3, 0xbc, 0xad, 0x37,
	0x5a, 0x25, 0xd0, 0x19, 0x99, 0x99, 0xb1, 0x05, 0x26, 0x47, 0x69, 0x94, 0xc7, 0xba, 0x91, 0xea,
	0x02, 0x8e, 0x27, 0xae, 0x7c, 0x42, 0x9b, 0xc4, 0x7c, 0x20, 0x4a, 0x29, 0x68, 0x99, 0xf1, 0xd8,
	0x25, 0xad, 0x34, 0xca, 0x7b, 0x9a, 0xd7, 0xd9, 0xef, 0x08, 0xae, 0x34, 0x16, 0x38, 0x5d, 0x84,
	0xcf, 0x6b, 0x2c, 0xa4, 0xf0, 0xd0, 0x19, 0xeb, 0x27, 0xe8, 0x08, 0x00, 0x69, 0x9b, 0xae, 0x45,
	0x7c, 0x6d, 0xa3, 0x55, 0x06, 0xad, 0x85, 0xc3, 0x5f, 0x5c, 0xfd, 0xe4, 0xe3, 0xff, 0x7d, 0xa6,
	0xef, 0x4b, 0x06, 0xcd, 0x67, 0x2a, 0x87, 0x4e, 0x0d, 0x1c, 0x98, 0xe5, 0x65, 0x58, 0x73, 0x9c,
	0x4d, 0xe0, 0x42, 0x38, 0x9e, 0x33, 0x34, 0x75, 0xa2, 0xd7, 0xd5, 0x39, 0xfa, 0x77, 0x9d, 0x01,
	0x80, 0xc6, 0xe5, 0x40, 0x5a, 0xf5, 0x1e, 0x7a, 0xd4, 0x06, 0xf4, 0x1e, 0x7d, 0x12, 0xa5, 0x71,
	0xde, 0xd3, 0xdb, 0x0d, 0x6a, 0x24, 0xb9, 0x45, 0xc7, 0x49, 0x7b, 0x5a, 0x54, 0xf6, 0x01, 0xba,
	0x92, 0xd7, 0xab, 0x14, 0x62, 0x53, 0x14, 0x7c, 0xf7, 0x65, 0x55, 0x3a, 0xca, 0xa6, 0xf0, 0x46,
	0xe3, 0x72, 0x58, 0x3e, 0xa2, 0x7d, 0x5d, 0xd9, 0x14, 0x4e, 0x02, 0x45, 0x3f, 0x54, 0xf3, 0x51,
	0x39, 0x93, 0xda, 0xbb, 0x5b, 0x3b, 0x60, 0xf1, 0x1e, 0xd8, 0x27, 0xe8, 0x0e, 0xcb, 0x60, 0x66,
	0x5f, 0x10, 0xd5, 0x19, 0xc4, 0x13, 0x44, 0xee, 0x5a, 0xac, 0x69, 0x49, 0x13, 0x13, 0xd6, 0xb7,
	0x04, 0xd6, 0x4c, 0x8c, 0xc8, 0xec, 0x09, 0xce, 0x34, 0x2e, 0xbf, 0x62, 0xe0, 0xdb, 0xb7, 0xe5,
	0xd4, 0xb2, 0x79, 0x5f, 0x03, 0xd4, 0x0f, 0x2f, 0x8a, 0xd8, 0x7d, 0x30, 0x01, 0xbf, 0x19, 0xff,
	0x93, 0xf3, 0x9c, 0xea, 0xed, 0x06, 0x0d, 0x8c, 0x0f, 0xc6, 0x85, 0xef, 0x58, 0x31, 0xdb, 0xa9,
	0xde, 0x68, 0x75, 0x0e, 0x6d, 0x6e, 0x0b, 0x0f, 0x60, 0xac, 0x6b, 0x91, 0x59, 0x50, 0x1a, 0x17,
	0xb3, 0x6a, 0xbf, 0xfa, 0x26, 0x36, 0xda, 0x89, 0x25, 0x4f, 0x76, 0x35, 0x17, 0x7a, 0x5a, 0x12,
	0xa5, 0x99, 0x73, 0xa0, 0xcc, 0x7a, 0xad, 0xc8, 0xab, 0xc5, 0x35, 0x63, 0xb4, 0x18, 0xa3, 0x91,
	0xd9, 0x0a, 0x2e, 0xef, 0x02, 0x3a, 0x13, 0x50, 0x1b, 0xfb, 0x03, 0x07, 0xd5, 0xc3, 0x06, 0x7e,
	0xcf, 0x5a, 0xf4, 0xdc, 0xda, 0x39, 0xb4, 0xd9, 0x8a, 0x98, 0xae, 0x05, 0x21, 0xa1, 0x1d, 0x8b,
	0x57, 0x5a, 0x1e, 0xb0, 0x79, 0x07, 0xef, 0x34, 0x2e, 0x9b, 0xd9, 0xa6, 0xb7, 0xbc, 0xf1, 0x1e,
	0x83, 0x27, 0x52, 0x79, 0x7a, 0x69, 0x74, 0x23, 0x0f, 0x8e, 0xdf, 0x3d, 0xc0, 0x36, 0xc1, 0xc1,
	0x77, 0xca, 0xa1, 0x23, 0xbf, 0x97, 0x43, 0x9f, 0x84, 0x1c, 0x67, 0xf7, 0x70, 0xc9, 0x2f, 0xf0,
	0x17, 0xb8, 0x6b, 0x19, 0xc5, 0x5a, 0xca, 0x94, 0xbf, 0x95, 0x44, 0xdb, 0x40, 0xbd, 0x1b, 0x35,
	0x3a, 0xe6, 0xdf, 0xd9, 0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0xb2, 0x83, 0xbc, 0xdf,
	0x04, 0x00, 0x00,
}
