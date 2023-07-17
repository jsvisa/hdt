package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/datatypes"
)

// copy from "github.com/forta-network/forta-core-go/clients/webhook/client/models"
type AlertBlock struct {
	// chain Id
	// Example: 1337
	ChainID uint64 `json:"chainId,omitempty"`

	// hash
	// Example: 0xf9e777b739cf90a197c74c461933422dcf26fadf50e0ef9aa72af76727da87ca
	Hash string `json:"hash,omitempty"`

	// number
	// Example: 1235678901234
	Number uint64 `json:"number,omitempty"`

	// Timestamp (RFC3339)
	// Example: 2022-03-01T12:24:33Z
	Timestamp string `json:"timestamp,omitempty"`
}

type AlertBot struct {
	// id
	// Example: 0x17381ae942ee1fe141d0652e9dad7d001761552f906fb1684b2812603de31049
	ID string `json:"id,omitempty"`

	// Docker image reference (Disco)
	// Example: bafybeibrigevnhic4befnkqbaagzgxqtdyv2fdgcbqwxe7ees3hw6fymme@sha256:9ca1547e130a6264bb1b4ad6b10f17cabf404957f23d457a30046b9afdf29fc8
	Image string `json:"image,omitempty"`

	// Bot reference (IPFS hash)
	// Example: QmU6L9Zo5rweF6QZLhLfwAAFUFRMF3uFdSnMiJzENXr37R
	Reference string `json:"reference,omitempty"`
}

type AlertSourceEvent struct {
	// Deterministic alert hash
	// Example: 0xe9cfda18f167de5cdd63c101e38ec0d4cb0a1c2dea80921ecc4405c2b010855f
	AlertHash string `json:"alertHash,omitempty"`

	// bot Id
	// Example: 0x17381ae942ee1fe141d0652e9dad7d001761552f906fb1684b2812603de31049
	BotID string `json:"botId,omitempty"`
}

type AlertSource struct {
	// block
	Block *AlertBlock `json:"block,omitempty" gorm:"type:json"`

	// bot
	Bot *AlertBot `json:"bot,omitempty" gorm:"type:json"`

	// source event
	SourceEvent *AlertSourceEvent `json:"sourceEvent,omitempty" gorm:"type:json"`

	// transaction hash
	// Example: 0x7040dd33cbfd3e9d880da80cb5f3697a717fc329abd0251f3dcd51599ab67b0a
	TransactionHash string `json:"transactionHash,omitempty"`
}

type RPCAlert struct {
	// Addresses involved in the source of this alert
	// Example: ["0x98883145049dec03c00cb7708cbc938058802520","0x1fFa3471A45C22B1284fE5a251eD74F40580a1E3"]
	Addresses pq.StringArray `json:"addresses" gorm:"type:text[]"`

	// AddressBloomFilter contains **all** addresses in the alert
	// Example: "addressBloomFilter": {"k": 11,"m": 44,"bitset": "AAAAAAAAACwAAAAAAAAACwAAAAAAAAAsAAALo5gpbbc=", item_count: 1}
	// AddressBloomFilter interface{} `json:"addressBloomFilter,omitempty"`

	// alert Id
	// Example: OZ-GNOSIS-EVENTS
	AlertID string `json:"alertId,omitempty"`

	// Timestamp (RFC3339Nano)
	// Example: 2022-03-01T12:24:33.379756298Z
	CreatedAt string `json:"createdAt,omitempty"`

	// description
	// Example: Detected Transfer event
	Description string `json:"description,omitempty"`

	// finding type
	// Enum: [UNKNOWN_TYPE EXPLOIT SUSPICIOUS DEGRADED INFORMATION SCAM]
	FindingType string `json:"findingType,omitempty"`

	// Deterministic alert hash
	// Example: 0xe9cfda18f167de5cdd63c101e38ec0d4cb0a1c2dea80921ecc4405c2b010855f
	Hash string `json:"hash,omitempty"`

	// An associative array of extra links values
	// Example: {"blockUrl":"https://etherscan.io/block/18646150","explorerUrl":"https://explorer.forta.network/alert/0xd795c365931762afeccf4a440ecee2f7e89820c59136aa46310a8eec54ba96d8"}
	// Links interface{} `json:"links,omitempty"`

	// An associative array of string values
	// Example: {"contractAddress":"0x98883145049dec03c00cb7708cbc938058802520","operator":"0x1fFa3471A45C22B1284fE5a251eD74F40580a1E3"}
	Metadata datatypes.JSON `json:"metadata,omitempty" gorm:"type:json"`

	// name
	// Example: Transfer Event
	Name string `json:"name,omitempty"`

	// protocol
	// Example: ethereum
	Protocol string `json:"protocol,omitempty"`

	// Related alerts involved in the source of this alert
	// Example: ["0xe9cfda18f167de5cdd63c101e38ec0d4cb0a1c2dea80921ecc4405c2b010855f","0x533c100d5d7a56ee8448b6b08b5b1ce41ea9d1667086e1d2d4c1f03d09d191b9"]
	RelatedAlerts []string `json:"relatedAlerts" gorm:"type:json"`

	// severity
	// Enum: [UNKNOWN INFO LOW MEDIUM HIGH CRITICAL]
	Severity string `json:"severity,omitempty"`

	// source
	Source *AlertSource `json:"source,omitempty" gorm:"type:json"`
}

type RPCAlerts struct {
	Alerts []*RPCAlert `json:"alerts"`
}

type Alert struct {
	ID             uint      `gorm:"primarykey"`
	Chain          string    `json:"chain,omitempty"`
	BlockTimestamp time.Time `json:"blockTimestamp,omitempty"`
	BlockNum       uint64    `json:"blockNum,omitempty"`
	TxHash         string    `json:"txhash,omitempty"`

	// Addresses involved in the source of this alert
	// Example: ["0x98883145049dec03c00cb7708cbc938058802520","0x1fFa3471A45C22B1284fE5a251eD74F40580a1E3"]
	Addresses pq.StringArray `json:"addresses" gorm:"type:text[]"`

	// AddressBloomFilter contains **all** addresses in the alert
	// Example: "addressBloomFilter": {"k": 11,"m": 44,"bitset": "AAAAAAAAACwAAAAAAAAACwAAAAAAAAAsAAALo5gpbbc=", item_count: 1}
	// AddressBloomFilter interface{} `json:"addressBloomFilter,omitempty"`

	// alert Id
	// Example: OZ-GNOSIS-EVENTS
	AlertID string `json:"alertId,omitempty"`

	// Timestamp (RFC3339Nano)
	// Example: 2022-03-01T12:24:33.379756298Z
	CreatedAt string `json:"createdAt,omitempty"`

	// name
	// Example: Transfer Event
	Name string `json:"name,omitempty"`

	// description
	// Example: Detected Transfer event
	Description string `json:"description,omitempty"`

	// finding type
	// Enum: [UNKNOWN_TYPE EXPLOIT SUSPICIOUS DEGRADED INFORMATION SCAM]
	FindingType string `json:"findingType,omitempty"`

	// An associative array of extra links values
	// Example: {"blockUrl":"https://etherscan.io/block/18646150","explorerUrl":"https://explorer.forta.network/alert/0xd795c365931762afeccf4a440ecee2f7e89820c59136aa46310a8eec54ba96d8"}
	// Links interface{} `json:"links,omitempty"`

	// An associative array of string values
	// Example: {"contractAddress":"0x98883145049dec03c00cb7708cbc938058802520","operator":"0x1fFa3471A45C22B1284fE5a251eD74F40580a1E3"}
	Metadata datatypes.JSON `json:"metadata,omitempty" gorm:"type:json"`

	// severity
	// Enum: [UNKNOWN INFO LOW MEDIUM HIGH CRITICAL]
	Severity string `json:"severity,omitempty"`
}
