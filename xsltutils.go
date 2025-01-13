package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	"unsafe"
)
//go:linkname XsltGetNsProp C.xsltGetNsProp
func XsltGetNsProp(node libxml_2_0.XmlNodePtr, name *libxml_2_0.XmlChar, nameSpace *libxml_2_0.XmlChar) *libxml_2_0.XmlChar
//go:linkname XsltGetCNsProp C.xsltGetCNsProp
func XsltGetCNsProp(style XsltStylesheetPtr, node libxml_2_0.XmlNodePtr, name *libxml_2_0.XmlChar, nameSpace *libxml_2_0.XmlChar) *libxml_2_0.XmlChar
//go:linkname XsltGetUTF8Char C.xsltGetUTF8Char
func XsltGetUTF8Char(utf *int8, len *c.Int) c.Int

type XsltDebugTraceCodes c.Int

const (
	XsltDebugTraceCodesXSLTTRACEALL            XsltDebugTraceCodes = -1
	XsltDebugTraceCodesXSLTTRACENONE           XsltDebugTraceCodes = 0
	XsltDebugTraceCodesXSLTTRACECOPYTEXT       XsltDebugTraceCodes = 1
	XsltDebugTraceCodesXSLTTRACEPROCESSNODE    XsltDebugTraceCodes = 2
	XsltDebugTraceCodesXSLTTRACEAPPLYTEMPLATE  XsltDebugTraceCodes = 4
	XsltDebugTraceCodesXSLTTRACECOPY           XsltDebugTraceCodes = 8
	XsltDebugTraceCodesXSLTTRACECOMMENT        XsltDebugTraceCodes = 16
	XsltDebugTraceCodesXSLTTRACEPI             XsltDebugTraceCodes = 32
	XsltDebugTraceCodesXSLTTRACECOPYOF         XsltDebugTraceCodes = 64
	XsltDebugTraceCodesXSLTTRACEVALUEOF        XsltDebugTraceCodes = 128
	XsltDebugTraceCodesXSLTTRACECALLTEMPLATE   XsltDebugTraceCodes = 256
	XsltDebugTraceCodesXSLTTRACEAPPLYTEMPLATES XsltDebugTraceCodes = 512
	XsltDebugTraceCodesXSLTTRACECHOOSE         XsltDebugTraceCodes = 1024
	XsltDebugTraceCodesXSLTTRACEIF             XsltDebugTraceCodes = 2048
	XsltDebugTraceCodesXSLTTRACEFOREACH        XsltDebugTraceCodes = 4096
	XsltDebugTraceCodesXSLTTRACESTRIPSPACES    XsltDebugTraceCodes = 8192
	XsltDebugTraceCodesXSLTTRACETEMPLATES      XsltDebugTraceCodes = 16384
	XsltDebugTraceCodesXSLTTRACEKEYS           XsltDebugTraceCodes = 32768
	XsltDebugTraceCodesXSLTTRACEVARIABLES      XsltDebugTraceCodes = 65536
)
// llgo:link XsltDebugTraceCodes.XsltDebugSetDefaultTrace C.xsltDebugSetDefaultTrace
func (recv_ XsltDebugTraceCodes) XsltDebugSetDefaultTrace() {
}
//go:linkname XsltDebugGetDefaultTrace C.xsltDebugGetDefaultTrace
func XsltDebugGetDefaultTrace() XsltDebugTraceCodes
//go:linkname XsltPrintErrorContext C.xsltPrintErrorContext
func XsltPrintErrorContext(ctxt XsltTransformContextPtr, style XsltStylesheetPtr, node libxml_2_0.XmlNodePtr)
//go:linkname XsltMessage C.xsltMessage
func XsltMessage(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr)
//go:linkname XsltSetGenericErrorFunc C.xsltSetGenericErrorFunc
func XsltSetGenericErrorFunc(ctx unsafe.Pointer, handler libxml_2_0.XmlGenericErrorFunc)
//go:linkname XsltSetGenericDebugFunc C.xsltSetGenericDebugFunc
func XsltSetGenericDebugFunc(ctx unsafe.Pointer, handler libxml_2_0.XmlGenericErrorFunc)
//go:linkname XsltSetTransformErrorFunc C.xsltSetTransformErrorFunc
func XsltSetTransformErrorFunc(ctxt XsltTransformContextPtr, ctx unsafe.Pointer, handler libxml_2_0.XmlGenericErrorFunc)
//go:linkname XsltTransformError C.xsltTransformError
func XsltTransformError(ctxt XsltTransformContextPtr, style XsltStylesheetPtr, node libxml_2_0.XmlNodePtr, msg *int8, __llgo_va_list ...interface{})
//go:linkname XsltSetCtxtParseOptions C.xsltSetCtxtParseOptions
func XsltSetCtxtParseOptions(ctxt XsltTransformContextPtr, options c.Int) c.Int
//go:linkname XsltDocumentSortFunction C.xsltDocumentSortFunction
func XsltDocumentSortFunction(list libxml_2_0.XmlNodeSetPtr)
//go:linkname XsltSetSortFunc C.xsltSetSortFunc
func XsltSetSortFunc(handler XsltSortFunc)
//go:linkname XsltSetCtxtSortFunc C.xsltSetCtxtSortFunc
func XsltSetCtxtSortFunc(ctxt XsltTransformContextPtr, handler XsltSortFunc)
//go:linkname XsltSetCtxtLocaleHandlers C.xsltSetCtxtLocaleHandlers
func XsltSetCtxtLocaleHandlers(ctxt XsltTransformContextPtr, newLocale XsltNewLocaleFunc, freeLocale XsltFreeLocaleFunc, genSortKey XsltGenSortKeyFunc)
//go:linkname XsltDefaultSortFunction C.xsltDefaultSortFunction
func XsltDefaultSortFunction(ctxt XsltTransformContextPtr, sorts *libxml_2_0.XmlNodePtr, nbsorts c.Int)
//go:linkname XsltDoSortFunction C.xsltDoSortFunction
func XsltDoSortFunction(ctxt XsltTransformContextPtr, sorts *libxml_2_0.XmlNodePtr, nbsorts c.Int)
//go:linkname XsltComputeSortResult C.xsltComputeSortResult
func XsltComputeSortResult(ctxt XsltTransformContextPtr, sort libxml_2_0.XmlNodePtr) *libxml_2_0.XmlXPathObjectPtr
//go:linkname XsltSplitQName C.xsltSplitQName
func XsltSplitQName(dict libxml_2_0.XmlDictPtr, name *libxml_2_0.XmlChar, prefix **libxml_2_0.XmlChar) *libxml_2_0.XmlChar
//go:linkname XsltGetQNameURI C.xsltGetQNameURI
func XsltGetQNameURI(node libxml_2_0.XmlNodePtr, name **libxml_2_0.XmlChar) *libxml_2_0.XmlChar
//go:linkname XsltGetQNameURI2 C.xsltGetQNameURI2
func XsltGetQNameURI2(style XsltStylesheetPtr, node libxml_2_0.XmlNodePtr, name **libxml_2_0.XmlChar) *libxml_2_0.XmlChar
//go:linkname XsltSaveResultTo C.xsltSaveResultTo
func XsltSaveResultTo(buf libxml_2_0.XmlOutputBufferPtr, result libxml_2_0.XmlDocPtr, style XsltStylesheetPtr) c.Int
//go:linkname XsltSaveResultToFilename C.xsltSaveResultToFilename
func XsltSaveResultToFilename(URI *int8, result libxml_2_0.XmlDocPtr, style XsltStylesheetPtr, compression c.Int) c.Int
//go:linkname XsltSaveResultToFile C.xsltSaveResultToFile
func XsltSaveResultToFile(file *c.FILE, result libxml_2_0.XmlDocPtr, style XsltStylesheetPtr) c.Int
//go:linkname XsltSaveResultToFd C.xsltSaveResultToFd
func XsltSaveResultToFd(fd c.Int, result libxml_2_0.XmlDocPtr, style XsltStylesheetPtr) c.Int
//go:linkname XsltSaveResultToString C.xsltSaveResultToString
func XsltSaveResultToString(doc_txt_ptr **libxml_2_0.XmlChar, doc_txt_len *c.Int, result libxml_2_0.XmlDocPtr, style XsltStylesheetPtr) c.Int
//go:linkname XsltXPathCompile C.xsltXPathCompile
func XsltXPathCompile(style XsltStylesheetPtr, str *libxml_2_0.XmlChar) libxml_2_0.XmlXPathCompExprPtr
//go:linkname XsltXPathCompileFlags C.xsltXPathCompileFlags
func XsltXPathCompileFlags(style XsltStylesheetPtr, str *libxml_2_0.XmlChar, flags c.Int) libxml_2_0.XmlXPathCompExprPtr
//go:linkname XsltSaveProfiling C.xsltSaveProfiling
func XsltSaveProfiling(ctxt XsltTransformContextPtr, output *c.FILE)
//go:linkname XsltGetProfileInformation C.xsltGetProfileInformation
func XsltGetProfileInformation(ctxt XsltTransformContextPtr) libxml_2_0.XmlDocPtr
//go:linkname XsltTimestamp C.xsltTimestamp
func XsltTimestamp() c.Long
//go:linkname XsltCalibrateAdjust C.xsltCalibrateAdjust
func XsltCalibrateAdjust(delta c.Long)

type XsltDebugStatusCodes c.Int

const (
	XsltDebugStatusCodesXSLTDEBUGNONE       XsltDebugStatusCodes = 0
	XsltDebugStatusCodesXSLTDEBUGINIT       XsltDebugStatusCodes = 1
	XsltDebugStatusCodesXSLTDEBUGSTEP       XsltDebugStatusCodes = 2
	XsltDebugStatusCodesXSLTDEBUGSTEPOUT    XsltDebugStatusCodes = 3
	XsltDebugStatusCodesXSLTDEBUGNEXT       XsltDebugStatusCodes = 4
	XsltDebugStatusCodesXSLTDEBUGSTOP       XsltDebugStatusCodes = 5
	XsltDebugStatusCodesXSLTDEBUGCONT       XsltDebugStatusCodes = 6
	XsltDebugStatusCodesXSLTDEBUGRUN        XsltDebugStatusCodes = 7
	XsltDebugStatusCodesXSLTDEBUGRUNRESTART XsltDebugStatusCodes = 8
	XsltDebugStatusCodesXSLTDEBUGQUIT       XsltDebugStatusCodes = 9
)
// llgo:type C
type XsltHandleDebuggerCallback func(libxml_2_0.XmlNodePtr, libxml_2_0.XmlNodePtr, XsltTemplatePtr, XsltTransformContextPtr)
// llgo:type C
type XsltAddCallCallback func(XsltTemplatePtr, libxml_2_0.XmlNodePtr) c.Int
// llgo:type C
type XsltDropCallCallback func()
//go:linkname XsltGetDebuggerStatus C.xsltGetDebuggerStatus
func XsltGetDebuggerStatus() c.Int
//go:linkname XsltSetDebuggerStatus C.xsltSetDebuggerStatus
func XsltSetDebuggerStatus(value c.Int)
//go:linkname XsltSetDebuggerCallbacks C.xsltSetDebuggerCallbacks
func XsltSetDebuggerCallbacks(no c.Int, block unsafe.Pointer) c.Int
//go:linkname XslAddCall C.xslAddCall
func XslAddCall(templ XsltTemplatePtr, source libxml_2_0.XmlNodePtr) c.Int
//go:linkname XslDropCall C.xslDropCall
func XslDropCall()
