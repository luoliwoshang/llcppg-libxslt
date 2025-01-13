package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	"unsafe"
)

const XSLTMAXSORT c.Int = 15

type X_XsltRuntimeExtra struct {
	Info       unsafe.Pointer
	Deallocate unsafe.Pointer
	Val        struct {
		Ptr unsafe.Pointer
	}
}
type XsltRuntimeExtra X_XsltRuntimeExtra
type XsltRuntimeExtraPtr *XsltRuntimeExtra

type X_XsltTemplate struct {
	Next           *X_XsltTemplate
	Style          *X_XsltStylesheet
	Match          *libxml_2_0.XmlChar
	Priority       float32
	Name           *libxml_2_0.XmlChar
	NameURI        *libxml_2_0.XmlChar
	Mode           *libxml_2_0.XmlChar
	ModeURI        *libxml_2_0.XmlChar
	Content        libxml_2_0.XmlNodePtr
	Elem           libxml_2_0.XmlNodePtr
	InheritedNsNr  c.Int
	InheritedNs    *libxml_2_0.XmlNsPtr
	NbCalls        c.Int
	Time           c.Ulong
	Params         unsafe.Pointer
	TemplNr        c.Int
	TemplMax       c.Int
	TemplCalledTab *XsltTemplatePtr
	TemplCountTab  *c.Int
	Position       c.Int
}
type XsltTemplate X_XsltTemplate
type XsltTemplatePtr *XsltTemplate

type X_XsltStylesheet struct {
	Parent             *X_XsltStylesheet
	Next               *X_XsltStylesheet
	Imports            *X_XsltStylesheet
	DocList            XsltDocumentPtr
	Doc                libxml_2_0.XmlDocPtr
	StripSpaces        libxml_2_0.XmlHashTablePtr
	StripAll           c.Int
	CdataSection       libxml_2_0.XmlHashTablePtr
	Variables          XsltStackElemPtr
	Templates          XsltTemplatePtr
	TemplatesHash      libxml_2_0.XmlHashTablePtr
	RootMatch          *X_XsltCompMatch
	KeyMatch           *X_XsltCompMatch
	ElemMatch          *X_XsltCompMatch
	AttrMatch          *X_XsltCompMatch
	ParentMatch        *X_XsltCompMatch
	TextMatch          *X_XsltCompMatch
	PiMatch            *X_XsltCompMatch
	CommentMatch       *X_XsltCompMatch
	NsAliases          libxml_2_0.XmlHashTablePtr
	AttributeSets      libxml_2_0.XmlHashTablePtr
	NsHash             libxml_2_0.XmlHashTablePtr
	NsDefs             unsafe.Pointer
	Keys               unsafe.Pointer
	Method             *libxml_2_0.XmlChar
	MethodURI          *libxml_2_0.XmlChar
	Version            *libxml_2_0.XmlChar
	Encoding           *libxml_2_0.XmlChar
	OmitXmlDeclaration c.Int
	DecimalFormat      XsltDecimalFormatPtr
	Standalone         c.Int
	DoctypePublic      *libxml_2_0.XmlChar
	DoctypeSystem      *libxml_2_0.XmlChar
	Indent             c.Int
	MediaType          *libxml_2_0.XmlChar
	PreComps           XsltElemPreCompPtr
	Warnings           c.Int
	Errors             c.Int
	ExclPrefix         *libxml_2_0.XmlChar
	ExclPrefixTab      **libxml_2_0.XmlChar
	ExclPrefixNr       c.Int
	ExclPrefixMax      c.Int
	X_Private          unsafe.Pointer
	ExtInfos           libxml_2_0.XmlHashTablePtr
	ExtrasNr           c.Int
	Includes           XsltDocumentPtr
	Dict               libxml_2_0.XmlDictPtr
	AttVTs             unsafe.Pointer
	DefaultAlias       *libxml_2_0.XmlChar
	Nopreproc          c.Int
	Internalized       c.Int
	LiteralResult      c.Int
	Principal          XsltStylesheetPtr
	ForwardsCompatible c.Int
	NamedTemplates     libxml_2_0.XmlHashTablePtr
	XpathCtxt          libxml_2_0.XmlXPathContextPtr
	OpLimit            c.Ulong
	OpCount            c.Ulong
}

type X_XsltDecimalFormat struct {
	Next             *X_XsltDecimalFormat
	Name             *libxml_2_0.XmlChar
	Digit            *libxml_2_0.XmlChar
	PatternSeparator *libxml_2_0.XmlChar
	MinusSign        *libxml_2_0.XmlChar
	Infinity         *libxml_2_0.XmlChar
	NoNumber         *libxml_2_0.XmlChar
	DecimalPoint     *libxml_2_0.XmlChar
	Grouping         *libxml_2_0.XmlChar
	Percent          *libxml_2_0.XmlChar
	Permille         *libxml_2_0.XmlChar
	ZeroDigit        *libxml_2_0.XmlChar
	NsUri            *libxml_2_0.XmlChar
}
type XsltDecimalFormat X_XsltDecimalFormat
type XsltDecimalFormatPtr *XsltDecimalFormat

type X_XsltDocument struct {
	Next           *X_XsltDocument
	Main           c.Int
	Doc            libxml_2_0.XmlDocPtr
	Keys           unsafe.Pointer
	Includes       *X_XsltDocument
	Preproc        c.Int
	NbKeysComputed c.Int
}
type XsltDocument X_XsltDocument
type XsltDocumentPtr *XsltDocument

type X_XsltKeyDef struct {
	Next    *X_XsltKeyDef
	Inst    libxml_2_0.XmlNodePtr
	Name    *libxml_2_0.XmlChar
	NameURI *libxml_2_0.XmlChar
	Match   *libxml_2_0.XmlChar
	Use     *libxml_2_0.XmlChar
	Comp    libxml_2_0.XmlXPathCompExprPtr
	Usecomp libxml_2_0.XmlXPathCompExprPtr
	NsList  *libxml_2_0.XmlNsPtr
	NsNr    c.Int
}
type XsltKeyDef X_XsltKeyDef
type XsltKeyDefPtr *XsltKeyDef

type X_XsltKeyTable struct {
	Next    *X_XsltKeyTable
	Name    *libxml_2_0.XmlChar
	NameURI *libxml_2_0.XmlChar
	Keys    libxml_2_0.XmlHashTablePtr
}
type XsltKeyTable X_XsltKeyTable
type XsltKeyTablePtr *XsltKeyTable
type XsltStylesheet X_XsltStylesheet
type XsltStylesheetPtr *XsltStylesheet

type X_XsltTransformContext struct {
	Style               XsltStylesheetPtr
	Type                XsltOutputType
	Templ               XsltTemplatePtr
	TemplNr             c.Int
	TemplMax            c.Int
	TemplTab            *XsltTemplatePtr
	Vars                XsltStackElemPtr
	VarsNr              c.Int
	VarsMax             c.Int
	VarsTab             *XsltStackElemPtr
	VarsBase            c.Int
	ExtFunctions        libxml_2_0.XmlHashTablePtr
	ExtElements         libxml_2_0.XmlHashTablePtr
	ExtInfos            libxml_2_0.XmlHashTablePtr
	Mode                *libxml_2_0.XmlChar
	ModeURI             *libxml_2_0.XmlChar
	DocList             XsltDocumentPtr
	Document            XsltDocumentPtr
	Node                libxml_2_0.XmlNodePtr
	NodeList            libxml_2_0.XmlNodeSetPtr
	Output              libxml_2_0.XmlDocPtr
	Insert              libxml_2_0.XmlNodePtr
	XpathCtxt           libxml_2_0.XmlXPathContextPtr
	State               XsltTransformState
	GlobalVars          libxml_2_0.XmlHashTablePtr
	Inst                libxml_2_0.XmlNodePtr
	Xinclude            c.Int
	OutputFile          *int8
	Profile             c.Int
	Prof                c.Long
	ProfNr              c.Int
	ProfMax             c.Int
	ProfTab             *c.Long
	X_Private           unsafe.Pointer
	ExtrasNr            c.Int
	ExtrasMax           c.Int
	Extras              XsltRuntimeExtraPtr
	StyleList           XsltDocumentPtr
	Sec                 unsafe.Pointer
	Error               unsafe.Pointer
	Errctx              unsafe.Pointer
	Sortfunc            unsafe.Pointer
	TmpRVT              libxml_2_0.XmlDocPtr
	PersistRVT          libxml_2_0.XmlDocPtr
	Ctxtflags           c.Int
	Lasttext            *libxml_2_0.XmlChar
	Lasttsize           c.Int
	Lasttuse            c.Int
	DebugStatus         c.Int
	TraceCode           *c.Ulong
	ParserOptions       c.Int
	Dict                libxml_2_0.XmlDictPtr
	TmpDoc              libxml_2_0.XmlDocPtr
	Internalized        c.Int
	NbKeys              c.Int
	HasTemplKeyPatterns c.Int
	CurrentTemplateRule XsltTemplatePtr
	InitialContextNode  libxml_2_0.XmlNodePtr
	InitialContextDoc   libxml_2_0.XmlDocPtr
	Cache               XsltTransformCachePtr
	ContextVariable     unsafe.Pointer
	LocalRVT            libxml_2_0.XmlDocPtr
	LocalRVTBase        libxml_2_0.XmlDocPtr
	KeyInitLevel        c.Int
	Depth               c.Int
	MaxTemplateDepth    c.Int
	MaxTemplateVars     c.Int
	OpLimit             c.Ulong
	OpCount             c.Ulong
	SourceDocDirty      c.Int
	CurrentId           c.Ulong
	NewLocale           unsafe.Pointer
	FreeLocale          unsafe.Pointer
	GenSortKey          unsafe.Pointer
}
type XsltTransformContext X_XsltTransformContext
type XsltTransformContextPtr *XsltTransformContext

type X_XsltElemPreComp struct {
	Next XsltElemPreCompPtr
	Type XsltStyleType
	Func unsafe.Pointer
	Inst libxml_2_0.XmlNodePtr
	Free unsafe.Pointer
}
type XsltElemPreComp X_XsltElemPreComp
type XsltElemPreCompPtr *XsltElemPreComp
// llgo:type C
type XsltTransformFunction func(XsltTransformContextPtr, libxml_2_0.XmlNodePtr, libxml_2_0.XmlNodePtr, XsltElemPreCompPtr)
// llgo:type C
type XsltSortFunc func(XsltTransformContextPtr, *libxml_2_0.XmlNodePtr, c.Int)
type XsltStyleType c.Int

const (
	XsltStyleTypeXSLTFUNCCOPY           XsltStyleType = 1
	XsltStyleTypeXSLTFUNCSORT           XsltStyleType = 2
	XsltStyleTypeXSLTFUNCTEXT           XsltStyleType = 3
	XsltStyleTypeXSLTFUNCELEMENT        XsltStyleType = 4
	XsltStyleTypeXSLTFUNCATTRIBUTE      XsltStyleType = 5
	XsltStyleTypeXSLTFUNCCOMMENT        XsltStyleType = 6
	XsltStyleTypeXSLTFUNCPI             XsltStyleType = 7
	XsltStyleTypeXSLTFUNCCOPYOF         XsltStyleType = 8
	XsltStyleTypeXSLTFUNCVALUEOF        XsltStyleType = 9
	XsltStyleTypeXSLTFUNCNUMBER         XsltStyleType = 10
	XsltStyleTypeXSLTFUNCAPPLYIMPORTS   XsltStyleType = 11
	XsltStyleTypeXSLTFUNCCALLTEMPLATE   XsltStyleType = 12
	XsltStyleTypeXSLTFUNCAPPLYTEMPLATES XsltStyleType = 13
	XsltStyleTypeXSLTFUNCCHOOSE         XsltStyleType = 14
	XsltStyleTypeXSLTFUNCIF             XsltStyleType = 15
	XsltStyleTypeXSLTFUNCFOREACH        XsltStyleType = 16
	XsltStyleTypeXSLTFUNCDOCUMENT       XsltStyleType = 17
	XsltStyleTypeXSLTFUNCWITHPARAM      XsltStyleType = 18
	XsltStyleTypeXSLTFUNCPARAM          XsltStyleType = 19
	XsltStyleTypeXSLTFUNCVARIABLE       XsltStyleType = 20
	XsltStyleTypeXSLTFUNCWHEN           XsltStyleType = 21
	XsltStyleTypeXSLTFUNCEXTENSION      XsltStyleType = 22
)
// llgo:type C
type XsltElemPreCompDeallocator func(XsltElemPreCompPtr)

type X_XsltStylePreComp struct {
	Next        XsltElemPreCompPtr
	Type        XsltStyleType
	Func        unsafe.Pointer
	Inst        libxml_2_0.XmlNodePtr
	Stype       *libxml_2_0.XmlChar
	HasStype    c.Int
	Number      c.Int
	Order       *libxml_2_0.XmlChar
	HasOrder    c.Int
	Descending  c.Int
	Lang        *libxml_2_0.XmlChar
	HasLang     c.Int
	CaseOrder   *libxml_2_0.XmlChar
	LowerFirst  c.Int
	Use         *libxml_2_0.XmlChar
	HasUse      c.Int
	Noescape    c.Int
	Name        *libxml_2_0.XmlChar
	HasName     c.Int
	Ns          *libxml_2_0.XmlChar
	HasNs       c.Int
	Mode        *libxml_2_0.XmlChar
	ModeURI     *libxml_2_0.XmlChar
	Test        *libxml_2_0.XmlChar
	Templ       XsltTemplatePtr
	Select      *libxml_2_0.XmlChar
	Ver11       c.Int
	Filename    *libxml_2_0.XmlChar
	HasFilename c.Int
	Numdata     XsltNumberData
	Comp        libxml_2_0.XmlXPathCompExprPtr
	NsList      *libxml_2_0.XmlNsPtr
	NsNr        c.Int
}
type XsltStylePreComp X_XsltStylePreComp
type XsltStylePreCompPtr *XsltStylePreComp

type X_XsltStackElem struct {
	Next     *X_XsltStackElem
	Comp     XsltStylePreCompPtr
	Computed c.Int
	Name     *libxml_2_0.XmlChar
	NameURI  *libxml_2_0.XmlChar
	Select   *libxml_2_0.XmlChar
	Tree     libxml_2_0.XmlNodePtr
	Value    libxml_2_0.XmlXPathObjectPtr
	Fragment libxml_2_0.XmlDocPtr
	Level    c.Int
	Context  XsltTransformContextPtr
	Flags    c.Int
}
type XsltStackElem X_XsltStackElem
type XsltStackElemPtr *XsltStackElem

type X_XsltTransformCache struct {
	RVT          libxml_2_0.XmlDocPtr
	NbRVT        c.Int
	StackItems   XsltStackElemPtr
	NbStackItems c.Int
}
type XsltTransformCache X_XsltTransformCache
type XsltTransformCachePtr *XsltTransformCache
type XsltOutputType c.Int

const (
	XsltOutputTypeXSLTOUTPUTXML  XsltOutputType = 0
	XsltOutputTypeXSLTOUTPUTHTML XsltOutputType = 1
	XsltOutputTypeXSLTOUTPUTTEXT XsltOutputType = 2
)
// llgo:type C
type XsltNewLocaleFunc func(*libxml_2_0.XmlChar, c.Int) unsafe.Pointer
// llgo:type C
type XsltFreeLocaleFunc func(unsafe.Pointer)
// llgo:type C
type XsltGenSortKeyFunc func(unsafe.Pointer, *libxml_2_0.XmlChar) *libxml_2_0.XmlChar
type XsltTransformState c.Int

const (
	XsltTransformStateXSLTSTATEOK      XsltTransformState = 0
	XsltTransformStateXSLTSTATEERROR   XsltTransformState = 1
	XsltTransformStateXSLTSTATESTOPPED XsltTransformState = 2
)
//go:linkname XsltNewStylesheet C.xsltNewStylesheet
func XsltNewStylesheet() XsltStylesheetPtr
// llgo:link (*XmlChar).XsltParseStylesheetFile C.xsltParseStylesheetFile
func (recv_ *libxml_2_0.XmlChar) XsltParseStylesheetFile() XsltStylesheetPtr {
	return nil
}
//go:linkname XsltFreeStylesheet C.xsltFreeStylesheet
func XsltFreeStylesheet(style XsltStylesheetPtr)
// llgo:link (*XmlChar).XsltIsBlank C.xsltIsBlank
func (recv_ *libxml_2_0.XmlChar) XsltIsBlank() c.Int {
	return 0
}
//go:linkname XsltFreeStackElemList C.xsltFreeStackElemList
func XsltFreeStackElemList(elem XsltStackElemPtr)
//go:linkname XsltDecimalFormatGetByName C.xsltDecimalFormatGetByName
func XsltDecimalFormatGetByName(style XsltStylesheetPtr, name *libxml_2_0.XmlChar) XsltDecimalFormatPtr
//go:linkname XsltDecimalFormatGetByQName C.xsltDecimalFormatGetByQName
func XsltDecimalFormatGetByQName(style XsltStylesheetPtr, nsUri *libxml_2_0.XmlChar, name *libxml_2_0.XmlChar) XsltDecimalFormatPtr
//go:linkname XsltParseStylesheetProcess C.xsltParseStylesheetProcess
func XsltParseStylesheetProcess(ret XsltStylesheetPtr, doc libxml_2_0.XmlDocPtr) XsltStylesheetPtr
//go:linkname XsltParseStylesheetOutput C.xsltParseStylesheetOutput
func XsltParseStylesheetOutput(style XsltStylesheetPtr, cur libxml_2_0.XmlNodePtr)
//go:linkname XsltParseStylesheetDoc C.xsltParseStylesheetDoc
func XsltParseStylesheetDoc(doc libxml_2_0.XmlDocPtr) XsltStylesheetPtr
//go:linkname XsltParseStylesheetImportedDoc C.xsltParseStylesheetImportedDoc
func XsltParseStylesheetImportedDoc(doc libxml_2_0.XmlDocPtr, style XsltStylesheetPtr) XsltStylesheetPtr
//go:linkname XsltParseStylesheetUser C.xsltParseStylesheetUser
func XsltParseStylesheetUser(style XsltStylesheetPtr, doc libxml_2_0.XmlDocPtr) c.Int
//go:linkname XsltLoadStylesheetPI C.xsltLoadStylesheetPI
func XsltLoadStylesheetPI(doc libxml_2_0.XmlDocPtr) XsltStylesheetPtr
//go:linkname XsltNumberFormat C.xsltNumberFormat
func XsltNumberFormat(ctxt XsltTransformContextPtr, data XsltNumberDataPtr, node libxml_2_0.XmlNodePtr)
//go:linkname XsltFormatNumberConversion C.xsltFormatNumberConversion
func XsltFormatNumberConversion(self XsltDecimalFormatPtr, format *libxml_2_0.XmlChar, number float64, result **libxml_2_0.XmlChar) libxml_2_0.XmlXPathError
//go:linkname XsltParseTemplateContent C.xsltParseTemplateContent
func XsltParseTemplateContent(style XsltStylesheetPtr, templ libxml_2_0.XmlNodePtr)
//go:linkname XsltAllocateExtra C.xsltAllocateExtra
func XsltAllocateExtra(style XsltStylesheetPtr) c.Int
//go:linkname XsltAllocateExtraCtxt C.xsltAllocateExtraCtxt
func XsltAllocateExtraCtxt(ctxt XsltTransformContextPtr) c.Int
//go:linkname XsltCreateRVT C.xsltCreateRVT
func XsltCreateRVT(ctxt XsltTransformContextPtr) libxml_2_0.XmlDocPtr
//go:linkname XsltRegisterTmpRVT C.xsltRegisterTmpRVT
func XsltRegisterTmpRVT(ctxt XsltTransformContextPtr, RVT libxml_2_0.XmlDocPtr) c.Int
//go:linkname XsltRegisterLocalRVT C.xsltRegisterLocalRVT
func XsltRegisterLocalRVT(ctxt XsltTransformContextPtr, RVT libxml_2_0.XmlDocPtr) c.Int
//go:linkname XsltRegisterPersistRVT C.xsltRegisterPersistRVT
func XsltRegisterPersistRVT(ctxt XsltTransformContextPtr, RVT libxml_2_0.XmlDocPtr) c.Int
//go:linkname XsltExtensionInstructionResultRegister C.xsltExtensionInstructionResultRegister
func XsltExtensionInstructionResultRegister(ctxt XsltTransformContextPtr, obj libxml_2_0.XmlXPathObjectPtr) c.Int
//go:linkname XsltExtensionInstructionResultFinalize C.xsltExtensionInstructionResultFinalize
func XsltExtensionInstructionResultFinalize(ctxt XsltTransformContextPtr) c.Int
//go:linkname XsltFlagRVTs C.xsltFlagRVTs
func XsltFlagRVTs(ctxt XsltTransformContextPtr, obj libxml_2_0.XmlXPathObjectPtr, val c.Int) c.Int
//go:linkname XsltFreeRVTs C.xsltFreeRVTs
func XsltFreeRVTs(ctxt XsltTransformContextPtr)
//go:linkname XsltReleaseRVT C.xsltReleaseRVT
func XsltReleaseRVT(ctxt XsltTransformContextPtr, RVT libxml_2_0.XmlDocPtr)
//go:linkname XsltCompileAttr C.xsltCompileAttr
func XsltCompileAttr(style XsltStylesheetPtr, attr libxml_2_0.XmlAttrPtr)
//go:linkname XsltEvalAVT C.xsltEvalAVT
func XsltEvalAVT(ctxt XsltTransformContextPtr, avt unsafe.Pointer, node libxml_2_0.XmlNodePtr) *libxml_2_0.XmlChar
//go:linkname XsltFreeAVTList C.xsltFreeAVTList
func XsltFreeAVTList(avt unsafe.Pointer)
//go:linkname XsltUninit C.xsltUninit
func XsltUninit()
/************************************************************************
 *									*
 *  Transformation-time functions for *internal* use only               *
 *									*
 ************************************************************************/
//go:linkname XsltInitCtxtKey C.xsltInitCtxtKey
func XsltInitCtxtKey(ctxt XsltTransformContextPtr, doc XsltDocumentPtr, keyd XsltKeyDefPtr) c.Int
//go:linkname XsltInitAllDocKeys C.xsltInitAllDocKeys
func XsltInitAllDocKeys(ctxt XsltTransformContextPtr) c.Int
