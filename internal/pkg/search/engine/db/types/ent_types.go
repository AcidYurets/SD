package types

import "entgo.io/ent/dialect/sql"

type Predicate func(*sql.Selector)

type Wrapper func(p Predicate) Predicate
