package types

const (
	// ModuleName defines the module name
	ModuleName = "match"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_match"
)

var (
	ParamsKey = []byte("p_match")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
