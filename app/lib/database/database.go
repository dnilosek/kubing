package database

import "github.com/go-redis/redis"

type DB struct {
	Client *redis.Client
}

// Close connection
func (db *DB) Close() error {
	return db.Client.Close()
}

// Set key/value pair
func (db *DB) Set(key, value string) error {
	return db.Client.Set(key, value, 0).Err()
}

// Get value from key
func (db *DB) Get(key string) (string, error) {
	return db.Client.Get(key).Result()
}

// Open connection to redis db
func Open(databaseConnection string) (*DB, error) {
	var db DB
	options, err := redis.ParseURL(databaseConnection)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(options)
	db.Client = client
	return &db, nil
}
