package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/llcppg-libxml"
	_ "unsafe"
)

type X_XsltCompMatch struct {
	Unused [8]uint8
}

type X_XsltNumberData struct {
	Level                *libxml_2_0.XmlChar
	Count                *libxml_2_0.XmlChar
	From                 *libxml_2_0.XmlChar
	Value                *libxml_2_0.XmlChar
	Format               *libxml_2_0.XmlChar
	HasFormat            c.Int
	DigitsPerGroup       c.Int
	GroupingCharacter    c.Int
	GroupingCharacterLen c.Int
	Doc                  libxml_2_0.XmlDocPtr
	Node                 libxml_2_0.XmlNodePtr
	CountPat             *X_XsltCompMatch
	FromPat              *X_XsltCompMatch
}
type XsltNumberData X_XsltNumberData
type XsltNumberDataPtr *XsltNumberData

type X_XsltFormatNumberInfo struct {
	IntegerHash       c.Int
	IntegerDigits     c.Int
	FracDigits        c.Int
	FracHash          c.Int
	Group             c.Int
	Multiplier        c.Int
	AddDecimal        int8
	IsMultiplierSet   int8
	IsNegativePattern int8
}
type XsltFormatNumberInfo X_XsltFormatNumberInfo
type XsltFormatNumberInfoPtr *XsltFormatNumberInfo
