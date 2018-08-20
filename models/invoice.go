package models

import (
	"stocks/models/saft/go_SaftT104"
)

type Invoice struct {
	go_SaftT104.TxsdSourceDocumentsSequenceSalesInvoicesSequenceInvoice `bson:"inline"`
}