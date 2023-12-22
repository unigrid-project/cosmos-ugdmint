package types

const (
	// ModuleName defines the module name
	ModuleName = "ugdmint"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ugdmint"
)

var (
	MinterKey = []byte{0x00}
	ParamsKey = []byte{0x01}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
