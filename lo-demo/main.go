package main


import (
	"fmt"
	"strconv"

	"github.com/samber/lo"
	// lop "github.com/samber/lo/parallel"
)

func main() {
	names := lo.Uniq[string]([]string{"Samuel", "Marc", "Samuel"}) // 类型参数会自动推断，可省略
	fmt.Println(names)

	res := lo.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})

	lo.ForEach[string](res, func(x string, _ int) {
		fmt.Print(x, " ")
	})
}
