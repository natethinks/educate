package bolt

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/natethinks/educate"
)

// Client represents a client to the underlying BoltDB data store.
type Client struct {
	// Filename to the BoltDB database
	Path string

	// Returns the current time
	Now func() time.Time

	db *bolt.DB
}

// NewClient returns a new client for communicating with the db
func NewClient() *Client {
	c := &Client{Now: time.Now}
	return c
}

// Open opens and initializes the BoltDB database.
func (c *Client) Open() error {
	// Open database file.
	db, err := bolt.Open(c.Path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	c.db = db

	// Initialize top-level buckets.
	tx, err := c.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.CreateBucketIfNotExists([]byte("Dials")); err != nil {
		return err
	}

	return tx.Commit()
}

// UserService functions

// User retrieves data on a single user
func (c *Client) User(id int) (*educate.User, error) {
	fmt.Println("User()")

	u := educate.User{ID: 12, Username: "asdf"}

	return &u, nil
}

// Users is a Bolt implementation of Users
func (c *Client) Users() ([]*educate.User, error) {
	fmt.Println("Users()")

	var users []*educate.User

	u := educate.User{ID: 12, Username: "asdf"}

	users = append(users, &u)

	return users, nil
}

// UserGroup returns a slice of users that belong to a specific group
func (c *Client) UserGroup(id int) ([]*educate.User, error) {
	fmt.Println("UserGroup()")

	var users []*educate.User

	u := educate.User{ID: 12, Username: "asdf"}

	users = append(users, &u)

	return users, nil
}
