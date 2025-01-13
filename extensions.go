package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	"unsafe"
)
/**
 * xsltInitGlobals:
 *
 * Initialize the global variables for extensions
 *
 */
//go:linkname XsltInitGlobals C.xsltInitGlobals
func XsltInitGlobals()
// llgo:type C
type XsltStyleExtInitFunction func(XsltStylesheetPtr, *libxml_2_0.XmlChar) unsafe.Pointer
// llgo:type C
type XsltStyleExtShutdownFunction func(XsltStylesheetPtr, *libxml_2_0.XmlChar, unsafe.Pointer)
// llgo:type C
type XsltExtInitFunction func(XsltTransformContextPtr, *libxml_2_0.XmlChar) unsafe.Pointer
// llgo:type C
type XsltExtShutdownFunction func(XsltTransformContextPtr, *libxml_2_0.XmlChar, unsafe.Pointer)
// llgo:link (*XmlChar).XsltRegisterExtModule C.xsltRegisterExtModule
func (recv_ *libxml_2_0.XmlChar) XsltRegisterExtModule(initFunc XsltExtInitFunction, shutdownFunc XsltExtShutdownFunction) c.Int {
	return 0
}
// llgo:link (*XmlChar).XsltRegisterExtModuleFull C.xsltRegisterExtModuleFull
func (recv_ *libxml_2_0.XmlChar) XsltRegisterExtModuleFull(initFunc XsltExtInitFunction, shutdownFunc XsltExtShutdownFunction, styleInitFunc XsltStyleExtInitFunction, styleShutdownFunc XsltStyleExtShutdownFunction) c.Int {
	return 0
}
// llgo:link (*XmlChar).XsltUnregisterExtModule C.xsltUnregisterExtModule
func (recv_ *libxml_2_0.XmlChar) XsltUnregisterExtModule() c.Int {
	return 0
}
//go:linkname XsltGetExtData C.xsltGetExtData
func XsltGetExtData(ctxt XsltTransformContextPtr, URI *libxml_2_0.XmlChar) unsafe.Pointer
//go:linkname XsltStyleGetExtData C.xsltStyleGetExtData
func XsltStyleGetExtData(style XsltStylesheetPtr, URI *libxml_2_0.XmlChar) unsafe.Pointer
//go:linkname XsltShutdownCtxtExts C.xsltShutdownCtxtExts
func XsltShutdownCtxtExts(ctxt XsltTransformContextPtr)
//go:linkname XsltShutdownExts C.xsltShutdownExts
func XsltShutdownExts(style XsltStylesheetPtr)
//go:linkname XsltXPathGetTransformContext C.xsltXPathGetTransformContext
func XsltXPathGetTransformContext(ctxt libxml_2_0.XmlXPathParserContextPtr) XsltTransformContextPtr
// llgo:link (*XmlChar).XsltRegisterExtModuleFunction C.xsltRegisterExtModuleFunction
func (recv_ *libxml_2_0.XmlChar) XsltRegisterExtModuleFunction(URI *libxml_2_0.XmlChar, function libxml_2_0.XmlXPathFunction) c.Int {
	return 0
}
// llgo:link (*XmlChar).XsltExtModuleFunctionLookup C.xsltExtModuleFunctionLookup
func (recv_ *libxml_2_0.XmlChar) XsltExtModuleFunctionLookup(URI *libxml_2_0.XmlChar) libxml_2_0.XmlXPathFunction {
	return nil
}
// llgo:link (*XmlChar).XsltUnregisterExtModuleFunction C.xsltUnregisterExtModuleFunction
func (recv_ *libxml_2_0.XmlChar) XsltUnregisterExtModuleFunction(URI *libxml_2_0.XmlChar) c.Int {
	return 0
}
// llgo:type C
type XsltPreComputeFunction func(XsltStylesheetPtr, libxml_2_0.XmlNodePtr, XsltTransformFunction) XsltElemPreCompPtr
//go:linkname XsltNewElemPreComp C.xsltNewElemPreComp
func XsltNewElemPreComp(style XsltStylesheetPtr, inst libxml_2_0.XmlNodePtr, function XsltTransformFunction) XsltElemPreCompPtr
//go:linkname XsltInitElemPreComp C.xsltInitElemPreComp
func XsltInitElemPreComp(comp XsltElemPreCompPtr, style XsltStylesheetPtr, inst libxml_2_0.XmlNodePtr, function XsltTransformFunction, freeFunc XsltElemPreCompDeallocator)
// llgo:link (*XmlChar).XsltRegisterExtModuleElement C.xsltRegisterExtModuleElement
func (recv_ *libxml_2_0.XmlChar) XsltRegisterExtModuleElement(URI *libxml_2_0.XmlChar, precomp XsltPreComputeFunction, transform XsltTransformFunction) c.Int {
	return 0
}
//go:linkname XsltExtElementLookup C.xsltExtElementLookup
func XsltExtElementLookup(ctxt XsltTransformContextPtr, name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar) XsltTransformFunction
// llgo:link (*XmlChar).XsltExtModuleElementLookup C.xsltExtModuleElementLookup
func (recv_ *libxml_2_0.XmlChar) XsltExtModuleElementLookup(URI *libxml_2_0.XmlChar) XsltTransformFunction {
	return nil
}
// llgo:link (*XmlChar).XsltExtModuleElementPreComputeLookup C.xsltExtModuleElementPreComputeLookup
func (recv_ *libxml_2_0.XmlChar) XsltExtModuleElementPreComputeLookup(URI *libxml_2_0.XmlChar) XsltPreComputeFunction {
	return nil
}
// llgo:link (*XmlChar).XsltUnregisterExtModuleElement C.xsltUnregisterExtModuleElement
func (recv_ *libxml_2_0.XmlChar) XsltUnregisterExtModuleElement(URI *libxml_2_0.XmlChar) c.Int {
	return 0
}
// llgo:type C
type XsltTopLevelFunction func(XsltStylesheetPtr, libxml_2_0.XmlNodePtr)
// llgo:link (*XmlChar).XsltRegisterExtModuleTopLevel C.xsltRegisterExtModuleTopLevel
func (recv_ *libxml_2_0.XmlChar) XsltRegisterExtModuleTopLevel(URI *libxml_2_0.XmlChar, function XsltTopLevelFunction) c.Int {
	return 0
}
// llgo:link (*XmlChar).XsltExtModuleTopLevelLookup C.xsltExtModuleTopLevelLookup
func (recv_ *libxml_2_0.XmlChar) XsltExtModuleTopLevelLookup(URI *libxml_2_0.XmlChar) XsltTopLevelFunction {
	return nil
}
// llgo:link (*XmlChar).XsltUnregisterExtModuleTopLevel C.xsltUnregisterExtModuleTopLevel
func (recv_ *libxml_2_0.XmlChar) XsltUnregisterExtModuleTopLevel(URI *libxml_2_0.XmlChar) c.Int {
	return 0
}
//go:linkname XsltRegisterExtFunction C.xsltRegisterExtFunction
func XsltRegisterExtFunction(ctxt XsltTransformContextPtr, name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar, function libxml_2_0.XmlXPathFunction) c.Int
//go:linkname XsltRegisterExtElement C.xsltRegisterExtElement
func XsltRegisterExtElement(ctxt XsltTransformContextPtr, name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar, function XsltTransformFunction) c.Int
//go:linkname XsltRegisterExtPrefix C.xsltRegisterExtPrefix
func XsltRegisterExtPrefix(style XsltStylesheetPtr, prefix *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar) c.Int
//go:linkname XsltCheckExtPrefix C.xsltCheckExtPrefix
func XsltCheckExtPrefix(style XsltStylesheetPtr, URI *libxml_2_0.XmlChar) c.Int
//go:linkname XsltCheckExtURI C.xsltCheckExtURI
func XsltCheckExtURI(style XsltStylesheetPtr, URI *libxml_2_0.XmlChar) c.Int
//go:linkname XsltInitCtxtExts C.xsltInitCtxtExts
func XsltInitCtxtExts(ctxt XsltTransformContextPtr) c.Int
//go:linkname XsltFreeCtxtExts C.xsltFreeCtxtExts
func XsltFreeCtxtExts(ctxt XsltTransformContextPtr)
//go:linkname XsltFreeExts C.xsltFreeExts
func XsltFreeExts(style XsltStylesheetPtr)
//go:linkname XsltPreComputeExtModuleElement C.xsltPreComputeExtModuleElement
func XsltPreComputeExtModuleElement(style XsltStylesheetPtr, inst libxml_2_0.XmlNodePtr) XsltElemPreCompPtr
//go:linkname XsltGetExtInfo C.xsltGetExtInfo
func XsltGetExtInfo(style XsltStylesheetPtr, URI *libxml_2_0.XmlChar) libxml_2_0.XmlHashTablePtr
/**
 * Test of the extension module API
 */
//go:linkname XsltRegisterTestModule C.xsltRegisterTestModule
func XsltRegisterTestModule()
//go:linkname XsltDebugDumpExtensions C.xsltDebugDumpExtensions
func XsltDebugDumpExtensions(output *c.FILE)
