package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type Set struct {
	CRUD_Base
	Gedis *redis.Client
}

// Create a key
func (c *Set) Create(key string, value interface{}) error {
	//validate value
	if value == nil {
		log.Println("Value invalid!")
		return fmt.Errorf("Value invalid!")
	}
	// Set key
	res := c.Gedis.SAdd(context.Background(), key, value)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return res.Err()
	}
	return nil
}

func (c *Set) Update(key string, value interface{}) error {
	//validate value
	if value == nil {
		log.Println("Value invalid!")
		return fmt.Errorf("Value invalid!")
	}
	// remove old element
	res := c.Gedis.Del(context.Background(), key)
	if res.Err() != nil {
		log.Println("Cant del old data: ", res.Err().Error())
		return res.Err()
	}

	// Set key
	res = c.Gedis.SAdd(context.Background(), key, value)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return res.Err()
	}
	return nil
}

func (c *Set) Read(key string) (interface{}, error) {
	res := c.Gedis.SMembers(context.Background(), key)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return nil, res.Err()
	}
	return res.Val(), nil
}

func (c *Set) Delete(key string) error {
	res := c.Gedis.Del(context.Background(), key)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return res.Err()
	}
	return nil
}

func (c *Set) Push(key string, value interface{}) error {
	//validate value
	if value == nil {
		log.Println("Value invalid!")
		return fmt.Errorf("Value invalid!")
	}
	// Set key
	res := c.Gedis.SAdd(context.Background(), key, value)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return res.Err()
	}
	return nil
}
