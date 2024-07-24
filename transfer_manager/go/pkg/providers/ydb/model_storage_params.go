package ydb

import server "github.com/doublecloud/tross/transfer_manager/go/pkg/abstract/model"

type YdbStorageParams struct {
	Database           string
	Instance           string
	Tables             []string
	TableColumnsFilter []YdbColumnsFilter
	UseFullPaths       bool

	// auth props
	Token            server.SecretString
	ServiceAccountID string
	UserdataAuth     bool
	SAKeyContent     string
	TokenServiceURL  string

	RootCAFiles []string
	TLSEnabled  bool
}
