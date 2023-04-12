package ent_types

import "entgo.io/ent/dialect/sql"

type Predicate func(*sql.Selector)
type SortOptions func(options *sql.OrderTermOptions)
