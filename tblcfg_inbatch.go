package repoutils

import (
	D "github.com/fbaube/dsmnd"
	FU "github.com/fbaube/fileutils"
)

// Inbatch describes a single import batch at the CLI.
// NOTE: Maybe rename this to FileSet ?
type Inbatch struct {
	Idx_Inbatch int
	FilCt       int
	RelFP       string
	AbsFP       FU.AbsFilePath
	T_Cre       string
	Descr       string
}

// TableSummary_Inbatch describes the table.
var TableSummary_Inbatch = D.TableSummary{
	D.TABL, "INB", "inbatch", "Batch import of files"}

// ColumnSpecs_Inbatch specifies two path fields
// (rel & abs), three time fields (creation, import,
// last-edit), a description, and the file count.
var ColumnSpecs_Inbatch = []D.ColumnSpec{
	D.DD_RelFP,
	D.DD_AbsFP,
	D.DD_T_Cre,
	D.ColumnSpec{D.TEXT, "descr", "Batch descr.", "Inbatch description"},
	D.ColumnSpec{D.INTG, "filct", "Nr. of files", "Number of files"},
}

// TableConfig_Inbatch specifies the table name
// "inbatch" and no foreign keys.
var TableConfig_Inbatch = TableConfig{
	"inbatch",
	// no foreign keys
	nil,
	ColumnSpecs_Inbatch,
}
