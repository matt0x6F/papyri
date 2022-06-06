package database

import (
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
	bolt "go.etcd.io/bbolt"
)

const (
	DefaultName = "papyri.db"
)

var (
	// DefaultConnection is the source connection that all other connections should create transactions off of.
	DefaultConnection *Connection
)

type Connection struct {
	db *bolt.DB
}

func (c *Connection) Initialize() error {
	debug := log.With().Str("pkg", "db").Logger()
	err := c.db.Update(
		func(tx *bolt.Tx) error {
			debug.Debug().Str("bucket", "main").Msg("Initializing bucket")
			main, err := tx.CreateBucketIfNotExists([]byte("main"))
			if err != nil {
				return err
			}

			debug.Debug().Str("bucket", "main.config").Msg("Initializing bucket")
			_, err = main.CreateBucketIfNotExists([]byte("config"))
			if err != nil {
				return err
			}

			debug.Debug().Str("bucket", "main.sources").Msg("Initializing bucket")
			_, err = main.CreateBucketIfNotExists([]byte("sources"))
			if err != nil {
				return err
			}

			debug.Debug().Str("bucket", "main.items").Msg("Initializing bucket")
			_, err = main.CreateBucketIfNotExists([]byte("items"))
			if err != nil {
				return err
			}

			return nil
		},
	)
	return err
}

func (c *Connection) Close() error {
	return c.db.Close()
}

func (c *Connection) Config() (Config, error) {
	var cfg Config

	err := c.db.View(
		func(tx *bolt.Tx) error {
			raw := tx.Bucket([]byte("main")).Get([]byte("config"))
			if raw != nil {
				err := json.Unmarshal(raw, &cfg)
				if err != nil {
					return err
				}
			}

			log.Logger.Debug().Bytes("content", raw).Msg("Fetched config")

			return nil
		},
	)
	return cfg, err
}

type Config struct {
	UpdateInterval time.Duration
}

func New(path string) (*Connection, error) {
	log.Debug().Str("path", path).Msg("Opening database")

	db, err := bolt.Open(path, 0o666, nil)
	if err != nil {
		return nil, err
	}

	return &Connection{
		db,
	}, nil
}
