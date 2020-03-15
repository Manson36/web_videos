package main

import (
	"fmt"
	"testing"
)

//测试文件文件名必须是xxx_test.go；使用go test 命令或 go test -v 命令进行测试
/*
Test的写法：
1.每个test文件需import一个testing
2.test文件下的每一个test case 均必须以Test开头并且符合TestXXX形式，否则go test会直接跳过测试不执行
3.test case 的入参为 t *testing.T 或 b *testing.B
4.t.Errorf 为打印错误信息，并且当前的test case会跳过
5.t.SkipNow()为直接跳过当前的test，并且直接按照PASS处理继续下一个Test。必须写在test的第一行才能发挥作用
 */
func TestPrint(t *testing.T) {
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("Wrong result of print1to20")
	}
}

/*Test注意要点：
1.Go的Test不会保证多个test的顺序执行，但是通常是按顺序执行。
当我们需要按顺序执行或一个函数是使用另一个函数的结果的时候，解决方法：
使用t.Run来执行subtests可以做到控制test输出以及test的顺序

2.使用TestMain作为初始化test，并且使用m.Run()来调用其他tests可以完成一些需要初始化的testing，
	比如数据库的连接，文件打开，REST服务登陆等
	如果我们在testMain中没有执行m.Run(),那么出了Test Main其他的Test都不会执行
*/
func TestPrint1(t *testing.T) {
	t.Run("a1", func(t *testing.T) {fmt.Println("a1")})
	t.Run("a2", func(t *testing.T) {fmt.Println("a2")})
	t.Run("a3", func(t *testing.T) {fmt.Println("a3")})
}

func TestPrint2(t *testing.T) {
	res := Print1to20()
	res ++
	if res != 210 {
		fmt.Println("Test print2 failed")
	}
}
func TestMain(m *testing.M) {
	fmt.Println("test main first")
	m.Run()
}

func TestAll(t *testing.T) {
	t.Run("TestPrint", TestPrint)//函数使用大写会执行两遍，所以我们通常使用小写
	t.Run("TestPrint2", TestPrint2)
}

/*
Test之Benchmark：
1.benchmark一般以Benchmark开头
2.benchmark的case一般会跑b.N次，而且每次执行都会如此
3.在执行过程中会根据实际case的执行时间是否稳定怎加b.N的次数，以达到稳态:所以这里要非常注意一点，一定不要让Runtime处于非稳态
 */
//go test -bench=. 命令执行
 func BenchmarkAll(b *testing.B) {
 	for n := 0; n < b.N; n++ {
 		Print1to20()
 		//aaa(n)
	}
 }

 func aaa(n int) int {
 	for n > 0 {
 		n--
	}
 	return n
 }
