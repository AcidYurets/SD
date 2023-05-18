package dataloader

import (
	"calend/internal/models/access"
	"strconv"
)

type StringKey string

func (k StringKey) String() string { return string(k) }

func (k StringKey) Raw() interface{} { return k }

type UintKey uint

func (k UintKey) String() string { return strconv.Itoa(int(k)) }

func (k UintKey) Raw() interface{} { return k }

type AccessRightKey access.Type

func (k AccessRightKey) String() string { return string(k) }

func (k AccessRightKey) Raw() interface{} { return k }
