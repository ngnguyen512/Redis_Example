package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type TypeString struct {
	CRUD_Base
	Gedis *redis.Client
}

// Create a key
func (c *TypeString) Create(key string, value interface{}) error {
	//validate value
	if value == nil {
		log.Println("Value invalid!")
		return fmt.Errorf("Value invalid!")
	}
	// Set key
	res := c.Gedis.Set(context.Background(), key, value, time.Second*0)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return res.Err()
	}
	return nil
}

func (c *TypeString) Update(key string, value interface{}) error {
	//validate value
	if value == nil {
		log.Println("Value invalid!")
		return fmt.Errorf("Value invalid!")
	}
	// Set key
	res := c.Gedis.Set(context.Background(), key, value, time.Second*0)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return res.Err()
	}
	return nil
}

func (c *TypeString) Read(key string) (interface{}, error) {
	res := c.Gedis.Get(context.Background(), key)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return nil, res.Err()
	}
	return res.Val(), nil
}

func (c *TypeString) Delete(key string) error {
	res := c.Gedis.Del(context.Background(), key)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return res.Err()
	}
	return nil
}

func (c *TypeString) IsExists(key string) (int64, error) {
	res := c.Gedis.Exists(context.Background(), key)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return 0, res.Err()
	}
	return res.Val(), nil
}
