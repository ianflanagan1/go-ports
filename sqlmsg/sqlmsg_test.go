package sqlmsg_test

import (
	"testing"
	"time"

	"github.com/ianflanagan1/go-ports/sqlmsg"
)

// describe reads a Value the way a caller does, with an exhaustive switch.
func describe(v sqlmsg.Value) string {
	switch v := v.(type) {
	case sqlmsg.Null:
		return "null"
	case sqlmsg.Int:
		return "int"
	case sqlmsg.Float:
		return "float"
	case sqlmsg.Text:
		return "text " + string(v)
	case sqlmsg.Bytes:
		return "bytes " + string(v)
	case sqlmsg.Bool:
		return "bool"
	case sqlmsg.Time:
		return "time " + v.Format(time.DateOnly)
	}
	return "unreachable"
}

func TestEachKindOfValueIsItsOwnType(t *testing.T) {
	for _, tc := range []struct {
		value sqlmsg.Value
		want  string
	}{
		{sqlmsg.Null{}, "null"},
		{sqlmsg.Int(7), "int"},
		{sqlmsg.Float(1.5), "float"},
		{sqlmsg.Text("a"), "text a"},
		{sqlmsg.Bytes("ab"), "bytes ab"},
		{sqlmsg.Bool(true), "bool"},
		{sqlmsg.Time{Time: time.Date(2026, time.October, 28, 9, 0, 0, 0, time.UTC)}, "time 2026-10-28"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			if got := describe(tc.value); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestColumnFindsAColumnByName(t *testing.T) {
	rows := sqlmsg.Rows{Columns: []string{"id", "name"}}
	if i, ok := rows.Column("name"); !ok || i != 1 {
		t.Fatalf("got %d, %v; want 1, true", i, ok)
	}
	if _, ok := rows.Column("missing"); ok {
		t.Fatal("found a missing column")
	}
}
