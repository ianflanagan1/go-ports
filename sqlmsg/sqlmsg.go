// Package sqlmsg holds SQL queries and their results as plain values. It
// does no I/O, so a project's core can build queries and read rows, and its
// shell can declare ports in these types without importing database/sql.
package sqlmsg

import "time"

// Value is a query argument or a column value: Null, Int, Float, Text,
// Bytes, Bool or Time. Read one with a type switch that has a case for each.
//
//sumtype:decl
type Value interface{ isValue() }

// Null is SQL NULL.
type Null struct{}

// Int is a whole number.
type Int int64

// Float is a floating-point number.
type Float float64

// Text is a string, or a value in its database's text format, such as a
// decimal number.
type Text string

// Bytes is binary data.
type Bytes []byte

// Bool is true or false.
type Bool bool

// Time is a timestamp or date.
type Time struct{ time.Time }

func (Null) isValue()  {}
func (Int) isValue()   {}
func (Float) isValue() {}
func (Text) isValue()  {}
func (Bytes) isValue() {}
func (Bool) isValue()  {}
func (Time) isValue()  {}

// Query is a SQL statement and its arguments, written with the database's
// placeholders: $1, $2 and so on for Postgres.
type Query struct {
	SQL  string
	Args []Value
}

// Rows is what a query returns: the column names, and each row's values in
// the same order.
type Rows struct {
	Columns []string
	Values  [][]Value
}

// Column returns the index of the column called name, and whether there is
// one.
func (r Rows) Column(name string) (int, bool) {
	for i, column := range r.Columns {
		if column == name {
			return i, true
		}
	}
	return 0, false
}

// Result is what a statement that returns no rows reports.
type Result struct {
	RowsAffected int64
}

// Tx is the port for running queries inside a transaction. A transaction
// adapter passes one to the function it runs; the transaction's context
// applies to every call.
type Tx interface {
	Query(q Query) (Rows, error)
	Exec(q Query) (Result, error)
}
