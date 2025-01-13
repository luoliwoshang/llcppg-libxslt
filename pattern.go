package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	"unsafe"
)

type XsltCompMatch X_XsltCompMatch
type XsltCompMatchPtr *XsltCompMatch
// llgo:link (*XmlChar).XsltCompilePattern C.xsltCompilePattern
func (recv_ *libxml_2_0.XmlChar) XsltCompilePattern(doc libxml_2_0.XmlDocPtr, node libxml_2_0.XmlNodePtr, style XsltStylesheetPtr, runtime XsltTransformContextPtr) XsltCompMatchPtr {
	return nil
}
//go:linkname XsltFreeCompMatchList C.xsltFreeCompMatchList
func XsltFreeCompMatchList(comp XsltCompMatchPtr)
//go:linkname XsltTestCompMatchList C.xsltTestCompMatchList
func XsltTestCompMatchList(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, comp XsltCompMatchPtr) c.Int
//go:linkname XsltCompMatchClearCache C.xsltCompMatchClearCache
func XsltCompMatchClearCache(ctxt XsltTransformContextPtr, comp XsltCompMatchPtr)
//go:linkname XsltNormalizeCompSteps C.xsltNormalizeCompSteps
func XsltNormalizeCompSteps(payload unsafe.Pointer, data unsafe.Pointer, name *libxml_2_0.XmlChar)
//go:linkname XsltAddTemplate C.xsltAddTemplate
func XsltAddTemplate(style XsltStylesheetPtr, cur XsltTemplatePtr, mode *libxml_2_0.XmlChar, modeURI *libxml_2_0.XmlChar) c.Int
//go:linkname XsltGetTemplate C.xsltGetTemplate
func XsltGetTemplate(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, style XsltStylesheetPtr) XsltTemplatePtr
//go:linkname XsltFreeTemplateHashes C.xsltFreeTemplateHashes
func XsltFreeTemplateHashes(style XsltStylesheetPtr)
//go:linkname XsltCleanupTemplates C.xsltCleanupTemplates
func XsltCleanupTemplates(style XsltStylesheetPtr)
