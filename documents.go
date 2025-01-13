package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	"unsafe"
)
//go:linkname XsltNewDocument C.xsltNewDocument
func XsltNewDocument(ctxt XsltTransformContextPtr, doc libxml_2_0.XmlDocPtr) XsltDocumentPtr
//go:linkname XsltLoadDocument C.xsltLoadDocument
func XsltLoadDocument(ctxt XsltTransformContextPtr, URI *libxml_2_0.XmlChar) XsltDocumentPtr
//go:linkname XsltFindDocument C.xsltFindDocument
func XsltFindDocument(ctxt XsltTransformContextPtr, doc libxml_2_0.XmlDocPtr) XsltDocumentPtr
//go:linkname XsltFreeDocuments C.xsltFreeDocuments
func XsltFreeDocuments(ctxt XsltTransformContextPtr)
//go:linkname XsltLoadStyleDocument C.xsltLoadStyleDocument
func XsltLoadStyleDocument(style XsltStylesheetPtr, URI *libxml_2_0.XmlChar) XsltDocumentPtr
//go:linkname XsltNewStyleDocument C.xsltNewStyleDocument
func XsltNewStyleDocument(style XsltStylesheetPtr, doc libxml_2_0.XmlDocPtr) XsltDocumentPtr
//go:linkname XsltFreeStyleDocuments C.xsltFreeStyleDocuments
func XsltFreeStyleDocuments(style XsltStylesheetPtr)

type XsltLoadType c.Int

const (
	XsltLoadTypeXSLTLOADSTART      XsltLoadType = 0
	XsltLoadTypeXSLTLOADSTYLESHEET XsltLoadType = 1
	XsltLoadTypeXSLTLOADDOCUMENT   XsltLoadType = 2
)
// llgo:type C
type XsltDocLoaderFunc func(*libxml_2_0.XmlChar, libxml_2_0.XmlDictPtr, c.Int, unsafe.Pointer, XsltLoadType) libxml_2_0.XmlDocPtr
//go:linkname XsltSetLoaderFunc C.xsltSetLoaderFunc
func XsltSetLoaderFunc(f XsltDocLoaderFunc)
