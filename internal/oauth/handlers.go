package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/hasahmad/go-api/internal/config"
	"github.com/hasahmad/go-api/internal/helpers"
	"github.com/hasahmad/go-api/internal/repository"
	"github.com/jmoiron/sqlx"
)

// passAuthenticator validates a user supplied username-password pair.  On validation, the userID is set.
func passwordCredsAuthHandler(db *sqlx.DB, cfg config.Config) server.PasswordAuthorizationHandler {
	return func(ctx context.Context, username, password string) (userID string, err error) {
		repo := repository.NewUserRepo(db, cfg, nil)
		user, err := repo.FindByUsername(ctx, username)
		userID = ""
		if err != nil {
			err = fmt.Errorf("unable to find userByUsername (%s): %w", username, err)
			return
		}
		if user.UserID.String() == "" {
			err = fmt.Errorf("no record found (%s): %w", username, err)
			return
		}

		matched, err := user.Password.Matches(password)
		if err != nil {
			err = fmt.Errorf("could not authorize user: %w", err)
			return
		}
		if !matched {
			err = fmt.Errorf("invalid credentials")
			return
		}

		userID = user.UserID.String()

		return
	}
}

// userAuthorizeHandler validates a user supplied access token.
//
// On validation, the request is processed.
// On failure, the user is redirected to the login page.
func userAuthorizeHandler(db *sqlx.DB, oauth2Server *server.Server) server.UserAuthorizationHandler {
	return func(wr http.ResponseWriter, req *http.Request) (userID string, err error) {
		token, err := oauth2Server.ValidationBearerToken(req)
		if err != nil {
			helpers.WriteJSON(wr, 403, helpers.Envelope{
				"error":         "invalid_login",
				"error_details": "Invalid Login",
			}, nil)
			return
		}

		userID = token.GetUserID()

		return
	}
}

// testHandler handles the "/test" route requests.
func testHandlerFunc(db *sqlx.DB, oauth2Server *server.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(wr http.ResponseWriter, req *http.Request) {
		token, err := oauth2Server.ValidationBearerToken(req)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		if err := req.ParseForm(); err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}
		scope := []string{req.Form.Get("scope")}
		if scope == nil {
			http.Error(wr, "Undefined scope", http.StatusBadRequest)
		}

		/*     data := map[string]interface{}{
		*         "expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		*         "client_id":  token.GetClientID(),
		*         "user_id":    token.GetUserID(),
		*     }
		*
		*     encoder := json.NewEncoder(wr)
		*     encoder.SetIndent("", "  ")
		*     encoder.Encode(data) */

		output := map[string]interface{}{
			"token": token,
			"scope": scope,
		}

		encoder := json.NewEncoder(wr)
		encoder.SetIndent("", "  ")
		_ = encoder.Encode(output)
	}
}
