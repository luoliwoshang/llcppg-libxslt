package libxslt

import (
	"github.com/luoliwoshang/llcppg-libxml"
	_ "unsafe"
)
//go:linkname XsltDocumentComp C.xsltDocumentComp
func XsltDocumentComp(style XsltStylesheetPtr, inst libxml_2_0.XmlNodePtr, function XsltTransformFunction) XsltElemPreCompPtr
//go:linkname XsltStylePreCompute C.xsltStylePreCompute
func XsltStylePreCompute(style XsltStylesheetPtr, inst libxml_2_0.XmlNodePtr)
//go:linkname XsltFreeStylePreComps C.xsltFreeStylePreComps
func XsltFreeStylePreComps(style XsltStylesheetPtr)
