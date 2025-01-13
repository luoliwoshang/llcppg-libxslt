package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	_ "unsafe"
)
//go:linkname XsltParseStylesheetImport C.xsltParseStylesheetImport
func XsltParseStylesheetImport(style XsltStylesheetPtr, cur libxml_2_0.XmlNodePtr) c.Int
//go:linkname XsltParseStylesheetInclude C.xsltParseStylesheetInclude
func XsltParseStylesheetInclude(style XsltStylesheetPtr, cur libxml_2_0.XmlNodePtr) c.Int
//go:linkname XsltNextImport C.xsltNextImport
func XsltNextImport(style XsltStylesheetPtr) XsltStylesheetPtr
//go:linkname XsltNeedElemSpaceHandling C.xsltNeedElemSpaceHandling
func XsltNeedElemSpaceHandling(ctxt XsltTransformContextPtr) c.Int
//go:linkname XsltFindElemSpaceHandling C.xsltFindElemSpaceHandling
func XsltFindElemSpaceHandling(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr) c.Int
//go:linkname XsltFindTemplate C.xsltFindTemplate
func XsltFindTemplate(ctxt XsltTransformContextPtr, name *libxml_2_0.XmlChar, nameURI *libxml_2_0.XmlChar) XsltTemplatePtr
