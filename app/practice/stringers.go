package practice

// 数値⇆文字列変換のパッケージをインポート
import (
	"fmt"
	"strconv"
)

// Stringメソッドの抽象化
type Stringer interface {
	String() string
}

// 型定義
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	var ipaddrString string

	// IPアドレスを文字列化する
	for num, octet := range ip {
		ipaddrString += strconv.Itoa(int(octet))
		if num < 3 {
			ipaddrString += "."
		}
	}
	return ipaddrString
}

func Action() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		// Stringerインターフェイスに型を入れ、内部のメソッドを具体化する
		var ipInterface Stringer = ip
		fmt.Printf("%v:%v\n", name, ipInterface.String())
	}
}
