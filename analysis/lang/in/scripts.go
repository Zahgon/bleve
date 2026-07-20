package in

import (
	"unicode"

	"github.com/bits-and-blooms/bitset"
)

type ScriptData struct {
	flag       rune
	base       rune
	decompMask *bitset.BitSet
}

var scripts = map[*unicode.RangeTable]*ScriptData{
	unicode.Devanagari: {
		flag: 1,
		base: 0x0900,
	},
	unicode.Bengali: {
		flag: 2,
		base: 0x0980,
	},
	unicode.Gurmukhi: {
		flag: 4,
		base: 0x0A00,
	},
	unicode.Gujarati: {
		flag: 8,
		base: 0x0A80,
	},
	unicode.Oriya: {
		flag: 16,
		base: 0x0B00,
	},
	unicode.Tamil: {
		flag: 32,
		base: 0x0B80,
	},
	unicode.Telugu: {
		flag: 64,
		base: 0x0C00,
	},
	unicode.Kannada: {
		flag: 128,
		base: 0x0C80,
	},
	unicode.Malayalam: {
		flag: 256,
		base: 0x0D00,
	},
}

func flag(ub *unicode.RangeTable) rune { _ = "STUB: not implemented"; return 0 }

var decompositions = [][]rune{

	{0x05, 0x3E, 0x45, 0x11, flag(unicode.Devanagari) | flag(unicode.Gujarati)},

	{0x05, 0x3E, 0x46, 0x12, flag(unicode.Devanagari)},

	{0x05, 0x3E, 0x47, 0x13, flag(unicode.Devanagari) | flag(unicode.Gujarati)},

	{0x05, 0x3E, 0x48, 0x14, flag(unicode.Devanagari) | flag(unicode.Gujarati)},

	{0x05, 0x3E, -1, 0x06, flag(unicode.Devanagari) | flag(unicode.Bengali) | flag(unicode.Gurmukhi) | flag(unicode.Gujarati) | flag(unicode.Oriya)},

	{0x05, 0x45, -1, 0x72, flag(unicode.Devanagari)},

	{0x05, 0x45, -1, 0x0D, flag(unicode.Gujarati)},

	{0x05, 0x46, -1, 0x04, flag(unicode.Devanagari)},

	{0x05, 0x47, -1, 0x0F, flag(unicode.Gujarati)},

	{0x05, 0x48, -1, 0x10, flag(unicode.Gurmukhi) | flag(unicode.Gujarati)},

	{0x05, 0x49, -1, 0x11, flag(unicode.Devanagari) | flag(unicode.Gujarati)},

	{0x05, 0x4A, -1, 0x12, flag(unicode.Devanagari)},

	{0x05, 0x4B, -1, 0x13, flag(unicode.Devanagari) | flag(unicode.Gujarati)},

	{0x05, 0x4C, -1, 0x14, flag(unicode.Devanagari) | flag(unicode.Gurmukhi) | flag(unicode.Gujarati)},

	{0x06, 0x45, -1, 0x11, flag(unicode.Devanagari) | flag(unicode.Gujarati)},

	{0x06, 0x46, -1, 0x12, flag(unicode.Devanagari)},

	{0x06, 0x47, -1, 0x13, flag(unicode.Devanagari) | flag(unicode.Gujarati)},

	{0x06, 0x48, -1, 0x14, flag(unicode.Devanagari) | flag(unicode.Gujarati)},

	{0x07, 0x57, -1, 0x08, flag(unicode.Malayalam)},

	{0x09, 0x41, -1, 0x0A, flag(unicode.Devanagari)},

	{0x09, 0x57, -1, 0x0A, flag(unicode.Tamil) | flag(unicode.Malayalam)},

	{0x0E, 0x46, -1, 0x10, flag(unicode.Malayalam)},

	{0x0F, 0x45, -1, 0x0D, flag(unicode.Devanagari)},

	{0x0F, 0x46, -1, 0x0E, flag(unicode.Devanagari)},

	{0x0F, 0x47, -1, 0x10, flag(unicode.Devanagari)},

	{0x0F, 0x57, -1, 0x10, flag(unicode.Oriya)},

	{0x12, 0x3E, -1, 0x13, flag(unicode.Malayalam)},

	{0x12, 0x4C, -1, 0x14, flag(unicode.Telugu) | flag(unicode.Kannada)},

	{0x12, 0x55, -1, 0x13, flag(unicode.Telugu)},

	{0x12, 0x57, -1, 0x14, flag(unicode.Tamil) | flag(unicode.Malayalam)},

	{0x13, 0x57, -1, 0x14, flag(unicode.Oriya)},

	{0x15, 0x3C, -1, 0x58, flag(unicode.Devanagari)},

	{0x16, 0x3C, -1, 0x59, flag(unicode.Devanagari) | flag(unicode.Gurmukhi)},

	{0x17, 0x3C, -1, 0x5A, flag(unicode.Devanagari) | flag(unicode.Gurmukhi)},

	{0x1C, 0x3C, -1, 0x5B, flag(unicode.Devanagari) | flag(unicode.Gurmukhi)},

	{0x21, 0x3C, -1, 0x5C, flag(unicode.Devanagari) | flag(unicode.Bengali) | flag(unicode.Oriya)},

	{0x22, 0x3C, -1, 0x5D, flag(unicode.Devanagari) | flag(unicode.Bengali) | flag(unicode.Oriya)},

	{0x23, 0x4D, 0xFF, 0x7A, flag(unicode.Malayalam)},

	{0x24, 0x4D, 0xFF, 0x4E, flag(unicode.Bengali)},

	{0x28, 0x3C, -1, 0x29, flag(unicode.Devanagari)},

	{0x28, 0x4D, 0xFF, 0x7B, flag(unicode.Malayalam)},

	{0x2B, 0x3C, -1, 0x5E, flag(unicode.Devanagari) | flag(unicode.Gurmukhi)},

	{0x2F, 0x3C, -1, 0x5F, flag(unicode.Devanagari) | flag(unicode.Bengali)},

	{0x2C, 0x41, 0x41, 0x0B, flag(unicode.Telugu)},

	{0x30, 0x3C, -1, 0x31, flag(unicode.Devanagari)},

	{0x30, 0x4D, 0xFF, 0x7C, flag(unicode.Malayalam)},

	{0x32, 0x4D, 0xFF, 0x7D, flag(unicode.Malayalam)},

	{0x33, 0x3C, -1, 0x34, flag(unicode.Devanagari)},

	{0x33, 0x4D, 0xFF, 0x7E, flag(unicode.Malayalam)},

	{0x35, 0x41, -1, 0x2E, flag(unicode.Telugu)},

	{0x3E, 0x45, -1, 0x49, flag(unicode.Devanagari) | flag(unicode.Gujarati)},

	{0x3E, 0x46, -1, 0x4A, flag(unicode.Devanagari)},

	{0x3E, 0x47, -1, 0x4B, flag(unicode.Devanagari) | flag(unicode.Gujarati)},

	{0x3E, 0x48, -1, 0x4C, flag(unicode.Devanagari) | flag(unicode.Gujarati)},

	{0x3F, 0x55, -1, 0x40, flag(unicode.Kannada)},

	{0x41, 0x41, -1, 0x42, flag(unicode.Gurmukhi)},

	{0x46, 0x3E, -1, 0x4A, flag(unicode.Tamil) | flag(unicode.Malayalam)},

	{0x46, 0x42, 0x55, 0x4B, flag(unicode.Kannada)},

	{0x46, 0x42, -1, 0x4A, flag(unicode.Kannada)},

	{0x46, 0x46, -1, 0x48, flag(unicode.Malayalam)},

	{0x46, 0x55, -1, 0x47, flag(unicode.Telugu) | flag(unicode.Kannada)},

	{0x46, 0x56, -1, 0x48, flag(unicode.Telugu) | flag(unicode.Kannada)},

	{0x46, 0x57, -1, 0x4C, flag(unicode.Tamil) | flag(unicode.Malayalam)},

	{0x47, 0x3E, -1, 0x4B, flag(unicode.Bengali) | flag(unicode.Oriya) | flag(unicode.Tamil) | flag(unicode.Malayalam)},

	{0x47, 0x57, -1, 0x4C, flag(unicode.Bengali) | flag(unicode.Oriya)},

	{0x4A, 0x55, -1, 0x4B, flag(unicode.Kannada)},

	{0x72, 0x3F, -1, 0x07, flag(unicode.Gurmukhi)},

	{0x72, 0x40, -1, 0x08, flag(unicode.Gurmukhi)},

	{0x72, 0x47, -1, 0x0F, flag(unicode.Gurmukhi)},

	{0x73, 0x41, -1, 0x09, flag(unicode.Gurmukhi)},

	{0x73, 0x42, -1, 0x0A, flag(unicode.Gurmukhi)},

	{0x73, 0x4B, -1, 0x13, flag(unicode.Gurmukhi)},
}

func init() {
	for _, scriptData := range scripts {
		scriptData.decompMask = bitset.New(0x7d)
		for _, decomposition := range decompositions {
			ch := decomposition[0]
			flags := decomposition[4]
			if (flags & scriptData.flag) != 0 {
				scriptData.decompMask.Set(uint(ch))
			}
		}
	}
}

func lookupScript(r rune) *unicode.RangeTable { _ = "STUB: not implemented"; return nil }

func normalize(input []rune) []rune { _ = "STUB: not implemented"; return nil }

func compose(ch0 rune, script0 *unicode.RangeTable, scriptData *ScriptData, input []rune, pos int, inputLen int) []rune {
	_ = "STUB: not implemented"
	return nil
}
