package practice

// "golang.org/x/tour/pic"

// unit8:8bitで数の範囲を示す(数の出力は対象の数255で割った剰余-1を出す)
// 数の範囲 0-255
// 例: 256 => 256%255 = 0

func Pic(dx, dy int) [][]uint8 {
	// 配列は256*256の二次元配列(dx,dyはpicメソッド内で使用している)
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			//　配列に任意の数を入れる
			pic[y][x] = uint8((x + y))
		}
	}
	//fmt.Printf("%d,%d\n",dx,dy)
	return pic
}

// exmple
// func main() {
// 	pic.Show(Pic)
// }
