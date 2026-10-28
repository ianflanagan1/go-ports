// Package sqlmsg holds SQL queries and their results as plain values. It
// does no I/O, so a project's core can build queries and read rows, and its
// shell can declare ports in these types without importing database/sql.
package sqlmsg

import "time"

// Kind is the type of a Value.
type Kind int

const (
	// Null is SQL NULL.
	Null Kind = iota
	// Int is a whole number.
	Int
	// Float is a floating-point number.
	Float
	// Text is a string, or a value in its database's text format, such as
	// a decimal number.
	Text
	// Bytes is binary data.
	Bytes
	// Bool is true or false.
	Bool
	// Time is a timestamp or date.
	Time
)

// Value is a query argument or a column value. Kind says which field holds
// it; the others are zero.
type Value struct {
	Kind  Kind
	Int   int64
	Float float64
	Text  string
	Bytes []byte
	Bool  bool
	Time  time.Time
}

// NullValue returns SQL NULL.
func NullValue() Value { return Value{Kind: Null} }

// IntValue returns v as a Value.
func IntValue(v int64) Value { return Value{Kind: Int, Int: v} }

// FloatValue returns v as a Value.
func FloatValue(v float64) Value { return Value{Kind: Float, Float: v} }

// TextValue returns v as a Value.
func TextValue(v string) Value { return Value{Kind: Text, Text: v} }

// BytesValue returns v as a Value.
func BytesValue(v []byte) Value { return Value{Kind: Bytes, Bytes: v} }

// BoolValue returns v as a Value.
func BoolValue(v bool) Value { return Value{Kind: Bool, Bool: v} }

// TimeValue returns v as a Value.
func TimeValue(v time.Time) Value { return Value{Kind: Time, Time: v} }

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
