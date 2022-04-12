package oauth

import (
	"fmt"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/hasahmad/go-api/internal/config"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func SetupOAuthServer(db *sqlx.DB, cfg config.Config, logger *logrus.Logger) *server.Server {
	const tokenGCInterval = time.Minute * 5

	clientStore, err := NewClientStore(db)
	if err != nil {
		fmt.Errorf("unable to create a PostgreSQL client store:  %v", err)
		return nil
	}

	tokenStore, err := NewTokenStore(db, WithTokenStoreGCInterval(tokenGCInterval))
	if err != nil {
		err = fmt.Errorf("unable to create a PostgreSQL token store:  %v", err)
		return nil
	}
	defer tokenStore.Close()

	manager := manage.NewDefaultManager()
	manager.MapClientStorage(clientStore)
	manager.MapTokenStorage(tokenStore)

	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// Set the authorization `code` expiry to 60 seconds.
	manager.SetAuthorizeCodeExp(cfg.Server.AuthCodeExp)

	// Configure the Authorization Code Grant.
	manager.SetAuthorizeCodeTokenCfg(&manage.Config{
		AccessTokenExp:    cfg.Server.AccessTokenExp,
		RefreshTokenExp:   cfg.Server.AuthCodeRefreshTokenExp,
		IsGenerateRefresh: true,
	})

	// Configure the Resource Owner Password Credentials Grant.
	manager.SetPasswordTokenCfg(&manage.Config{
		AccessTokenExp:    cfg.Server.AccessTokenExp,
		RefreshTokenExp:   cfg.Server.PassCredsRefreshTokenExp,
		IsGenerateRefresh: true,
	})

	// Configure the Client Credentials Grant.
	// NOTE: A refresh token is not necessary for a Client Credentials Grant client (a Resource
	// Owner).
	manager.SetClientTokenCfg(&manage.Config{
		AccessTokenExp: cfg.Server.AccessTokenExp,
	})

	oauth2Server := server.NewDefaultServer(manager)

	oauth2Server.SetAllowGetAccessRequest(true)
	oauth2Server.SetAllowedGrantType(oauth2.AuthorizationCode, oauth2.ClientCredentials,
		oauth2.PasswordCredentials, oauth2.Refreshing)
	oauth2Server.SetClientInfoHandler(server.ClientFormHandler)
	oauth2Server.SetPasswordAuthorizationHandler(passwordCredsAuthHandler(db, cfg, logger))
	oauth2Server.SetUserAuthorizationHandler(userAuthorizeHandler(db, oauth2Server))

	oauth2Server.SetInternalErrorHandler(func(err error) (resp *errors.Response) {
		logger.WithFields(logrus.Fields{
			"internal_error": err,
		})
		return
	})
	oauth2Server.SetResponseErrorHandler(func(resp *errors.Response) {
		logger.WithFields(logrus.Fields{
			"response_error": resp.Error,
		})
	})

	return oauth2Server
}
