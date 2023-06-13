package repoutils

import D "github.com/fbaube/dsmnd"

// TableConfig describes the field structure of a database table.
// We assume that it is enough to use just two column types, TEXT
// (for strings) and INTEGER (for integers). Also, for each table,
// a primary key is assumed and foreign keys are allowed.
//
// Note that the field Columns is a slice of [DbColSpec], each of
// which is four text fields: [TxtIntKeyEtc], Code, Name, Descr.
//
// The fields [ForenKeys] and [Columns] may be nil or length [0].
//
// Date-time's are not an issue for SQLite, since either a string
// or an int can be used. We favor using strings ("TEXT"), which
// are expected to be ISO-8601 / RFC 3339. It is the first option
// listed here:
//
// https://www.sqlite.org/datatype3.html#date_and_time_datatype:
//   - TEXT: "YYYY-MM-DD HH:MM:SS.SSS" (or with "T" in the blank position)
//   - REAL as Julian day numbers: the day count since 24 November 4714 BC
//   - INTEGER as Unix time: the seconds count since 1970-01-01 00:00:00 UTC
type TableConfig struct {
	TableName string
	ForenKeys []string
	Columns   []D.ColumnSpec
}
