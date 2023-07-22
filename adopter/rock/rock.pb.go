package rock

type PutStateRequest struct {
	Keys  []string `json:"keys,omitempty"`
	Value string   `json:"value,omitempty"`
}

type GetStateRequest struct {
	Keys []string `json:"keys,omitempty"`
}

type GetHistoryRequest struct {
	Keys []string `json:"keys,omitempty"`
}

type GetHistoryResponse struct {
	KeyModifications []*KeyModification `json:"key_modifications,omitempty"`
}

type KeyModification struct {
	TxId      string `json:"tx_id,omitempty"`
	Value     string `json:"value,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	IsDelete  bool   `json:"is_delete,omitempty"`
}
