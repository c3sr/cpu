// +build ppc64le

package cpuid

func SupportsAVX() bool {
	return false
}

func SupportsAVX2() bool {
	return false
}

func SupportsAVX512() bool {
	return false
}
