package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	_ "unsafe"
)
/**
 * XInclude default processing.
 */
//go:linkname XsltSetXIncludeDefault C.xsltSetXIncludeDefault
func XsltSetXIncludeDefault(xinclude c.Int)
//go:linkname XsltGetXIncludeDefault C.xsltGetXIncludeDefault
func XsltGetXIncludeDefault() c.Int
/**
 * Export context to users.
 */
//go:linkname XsltNewTransformContext C.xsltNewTransformContext
func XsltNewTransformContext(style XsltStylesheetPtr, doc libxml_2_0.XmlDocPtr) XsltTransformContextPtr
//go:linkname XsltFreeTransformContext C.xsltFreeTransformContext
func XsltFreeTransformContext(ctxt XsltTransformContextPtr)
//go:linkname XsltApplyStylesheetUser C.xsltApplyStylesheetUser
func XsltApplyStylesheetUser(style XsltStylesheetPtr, doc libxml_2_0.XmlDocPtr, params **int8, output *int8, profile *c.FILE, userCtxt XsltTransformContextPtr) libxml_2_0.XmlDocPtr
//go:linkname XsltProcessOneNode C.xsltProcessOneNode
func XsltProcessOneNode(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, params XsltStackElemPtr)
/**
 * Private Interfaces.
 */
//go:linkname XsltApplyStripSpaces C.xsltApplyStripSpaces
func XsltApplyStripSpaces(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr)
//go:linkname XsltApplyStylesheet C.xsltApplyStylesheet
func XsltApplyStylesheet(style XsltStylesheetPtr, doc libxml_2_0.XmlDocPtr, params **int8) libxml_2_0.XmlDocPtr
//go:linkname XsltProfileStylesheet C.xsltProfileStylesheet
func XsltProfileStylesheet(style XsltStylesheetPtr, doc libxml_2_0.XmlDocPtr, params **int8, output *c.FILE) libxml_2_0.XmlDocPtr
//go:linkname XsltRunStylesheet C.xsltRunStylesheet
func XsltRunStylesheet(style XsltStylesheetPtr, doc libxml_2_0.XmlDocPtr, params **int8, output *int8, SAX libxml_2_0.XmlSAXHandlerPtr, IObuf libxml_2_0.XmlOutputBufferPtr) c.Int
//go:linkname XsltRunStylesheetUser C.xsltRunStylesheetUser
func XsltRunStylesheetUser(style XsltStylesheetPtr, doc libxml_2_0.XmlDocPtr, params **int8, output *int8, SAX libxml_2_0.XmlSAXHandlerPtr, IObuf libxml_2_0.XmlOutputBufferPtr, profile *c.FILE, userCtxt XsltTransformContextPtr) c.Int
//go:linkname XsltApplyOneTemplate C.xsltApplyOneTemplate
func XsltApplyOneTemplate(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, list libxml_2_0.XmlNodePtr, templ XsltTemplatePtr, params XsltStackElemPtr)
//go:linkname XsltDocumentElem C.xsltDocumentElem
func XsltDocumentElem(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltSort C.xsltSort
func XsltSort(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltCopy C.xsltCopy
func XsltCopy(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltText C.xsltText
func XsltText(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltElement C.xsltElement
func XsltElement(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltComment C.xsltComment
func XsltComment(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltAttribute C.xsltAttribute
func XsltAttribute(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltProcessingInstruction C.xsltProcessingInstruction
func XsltProcessingInstruction(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltCopyOf C.xsltCopyOf
func XsltCopyOf(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltValueOf C.xsltValueOf
func XsltValueOf(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltNumber C.xsltNumber
func XsltNumber(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltApplyImports C.xsltApplyImports
func XsltApplyImports(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltCallTemplate C.xsltCallTemplate
func XsltCallTemplate(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltApplyTemplates C.xsltApplyTemplates
func XsltApplyTemplates(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltChoose C.xsltChoose
func XsltChoose(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltIf C.xsltIf
func XsltIf(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltForEach C.xsltForEach
func XsltForEach(ctxt XsltTransformContextPtr, node libxml_2_0.XmlNodePtr, inst libxml_2_0.XmlNodePtr, comp XsltElemPreCompPtr)
//go:linkname XsltRegisterAllElement C.xsltRegisterAllElement
func XsltRegisterAllElement(ctxt XsltTransformContextPtr)
//go:linkname XsltCopyTextString C.xsltCopyTextString
func XsltCopyTextString(ctxt XsltTransformContextPtr, target libxml_2_0.XmlNodePtr, string *libxml_2_0.XmlChar, noescape c.Int) libxml_2_0.XmlNodePtr
//go:linkname XsltLocalVariablePop C.xsltLocalVariablePop
func XsltLocalVariablePop(ctxt XsltTransformContextPtr, limitNr c.Int, level c.Int)
//go:linkname XsltLocalVariablePush C.xsltLocalVariablePush
func XsltLocalVariablePush(ctxt XsltTransformContextPtr, variable XsltStackElemPtr, level c.Int) c.Int
//go:linkname XslHandleDebugger C.xslHandleDebugger
func XslHandleDebugger(cur libxml_2_0.XmlNodePtr, node libxml_2_0.XmlNodePtr, templ XsltTemplatePtr, ctxt XsltTransformContextPtr)
