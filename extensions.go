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
//go:linkname XsltRegisterExtModule C.xsltRegisterExtModule
func XsltRegisterExtModule(URI *libxml_2_0.XmlChar, initFunc XsltExtInitFunction, shutdownFunc XsltExtShutdownFunction) c.Int
//go:linkname XsltRegisterExtModuleFull C.xsltRegisterExtModuleFull
func XsltRegisterExtModuleFull(URI *libxml_2_0.XmlChar, initFunc XsltExtInitFunction, shutdownFunc XsltExtShutdownFunction, styleInitFunc XsltStyleExtInitFunction, styleShutdownFunc XsltStyleExtShutdownFunction) c.Int
//go:linkname XsltUnregisterExtModule C.xsltUnregisterExtModule
func XsltUnregisterExtModule(URI *libxml_2_0.XmlChar) c.Int
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
//go:linkname XsltRegisterExtModuleFunction C.xsltRegisterExtModuleFunction
func XsltRegisterExtModuleFunction(name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar, function libxml_2_0.XmlXPathFunction) c.Int
//go:linkname XsltExtModuleFunctionLookup C.xsltExtModuleFunctionLookup
func XsltExtModuleFunctionLookup(name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar) libxml_2_0.XmlXPathFunction
//go:linkname XsltUnregisterExtModuleFunction C.xsltUnregisterExtModuleFunction
func XsltUnregisterExtModuleFunction(name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar) c.Int
// llgo:type C
type XsltPreComputeFunction func(XsltStylesheetPtr, libxml_2_0.XmlNodePtr, XsltTransformFunction) XsltElemPreCompPtr
//go:linkname XsltNewElemPreComp C.xsltNewElemPreComp
func XsltNewElemPreComp(style XsltStylesheetPtr, inst libxml_2_0.XmlNodePtr, function XsltTransformFunction) XsltElemPreCompPtr
//go:linkname XsltInitElemPreComp C.xsltInitElemPreComp
func XsltInitElemPreComp(comp XsltElemPreCompPtr, style XsltStylesheetPtr, inst libxml_2_0.XmlNodePtr, function XsltTransformFunction, freeFunc XsltElemPreCompDeallocator)
//go:linkname XsltRegisterExtModuleElement C.xsltRegisterExtModuleElement
func XsltRegisterExtModuleElement(name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar, precomp XsltPreComputeFunction, transform XsltTransformFunction) c.Int
//go:linkname XsltExtElementLookup C.xsltExtElementLookup
func XsltExtElementLookup(ctxt XsltTransformContextPtr, name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar) XsltTransformFunction
//go:linkname XsltExtModuleElementLookup C.xsltExtModuleElementLookup
func XsltExtModuleElementLookup(name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar) XsltTransformFunction
//go:linkname XsltExtModuleElementPreComputeLookup C.xsltExtModuleElementPreComputeLookup
func XsltExtModuleElementPreComputeLookup(name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar) XsltPreComputeFunction
//go:linkname XsltUnregisterExtModuleElement C.xsltUnregisterExtModuleElement
func XsltUnregisterExtModuleElement(name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar) c.Int
// llgo:type C
type XsltTopLevelFunction func(XsltStylesheetPtr, libxml_2_0.XmlNodePtr)
//go:linkname XsltRegisterExtModuleTopLevel C.xsltRegisterExtModuleTopLevel
func XsltRegisterExtModuleTopLevel(name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar, function XsltTopLevelFunction) c.Int
//go:linkname XsltExtModuleTopLevelLookup C.xsltExtModuleTopLevelLookup
func XsltExtModuleTopLevelLookup(name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar) XsltTopLevelFunction
//go:linkname XsltUnregisterExtModuleTopLevel C.xsltUnregisterExtModuleTopLevel
func XsltUnregisterExtModuleTopLevel(name *libxml_2_0.XmlChar, URI *libxml_2_0.XmlChar) c.Int
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
