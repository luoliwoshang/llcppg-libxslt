package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	"unsafe"
)
// llgo:link (*XmlChar).XsltNewLocale C.xsltNewLocale
func (recv_ *libxml_2_0.XmlChar) XsltNewLocale(lowerFirst c.Int) unsafe.Pointer {
	return nil
}
//go:linkname XsltFreeLocale C.xsltFreeLocale
func XsltFreeLocale(locale unsafe.Pointer)
//go:linkname XsltStrxfrm C.xsltStrxfrm
func XsltStrxfrm(locale unsafe.Pointer, string *libxml_2_0.XmlChar) *libxml_2_0.XmlChar
//go:linkname XsltFreeLocales C.xsltFreeLocales
func XsltFreeLocales()

type XsltLocale unsafe.Pointer
type XsltLocaleChar libxml_2_0.XmlChar
//go:linkname XsltLocaleStrcmp C.xsltLocaleStrcmp
func XsltLocaleStrcmp(locale unsafe.Pointer, str1 *libxml_2_0.XmlChar, str2 *libxml_2_0.XmlChar) c.Int
