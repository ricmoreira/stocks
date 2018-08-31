package models

import (
	"stocks/models/saft/go_SaftT104"
)

type Invoice struct {
	ID                                                                  string `json:"id,omitempty"`
	go_SaftT104.TxsdSourceDocumentsSequenceSalesInvoicesSequenceInvoice `bson:"inline"`
}