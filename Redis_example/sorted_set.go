package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type SortedSet struct {
	CRUD_Base
	Gedis *redis.Client
}

// Create a key
func (c *SortedSet) Create(key string, value interface{}) error {
	//validate value
	v, ok := value.(redis.Z)
	if !ok || value == nil {
		log.Println("Value invalid!")
		return fmt.Errorf("Value invalid!")
	}

	// Set key
	res := c.Gedis.ZAdd(context.Background(), key, v)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return res.Err()
	}
	return nil
}

func (c *SortedSet) Update(key string, value interface{}) error {
	//validate value
	v, ok := value.(redis.Z)
	if !ok || value == nil {
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
	res = c.Gedis.ZAdd(context.Background(), key, v)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return res.Err()
	}
	return nil
}

func (c *SortedSet) Read(key string) (interface{}, error) {
	res := c.Gedis.ZRange(context.Background(), key, 0, -1)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return nil, res.Err()
	}
	return res.Val(), nil
}

func (c *SortedSet) Delete(key string) error {
	res := c.Gedis.Del(context.Background(), key)
	if res.Err() != nil {
		log.Println("Error: ", res.Err().Error())
		return res.Err()
	}
	return nil
}
