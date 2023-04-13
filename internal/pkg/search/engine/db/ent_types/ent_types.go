package ent_types

import "entgo.io/ent/dialect/sql"

type Predicate func(*sql.Selector)
