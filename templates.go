package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	_ "unsafe"
)
//go:linkname XsltEvalXPathPredicate C.xsltEvalXPathPredicate
func XsltEvalXPathPredicate(ctxt XsltTransformContextPtr, comp libxml_2_0.XmlXPathCompExprPtr, nsList *libxml_2_0.XmlNsPtr, nsNr c.Int) c.Int
//go:linkname XsltEvalTemplateString C.xsltEvalTemplateString
func XsltEvalTemplateString(ctxt XsltTransformContextPtr, contextNode libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr) *libxml_2_0.XmlChar
//go:linkname XsltEvalAttrValueTemplate C.xsltEvalAttrValueTemplate
func XsltEvalAttrValueTemplate(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, name *libxml_2_0.XmlChar, ns *libxml_2_0.XmlChar) *libxml_2_0.XmlChar
//go:linkname XsltEvalStaticAttrValueTemplate C.xsltEvalStaticAttrValueTemplate
func XsltEvalStaticAttrValueTemplate(style XsltStylesheetPtr, node libxml_2_0.XmlNodePtr, name *libxml_2_0.XmlChar, ns *libxml_2_0.XmlChar, found *c.Int) *libxml_2_0.XmlChar
//go:linkname XsltEvalXPathString C.xsltEvalXPathString
func XsltEvalXPathString(ctxt XsltTransformContextPtr, comp libxml_2_0.XmlXPathCompExprPtr) *libxml_2_0.XmlChar
//go:linkname XsltEvalXPathStringNs C.xsltEvalXPathStringNs
func XsltEvalXPathStringNs(ctxt XsltTransformContextPtr, comp libxml_2_0.XmlXPathCompExprPtr, nsNr c.Int, nsList *libxml_2_0.XmlNsPtr) *libxml_2_0.XmlChar
//go:linkname XsltTemplateProcess C.xsltTemplateProcess
func XsltTemplateProcess(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr) *libxml_2_0.XmlNodePtr
//go:linkname XsltAttrListTemplateProcess C.xsltAttrListTemplateProcess
func XsltAttrListTemplateProcess(ctxt XsltTransformContextPtr, target libxml_2_0.XmlNodePtr, cur libxml_2_0.XmlAttrPtr) libxml_2_0.XmlAttrPtr
//go:linkname XsltAttrTemplateProcess C.xsltAttrTemplateProcess
func XsltAttrTemplateProcess(ctxt XsltTransformContextPtr, target libxml_2_0.XmlNodePtr, attr libxml_2_0.XmlAttrPtr) libxml_2_0.XmlAttrPtr
//go:linkname XsltAttrTemplateValueProcess C.xsltAttrTemplateValueProcess
func XsltAttrTemplateValueProcess(ctxt XsltTransformContextPtr, attr *libxml_2_0.XmlChar) *libxml_2_0.XmlChar
//go:linkname XsltAttrTemplateValueProcessNode C.xsltAttrTemplateValueProcessNode
func XsltAttrTemplateValueProcessNode(ctxt XsltTransformContextPtr, str *libxml_2_0.XmlChar, node libxml_2_0.XmlNodePtr) *libxml_2_0.XmlChar
