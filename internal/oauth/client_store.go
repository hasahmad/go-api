package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"

	"github.com/jmoiron/sqlx"
)

// ClientStore PostgreSQL client store
type ClientStore struct {
	db        *sqlx.DB
	tableName string
	logger    Logger

	initTableDisabled bool
}

// ClientStoreItem data item
type ClientStoreItem struct {
	ID     string `db:"id"`
	Secret string `db:"secret"`
	Domain string `db:"domain"`
	Data   []byte `db:"data"`
}

// NewClientStore creates PostgreSQL store instance
func NewClientStore(db *sqlx.DB, options ...ClientStoreOption) (*ClientStore, error) {
	store := &ClientStore{
		db:        db,
		tableName: "oauth2_clients",
		logger:    log.New(os.Stderr, "[OAUTH2-PG-ERROR]", log.LstdFlags),
	}

	for _, o := range options {
		o(store)
	}

	var err error
	if !store.initTableDisabled {
		err = store.initTable()
	}

	if err != nil {
		return store, err
	}

	return store, err
}

func (s *ClientStore) initTable() error {
	_, err := s.db.ExecContext(context.Background(), fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
  id     TEXT  NOT NULL,
  secret TEXT  NOT NULL,
  domain TEXT  NOT NULL,
  data   JSONB NOT NULL,
  CONSTRAINT %[1]s_pkey PRIMARY KEY (id)
);
`, s.tableName))
	return err
}

// GetByID retrieves and returns client information by id
func (s *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	if id == "" {
		return nil, nil
	}

	var item ClientStoreItem
	if err := s.db.GetContext(ctx, &item, fmt.Sprintf("SELECT * FROM %s WHERE id = $1", s.tableName), id); err != nil {
		return nil, err
	}

	return s.toClientInfo(item.Data)
}

// Create creates and stores the new client information
func (s *ClientStore) Create(info oauth2.ClientInfo) error {
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(
		context.Background(),
		fmt.Sprintf("INSERT INTO %s (id, secret, domain, data) VALUES ($1, $2, $3, $4)", s.tableName),
		info.GetID(),
		info.GetSecret(),
		info.GetDomain(),
		data,
	)

	return err
}

func (s *ClientStore) toClientInfo(data []byte) (oauth2.ClientInfo, error) {
	var cm models.Client
	err := json.Unmarshal(data, &cm)
	return &cm, err
}
