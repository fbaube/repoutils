package repoutils

import (
	"fmt"
	D "github.com/fbaube/dsmnd"
	FU "github.com/fbaube/fileutils"
	// "github.com/fbaube/nurepo/db"
)

// TableSummary_ContentityRow summarizes the table.
var TableSummary_ContentityRow = D.TableSummary{
	D.TABL, "contentity", "cnt", "Content entity"}

// TableDescriptor_ContentityRow specifies 11 DB columns,
// incl primary key (assumed) and one foreign key, "inbatch".
var TableDescriptor_ContentityRow = TableDescriptor{
	"contentity",     // Name
	"cnt",            // ShortName
	"idx_contentity", // IDName
	"RelFP, AbsFP, T_Cre, T_Imp, T_Edt, Descr, " +
		"RawMT, Mimtp, MType", // ColumnNames
	// One foreign key: "inbatch"
	ColumnSpecs_ContentityRow, // []D.ColumnSpecs
}

// PtrFields implements interface [Row].
// TODO: Still can't sort out the notation for ptr constraints ?!
func (cro *ContentityRow) PtrFields() []any { // barfs on []db.PtrFields
	return []any{
		&cro.Idx_Contentity, &cro.Idx_Inbatch,
		&cro.PathProps.RelFP, &cro.PathProps.AbsFP,
		&cro.T_Cre, &cro.T_Imp, &cro.T_Edt, &cro.Descr,
		&cro.PathProps.TypedRaw.MarkupType,
		&cro.PathAnalysis.ContypingInfo.MimeType,
		&cro.PathAnalysis.ContypingInfo.MType,
		&cro.PathProps.TypedRaw.Raw}
}

// ColumnSpecs_ContentityRow specifies
//   - a primary key (actually, it doescNOT - a primary
//     key is assumed, and handled elsewhere)
//   - a foreign key "inbatch"
//   - two path fields (rel & abs)
//   - three time fields (creation, import, last-edit)
//   - a description
//   - three content-type fields (raw markup type, MIME-type, MType)
//     NOTE: these are persisted in the DB because
//   - - they are useful in searching thru content
//   - - they can be expensive to calculate at import time
//   - - they can be overridden by choices made by users
//   - the content itself
//   - (not for now!) XML content type and XML DOCTYPE
//   - (not for now!) two LwDITA fields (flavor
//     [xdita,hdita!,mdita]), LwDITA content type)
//
// .
var ColumnSpecs_ContentityRow = []D.ColumnSpec{
	D.ColumnSpec{D.FKEY, "idx_inbatch", "inbatch",
		"Input batch of imported content"},
	D.DD_RelFP,
	D.DD_AbsFP,
	D.DD_T_Cre,
	D.DD_T_Imp,
	D.DD_T_Edt,
	D.ColumnSpec{D.TEXT, "descr", "Description",
		"Content entity description"},
	D.ColumnSpec{D.TEXT, "rawmt", "Markup type", "Raw markup type"},
	D.ColumnSpec{D.TEXT, "mimtp", "MIME type", "MIME type"},
	D.ColumnSpec{D.TEXT, "mtype", "MType", "MType"},
	D.ColumnSpec{D.TEXT, "contt", "Content", "Entity raw content"},
	// D.ColSpec{D.TEXT, "xmlcontype", "XML contype", "XML content type"},
	// D.ColSpec{D.TEXT, "xmldoctype", "XML Doctype", "XML Doctype"},
	// D.ColSpec{D.TEXT, "ditaflavor", "LwDITA flavor", "LwDITA flavor"},
	// D.ColSpec{D.TEXT, "ditacontype", "LwDITA contype", "LwDITA cnt type"},
}

// ContentityRow describes (in the DB) the entity's content
// plus its "dead properties" - basically, properties that
// are set by the user, rather than calculated as needed.
// Has the entity Raw content, in [PathProps.TypedRaw.Raw].
type ContentityRow struct {
	Idx_Contentity int
	Idx_Inbatch    int // NOTE: Rename to FILESET? Could be multiple?
	// Times is T_Cre, T_Imp, T_Edt string
	Times
	// PathProps has Raw and is // => EntityProps !!
	// CT.TypedRaw { Raw, SU.MarkupType string };
	// RelFP, ShortFP string;
	// FileMeta { os.FileInfo, exists bool, MU.Errer }
	FU.PathProps
	// PathAnalysis is a ptr, so that we get a
	// NPE if it is not initialized properly;
	// or if analysis failed, if (for example)
	// the content is too short.
	// FU.PathAnalysis is
	// XU.ContypingInfo { FileExt, MimeType, =>
	//   MimeTypeAsSnift, MType string }
	// ContentityBasics { XmlRoot, Text, Meta CT.Span; // => TopLevel !!
	//     MetaFormat string; MetaProps SU.PropSet }
	// XmlContype string
	// *XU.ParsedPreamble
	// *XU.ParsedDoctype
	// DitaFlavor  string
	// DitaContype string
	*FU.PathAnalysis // NEED DETAIL

	Descr string
	// Contt string

	// For these next two fields, instead put the refs & defs
	//   into another table that FKEY's into this table.
	// ExtlLinkRefs // links that point outside this File
	// ExtlLinkDefs // link targets that are visible outside this File
	// Linker = an outgoing link
	// Linkee = the target of an outgoing link
	// Linkable = a symbol that CAN be a Linkee
}

/* REF
OLD: type TableConfig struct {
	TableName string
	ForenKeys []string
	Columns   []D.ColumnSpec
}
NEW: type TableDescriptor struct {
	Name string
	ShortName string
	IDName string
	ColumnNames string
	ColumnSpecs []D.ColumnSpec
} */

// TODO Write col desc's using Desmond !
// TODO Generate ColNames from ColumnSpecs_ContentityRow

// FIXME: String implements Stringer.
func (p *ContentityRow) String() string {
	return fmt.Sprintf("PP<%s> AR <%s>", "", "")
}
