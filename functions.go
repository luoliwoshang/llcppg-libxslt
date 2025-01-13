package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	"unsafe"
)
//go:linkname XsltXPathFunctionLookup C.xsltXPathFunctionLookup
func XsltXPathFunctionLookup(vctxt unsafe.Pointer, name *libxml_2_0.XmlChar, ns_uri *libxml_2_0.XmlChar) libxml_2_0.XmlXPathFunction
//go:linkname XsltDocumentFunction C.xsltDocumentFunction
func XsltDocumentFunction(ctxt libxml_2_0.XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XsltKeyFunction C.xsltKeyFunction
func XsltKeyFunction(ctxt libxml_2_0.XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XsltUnparsedEntityURIFunction C.xsltUnparsedEntityURIFunction
func XsltUnparsedEntityURIFunction(ctxt libxml_2_0.XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XsltFormatNumberFunction C.xsltFormatNumberFunction
func XsltFormatNumberFunction(ctxt libxml_2_0.XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XsltGenerateIdFunction C.xsltGenerateIdFunction
func XsltGenerateIdFunction(ctxt libxml_2_0.XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XsltSystemPropertyFunction C.xsltSystemPropertyFunction
func XsltSystemPropertyFunction(ctxt libxml_2_0.XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XsltElementAvailableFunction C.xsltElementAvailableFunction
func XsltElementAvailableFunction(ctxt libxml_2_0.XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XsltFunctionAvailableFunction C.xsltFunctionAvailableFunction
func XsltFunctionAvailableFunction(ctxt libxml_2_0.XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XsltRegisterAllFunctions C.xsltRegisterAllFunctions
func XsltRegisterAllFunctions(ctxt libxml_2_0.XmlXPathContextPtr)
