package numeric

var interleaveMagic = []uint64{
	0x5555555555555555,
	0x3333333333333333,
	0x0F0F0F0F0F0F0F0F,
	0x00FF00FF00FF00FF,
	0x0000FFFF0000FFFF,
	0x00000000FFFFFFFF,
	0xAAAAAAAAAAAAAAAA,
}

var interleaveShift = []uint{1, 2, 4, 8, 16}

func Interleave(v1, v2 uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func Deinterleave(b uint64) uint64 { _ = "STUB: not implemented"; return 0 }
