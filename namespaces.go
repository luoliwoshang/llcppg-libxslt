package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	_ "unsafe"
)
//go:linkname XsltNamespaceAlias C.xsltNamespaceAlias
func XsltNamespaceAlias(style c.Int, node libxml_2_0.XmlNodePtr)
//go:linkname XsltGetNamespace C.xsltGetNamespace
func XsltGetNamespace(ctxt c.Int, cur libxml_2_0.XmlNodePtr, ns libxml_2_0.XmlNsPtr, out libxml_2_0.XmlNodePtr) libxml_2_0.XmlNsPtr
//go:linkname XsltGetPlainNamespace C.xsltGetPlainNamespace
func XsltGetPlainNamespace(ctxt c.Int, cur libxml_2_0.XmlNodePtr, ns libxml_2_0.XmlNsPtr, out libxml_2_0.XmlNodePtr) libxml_2_0.XmlNsPtr
//go:linkname XsltGetSpecialNamespace C.xsltGetSpecialNamespace
func XsltGetSpecialNamespace(ctxt c.Int, cur libxml_2_0.XmlNodePtr, URI *libxml_2_0.XmlChar, prefix *libxml_2_0.XmlChar, out libxml_2_0.XmlNodePtr) libxml_2_0.XmlNsPtr
//go:linkname XsltCopyNamespace C.xsltCopyNamespace
func XsltCopyNamespace(ctxt c.Int, elem libxml_2_0.XmlNodePtr, ns libxml_2_0.XmlNsPtr) libxml_2_0.XmlNsPtr
//go:linkname XsltCopyNamespaceList C.xsltCopyNamespaceList
func XsltCopyNamespaceList(ctxt c.Int, node libxml_2_0.XmlNodePtr, cur libxml_2_0.XmlNsPtr) libxml_2_0.XmlNsPtr
//go:linkname XsltFreeNamespaceAliasHashes C.xsltFreeNamespaceAliasHashes
func XsltFreeNamespaceAliasHashes(style c.Int)
