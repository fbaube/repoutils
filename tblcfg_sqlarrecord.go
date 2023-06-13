package repoutils

import (
	"fmt"
	D "github.com/fbaube/dsmnd"
	FU "github.com/fbaube/fileutils"
)

// ContentityRecord is basically the content plus its "dead properties" -
// properties that are set by the user, rather than dynamically determined.
type SqlarRecord struct {
	Idx_Contentity int
	Idx_Inbatch    int // NOTE: Maybe rename to FILESET. Could be multiple?
	Descr          string
	// FU.PathProps
	Times
	FU.PathProps
	// PathAnalysis is a ptr, so that we get a
	// NPE if it is not initialized properly.
	*FU.PathAnalysis
	// For these next two fields, instead put the refs & defs
	//   into another table that FKEY's into this table.
	// ExtlLinkRefs // links that point outside this File
	// ExtlLinkDefs // link targets that are visible outside this File
	// Linker = an outgoing link
	// Linkee = the target of an outgoing link
	// Linkable = a symbol that CAN be a Linkee
}

// String implements Stringer.
func (p *SqlarRecord) String() string {
	return fmt.Sprintf("PP<%s> AR <%s>",
		p.PathProps.String(), p.PathAnalysis.String())
}

// ColumnSpecs_Contentity specifies two path fields (rel & abs),
// three time fields (creation, import, last-edit), a description,
// four XML-related fields (MIME-type, MType, XML content type, and
// XML DOCTYPE), and two LwDITA fields (flavor [xdita,hdita,mdita]),
// LwDITA content type).
var ColumnSpecs_Sqlar = []D.ColumnSpec{
	D.DD_RelFP,
	D.DD_AbsFP,
	D.DD_T_Cre,
	D.DD_T_Imp,
	D.DD_T_Edt,
	D.ColumnSpec{D.TEXT, "descr", "Description", "Content item description"},
	// D.ColumnSpec{D.TXT, "metaraw", "Meta (raw)", "Metadata/header (raw)"},
	// D.ColumnSpec{D.TXT, "textraw", "Text (raw)", "Text/body (raw)"},
	D.ColumnSpec{D.TEXT, "mimetype", "MIME type", "MIME type"},
	D.ColumnSpec{D.TEXT, "mtype", "MType", "MType"},
	// D.ColumnSpec{D.TXT, "roottag", "Root tag", "XML root tag"},
	// D.ColumnSpec{D.TXT, "rootatts", "Root att's", "XML root tag attributes"},
	D.ColumnSpec{D.TEXT, "xmlcontype", "XML contype", "XML content type"},
	// // 2022.09 OUCH // D.ColumnSpec{D.D_TXT, "xmldoctype", "XML Doctype", "XML Doctype"},
	D.ColumnSpec{D.TEXT, "ditaflavor", "(Lw)DITA flavor", "(Lw)DITA flavor"},
	D.ColumnSpec{D.TEXT, "ditacontype", "(Lw)DITA contype", "(Lw)DITA content type"},
}

// TableConfig_Contentity specifies the table name
// "contentity" and one foreign key, "inbatch".
var TableConfig_Sqlar = TableConfig{
	"contentity",
	// One foreign key
	[]string{"inbatch"},
	ColumnSpecs_Contentity,
}
