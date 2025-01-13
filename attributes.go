package libxslt

import (
	"github.com/luoliwoshang/llcppg-libxml"
	_ "unsafe"
)
//go:linkname XsltParseStylesheetAttributeSet C.xsltParseStylesheetAttributeSet
func XsltParseStylesheetAttributeSet(style XsltStylesheetPtr, cur libxml_2_0.XmlNodePtr)
//go:linkname XsltFreeAttributeSetsHashes C.xsltFreeAttributeSetsHashes
func XsltFreeAttributeSetsHashes(style XsltStylesheetPtr)
//go:linkname XsltApplyAttributeSet C.xsltApplyAttributeSet
func XsltApplyAttributeSet(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, attributes *libxml_2_0.XmlChar)
//go:linkname XsltResolveStylesheetAttributeSet C.xsltResolveStylesheetAttributeSet
func XsltResolveStylesheetAttributeSet(style XsltStylesheetPtr)
