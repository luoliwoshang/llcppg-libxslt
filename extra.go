package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	_ "unsafe"
)
//go:linkname XsltFunctionNodeSet C.xsltFunctionNodeSet
func XsltFunctionNodeSet(ctxt libxml_2_0.XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XsltDebug C.xsltDebug
func XsltDebug(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltRegisterExtras C.xsltRegisterExtras
func XsltRegisterExtras(ctxt XsltTransformContextPtr)
//go:linkname XsltRegisterAllExtras C.xsltRegisterAllExtras
func XsltRegisterAllExtras()
