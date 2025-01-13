package libxslt

import _ "unsafe"

const XSLTDEFAULTVERSION string = "1.0"
const XSLTDEFAULTVENDOR string = "libxslt"
const XSLTDEFAULTURL string = "http://xmlsoft.org/XSLT/"
//go:linkname XsltInit C.xsltInit
func XsltInit()
//go:linkname XsltCleanupGlobals C.xsltCleanupGlobals
func XsltCleanupGlobals()
