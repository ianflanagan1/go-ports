package sqlmsg_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/ianflanagan1/go-ports/sqlmsg"
)

func TestEachConstructorSetsItsKindAndField(t *testing.T) {
	now := time.Date(2026, time.October, 28, 9, 0, 0, 0, time.UTC)
	for _, tc := range []struct {
		name string
		got  sqlmsg.Value
		want sqlmsg.Value
	}{
		{"null", sqlmsg.NullValue(), sqlmsg.Value{Kind: sqlmsg.Null}},
		{"int", sqlmsg.IntValue(7), sqlmsg.Value{Kind: sqlmsg.Int, Int: 7}},
		{"float", sqlmsg.FloatValue(1.5), sqlmsg.Value{Kind: sqlmsg.Float, Float: 1.5}},
		{"text", sqlmsg.TextValue("a"), sqlmsg.Value{Kind: sqlmsg.Text, Text: "a"}},
		{"bool", sqlmsg.BoolValue(true), sqlmsg.Value{Kind: sqlmsg.Bool, Bool: true}},
		{"time", sqlmsg.TimeValue(now), sqlmsg.Value{Kind: sqlmsg.Time, Time: now}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if !reflect.DeepEqual(tc.got, tc.want) {
				t.Fatalf("got %+v, want %+v", tc.got, tc.want)
			}
		})
	}
}

func TestBytesValueKeepsTheBytes(t *testing.T) {
	got := sqlmsg.BytesValue([]byte("ab"))
	if got.Kind != sqlmsg.Bytes || string(got.Bytes) != "ab" {
		t.Fatalf("got %+v", got)
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
