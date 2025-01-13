package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	"unsafe"
)

const XSLTRVTLOCAL c.Int = 1
const XSLTRVTFUNCRESULT c.Int = 2
const XSLTRVTGLOBAL c.Int = 3
//go:linkname XsltEvalGlobalVariables C.xsltEvalGlobalVariables
func XsltEvalGlobalVariables(ctxt XsltTransformContextPtr) c.Int
//go:linkname XsltEvalUserParams C.xsltEvalUserParams
func XsltEvalUserParams(ctxt XsltTransformContextPtr, params **int8) c.Int
//go:linkname XsltQuoteUserParams C.xsltQuoteUserParams
func XsltQuoteUserParams(ctxt XsltTransformContextPtr, params **int8) c.Int
//go:linkname XsltEvalOneUserParam C.xsltEvalOneUserParam
func XsltEvalOneUserParam(ctxt XsltTransformContextPtr, name *libxml_2_0.XmlChar, value *libxml_2_0.XmlChar) c.Int
//go:linkname XsltQuoteOneUserParam C.xsltQuoteOneUserParam
func XsltQuoteOneUserParam(ctxt XsltTransformContextPtr, name *libxml_2_0.XmlChar, value *libxml_2_0.XmlChar) c.Int
//go:linkname XsltParseGlobalVariable C.xsltParseGlobalVariable
func XsltParseGlobalVariable(style XsltStylesheetPtr, cur libxml_2_0.XmlNodePtr)
//go:linkname XsltParseGlobalParam C.xsltParseGlobalParam
func XsltParseGlobalParam(style XsltStylesheetPtr, cur libxml_2_0.XmlNodePtr)
//go:linkname XsltParseStylesheetVariable C.xsltParseStylesheetVariable
func XsltParseStylesheetVariable(ctxt XsltTransformContextPtr, cur libxml_2_0.XmlNodePtr)
//go:linkname XsltParseStylesheetParam C.xsltParseStylesheetParam
func XsltParseStylesheetParam(ctxt XsltTransformContextPtr, cur libxml_2_0.XmlNodePtr)
//go:linkname XsltParseStylesheetCallerParam C.xsltParseStylesheetCallerParam
func XsltParseStylesheetCallerParam(ctxt XsltTransformContextPtr, cur libxml_2_0.XmlNodePtr) XsltStackElemPtr
//go:linkname XsltAddStackElemList C.xsltAddStackElemList
func XsltAddStackElemList(ctxt XsltTransformContextPtr, elems XsltStackElemPtr) c.Int
//go:linkname XsltFreeGlobalVariables C.xsltFreeGlobalVariables
func XsltFreeGlobalVariables(ctxt XsltTransformContextPtr)
//go:linkname XsltVariableLookup C.xsltVariableLookup
func XsltVariableLookup(ctxt XsltTransformContextPtr, name *libxml_2_0.XmlChar, ns_uri *libxml_2_0.XmlChar) libxml_2_0.XmlXPathObjectPtr
//go:linkname XsltXPathVariableLookup C.xsltXPathVariableLookup
func XsltXPathVariableLookup(ctxt unsafe.Pointer, name *libxml_2_0.XmlChar, ns_uri *libxml_2_0.XmlChar) libxml_2_0.XmlXPathObjectPtr
