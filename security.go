package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	_ "unsafe"
)

type X_XsltSecurityPrefs struct {
	Unused [8]uint8
}
type XsltSecurityPrefs X_XsltSecurityPrefs
type XsltSecurityPrefsPtr *XsltSecurityPrefs
type XsltSecurityOption c.Int

const (
	XsltSecurityOptionXSLTSECPREFREADFILE        XsltSecurityOption = 1
	XsltSecurityOptionXSLTSECPREFWRITEFILE       XsltSecurityOption = 2
	XsltSecurityOptionXSLTSECPREFCREATEDIRECTORY XsltSecurityOption = 3
	XsltSecurityOptionXSLTSECPREFREADNETWORK     XsltSecurityOption = 4
	XsltSecurityOptionXSLTSECPREFWRITENETWORK    XsltSecurityOption = 5
)
// llgo:type C
type XsltSecurityCheck func(XsltSecurityPrefsPtr, XsltTransformContextPtr, *int8) c.Int
//go:linkname XsltNewSecurityPrefs C.xsltNewSecurityPrefs
func XsltNewSecurityPrefs() XsltSecurityPrefsPtr
//go:linkname XsltFreeSecurityPrefs C.xsltFreeSecurityPrefs
func XsltFreeSecurityPrefs(sec XsltSecurityPrefsPtr)
//go:linkname XsltSetSecurityPrefs C.xsltSetSecurityPrefs
func XsltSetSecurityPrefs(sec XsltSecurityPrefsPtr, option XsltSecurityOption, func_ XsltSecurityCheck) c.Int
//go:linkname XsltGetSecurityPrefs C.xsltGetSecurityPrefs
func XsltGetSecurityPrefs(sec XsltSecurityPrefsPtr, option XsltSecurityOption) XsltSecurityCheck
//go:linkname XsltSetDefaultSecurityPrefs C.xsltSetDefaultSecurityPrefs
func XsltSetDefaultSecurityPrefs(sec XsltSecurityPrefsPtr)
//go:linkname XsltGetDefaultSecurityPrefs C.xsltGetDefaultSecurityPrefs
func XsltGetDefaultSecurityPrefs() XsltSecurityPrefsPtr
//go:linkname XsltSetCtxtSecurityPrefs C.xsltSetCtxtSecurityPrefs
func XsltSetCtxtSecurityPrefs(sec XsltSecurityPrefsPtr, ctxt XsltTransformContextPtr) c.Int
//go:linkname XsltSecurityAllow C.xsltSecurityAllow
func XsltSecurityAllow(sec XsltSecurityPrefsPtr, ctxt XsltTransformContextPtr, value *int8) c.Int
//go:linkname XsltSecurityForbid C.xsltSecurityForbid
func XsltSecurityForbid(sec XsltSecurityPrefsPtr, ctxt XsltTransformContextPtr, value *int8) c.Int
//go:linkname XsltCheckWrite C.xsltCheckWrite
func XsltCheckWrite(sec XsltSecurityPrefsPtr, ctxt XsltTransformContextPtr, URL *libxml_2_0.XmlChar) c.Int
//go:linkname XsltCheckRead C.xsltCheckRead
func XsltCheckRead(sec XsltSecurityPrefsPtr, ctxt XsltTransformContextPtr, URL *libxml_2_0.XmlChar) c.Int
