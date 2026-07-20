package geo

type encoding struct {
	enc string
	dec [256]byte
}

func newEncoding(encoder string) *encoding { _ = "STUB: not implemented"; return nil }

var base32encoding = newEncoding("0123456789bcdefghjkmnpqrstuvwxyz")

var masks = []uint64{16, 8, 4, 2, 1}

func DecodeGeoHash(geoHash string) (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

func EncodeGeoHash(lat, lon float64) string { _ = "STUB: not implemented"; return "" }
