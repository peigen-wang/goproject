package uniqueid

import "testing"

func TestGenSn(t *testing.T) {
	code := GenSn(SN_PREFIX_HOMESTAY_ORDER)
	println(code)
}

func TestHelloWorld(t *testing.T) {
	println("hello world")

}
