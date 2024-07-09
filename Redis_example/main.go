package main

type value interface{}

type CRUD_Base interface {
	Create(string, value) error
	Read(string) (value, error)
	Update(string, value) error
	Delete(string) error
}
