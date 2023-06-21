package repoutils

import D "github.com/fbaube/dsmnd"

// TableDescriptor describes the field structure of a database table.
// or each table, a primary key is assumed and foreign keys are allowed.
//
// All text fields should be in lower-case. Enforcement of this will
// inevitably be patchy.
//
// Notes on particular fields:
//   - The field [ColumnSpecs] is a slice of [db.ColumnSpec],
//     which is [dsmnd.Datum] and has four text fields:
//     [Fundatype], StorNam, DispName, Descrription.
//   - The field [ForenKeys] may be nil or length [0].
//   - TODO: (Maybe): the field [ColumnSpecs] could be nil or len 0.
//     If so then it should "probably" be autogenerated (partially
//     only) by reflection from the contents of a same-named table
//     currently existing in the DB.
//
// Notes on date-time fields:
//   - These are not an issue for SQLite, since either a string or an
//     int can be used. However, date-time fields referenced using THIS
//     system (i.e. [db.TableDescriptor] and [db.DbColumnSpec]) use
//     strings (SQLite DDL "TEXT"), which are expected to be ISO-8601
//     / RFC 3339 (and probly UTC). It is the first option listed in
//     https://www.sqlite.org/datatype3.html#date_and_time_datatype:
//   - TEXT: "YYYY-MM-DD HH:MM:SS.SSS"
//   - REAL as Julian day numbers: the day count since 24 November 4714 BC
//   - INTEGER as Unix time: the seconds count since 1970-01-01 00:00:00 UTC
//   - NOTE: For TEXT "YYYY-MM-DD HH:MM:SS.SSS", this might often end up
//     in ISO format, which has a "T" instead of the blank " " . So for
//     better readability, and to avoid line breaks, we have a utility
//     that replaces either a blank (" ") or an ISO "T" with a "_".
//
// .
type TableDescriptor struct {
	// Name is the name of the table in the DB,
	// e.g. inbatch, contentity, topicref.
	Name string
	// Shortname is a short version for use in the
	// names of other variables, e.g. inb, cnt, trf.
	ShortName string
	// IDName is the name of the index field, which (for now)
	// we use in the same format BOTH for a primary key AND
	// as a foreign key.
	IDName string
	// ColumnNames is all column names (except primary key), in a
	// specified order, comma-separated, for use in SQL statements.
	// We omit the primary key so that we can use this for SQL
	// INSERT staements too.
	ColumnNames string
	// We used to have ForenKeys defined by name only, but this is
	// insufficient information, because we need the field name AND
	// the table name. In principle we could derive one from the
	// other using our other DB-related data structures, and maybe
	// we used to, but it adds significant complexity.
	// ForenKeys   []string
	ColumnSpecs []D.ColumnSpec
}

/*
func (p *TableDescriptor) NameOfKey() string {
	return IDName
}
*/
