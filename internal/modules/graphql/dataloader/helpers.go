package dataloader

import (
	"calend/internal/models/access"
	"github.com/graph-gophers/dataloader"
)

func KeysToUintSlice(keys dataloader.Keys) []uint {

	var s []uint
	for _, key := range keys {
		if val, ok := key.(UintKey); ok {
			s = append(s, uint(val))
		}
	}

	return s
}

func KeysToStringSlice(keys dataloader.Keys) []string {

	var s []string
	for _, key := range keys {
		if val, ok := key.(StringKey); ok {
			s = append(s, string(val))
		}
	}

	return s
}

func KeysToAccessRightSlice(keys dataloader.Keys) []access.Type {

	var s []access.Type
	for _, key := range keys {
		if val, ok := key.(AccessRightKey); ok {
			s = append(s, access.Type(val))
		}
	}

	return s
}
