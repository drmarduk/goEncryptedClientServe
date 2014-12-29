// channel
package main

import (
	"errors"
	"net"
	"time"

	"golang.org/x/crypto/otr"
)

type Channel struct {
	Name  string
	Topic string
	Users []User
}

type User struct {
	Username   string
	PublicKey  otr.PublicKey
	Connection net.Conn
	JoinTime   time.Time
}

func (c *Channel) AddUser(name string, key otr.PublicKey, cn net.Conn) error {
	u := User{
		Username:   name,
		PublicKey:  key,
		Connection: cn,
		JoinTime:   time.Now(),
	}

	c.Users = append(c.Users, u)
	return nil
}

func (c *Channel) RemoveUser(name string) error {
	for i, u := range c.Users {

		if u.Username == name {
			c.Users = append(c.Users[:i], c.Users[i+1:]...)
			return nil
		}
	}

	return errors.New("User not found in Userlist.")
}

func (c *Channel) UserList() []User {
	return c.Users
}
