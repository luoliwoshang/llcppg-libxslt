package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	_ "unsafe"
)
//go:linkname XsltAddKey C.xsltAddKey
func XsltAddKey(style XsltStylesheetPtr, name *libxml_2_0.XmlChar, nameURI *libxml_2_0.XmlChar, match *libxml_2_0.XmlChar, use *libxml_2_0.XmlChar, inst libxml_2_0.XmlNodePtr) c.Int
//go:linkname XsltGetKey C.xsltGetKey
func XsltGetKey(ctxt XsltTransformContextPtr, name *libxml_2_0.XmlChar, nameURI *libxml_2_0.XmlChar, value *libxml_2_0.XmlChar) libxml_2_0.XmlNodeSetPtr
//go:linkname XsltInitCtxtKeys C.xsltInitCtxtKeys
func XsltInitCtxtKeys(ctxt XsltTransformContextPtr, doc XsltDocumentPtr)
//go:linkname XsltFreeKeys C.xsltFreeKeys
func XsltFreeKeys(style XsltStylesheetPtr)
//go:linkname XsltFreeDocumentKeys C.xsltFreeDocumentKeys
func XsltFreeDocumentKeys(doc XsltDocumentPtr)
