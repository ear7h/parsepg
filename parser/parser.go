//go:generate pigeon -o pg_gen.go pg.peg

package parser

type Query string
type Queries []Query
