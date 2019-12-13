package main

type Command interface {
	Run(...*Schema) error
}
