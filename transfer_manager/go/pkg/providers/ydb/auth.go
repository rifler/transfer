package ydb

import (
	"context"

	"github.com/doublecloud/tross/kikimr/public/sdk/go/ydb"
	"github.com/doublecloud/tross/library/go/core/log"
	"github.com/doublecloud/tross/library/go/core/xerrors"
	"github.com/doublecloud/tross/transfer_manager/go/pkg/credentials"
)

var JWTCredentials = func(content string, tokenServiceURL string) (ydb.Credentials, error) {
	return nil, xerrors.Errorf("not implemented")
}

// Credentials is an abstraction of API authorization credentials.
// See https://cloud.yandex.ru/docs/iam/concepts/authorization/authorization for details.
// Note that functions that return Credentials may return different Credentials implementation
// in next SDK version, and this is not considered breaking change.
type Credentials interface {
	// YandexCloudAPICredentials is a marker method. All compatible Credentials implementations have it
	YandexCloudAPICredentials()
}

var NewYDBCredsFromYCCreds = func(ycCreds Credentials, tokenService string) ydb.Credentials {
	return nil
}

type JWTAuthParams struct {
	KeyContent      string
	TokenServiceURL string
}

func ResolveCredentials(
	userDataAuth bool,
	oauthToken string,
	jwt JWTAuthParams,
	serviceAccountID string,
	logger log.Logger,
) (ydb.Credentials, error) {
	if serviceAccountID != "" {
		cc, err := credentials.NewServiceAccountCreds(logger, serviceAccountID)
		if err != nil {
			logger.Error("err", log.Error(err))
			return nil, xerrors.Errorf("cannot init kinesis reader config without credentials client: %w", err)
		}
		logger.Infof("try SA account: %v", serviceAccountID)
		if _, err := cc.Token(context.Background()); err != nil {
			logger.Error("failed resolve token from SA", log.Error(err))
			return nil, xerrors.Errorf("cannot resolve token from %T: %w", cc, err)
		}
		logger.Infof("bind SA account: %v", serviceAccountID)
		return cc, nil
	}
	if oauthToken != "" {
		cc := credentials.NewStaticCreds(oauthToken)
		return cc, nil
	}
	if len(jwt.KeyContent) > 0 {
		cc, err := JWTCredentials(jwt.KeyContent, jwt.TokenServiceURL)
		if err != nil {
			return nil, xerrors.Errorf("cannot create jwt token: %w", err)
		}
		return cc, nil
	}
	if userDataAuth {
		return credentials.NewIamCreds(logger)
	}
	return nil, nil // no creds
}
