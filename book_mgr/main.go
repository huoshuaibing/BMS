package main
import (
	"fmt"
	"os"
)
//使用函数实现一个简单的图书管理系统
//每本书有书名、作者、价格、上架信息
//用户可以在控制台添加书籍、修改书籍信息、打印所有书籍列表

//需求分析
//0. 定义结构体
type book struct{
	title string
	author string
	price float32
	publish bool
}
//1. 打印菜单
func showmenu(){
	fmt.Println("欢迎登陆BMS!")
	fmt.Println("1.添加书籍")
	fmt.Println("2.修改书籍")
	fmt.Println("3.展示所有书籍")
	fmt.Println("4.退出")
}

func userInput() *book {
	var (
		title string
		author string
		price float32
		publish bool
		)
	fmt.Println("请根据提示输入相关内容")
	fmt.Print("请输入书名:")
	fmt.Scanln(&title)
	fmt.Print("请输入作者:")
	fmt.Scanln(&author)
	fmt.Print("请输入价格:")
	fmt.Scanln(&price)
	fmt.Print("请输入是否上架(true|false):")
	fmt.Scanln(&publish)
	fmt.Println(title,author,price,publish)
	book := newbook(title,author,price,publish)
	return book
}
//2. 等待用户输入菜单选项
//定义一个book指针的切片，用来存储所有书籍
var allbooks = make([]*book,0,200)

//定义一个创建新书的构造函数
func newbook(title,author string, price float32, publish bool) *book{
	return &book{
		title: title,
		author: author,
		price: price,
		publish: publish,
	}
}
//3. 添加书籍的函数
func addbook(){
	var (
		title string
		author string
		price float32
		publish bool
	)
	fmt.Println("请根据提示输入相关内容")
	fmt.Print("请输入书名:")
	fmt.Scanln(&title)
	fmt.Print("请输入作者:")
	fmt.Scanln(&author)
	fmt.Print("请输入价格:")
	fmt.Scanln(&price)
	fmt.Print("请输入是否上架(true|false):")
	fmt.Scanln(&publish)
	fmt.Println(title,author,price,publish)
	book := newbook(title,author,price,publish)
	for _, b := range allbooks{
		if b.title == book.title{
			fmt.Printf("《%s》这本书已经存在",book.title)
			return
		}
	}
	allbooks = append(allbooks,book)
	fmt.Println("添加书籍成功！")
}
//4. 修改书籍的函数
func updatebook(){
	book := userInput()
	for index, b := range allbooks{
		if b.title == book.title{
			allbooks[index] = book
			fmt.Printf("书名:《%s》更新成功！",book.title)
			return 
		}
	}
	fmt.Printf("书名:《%s》不存在！", book.title )
}
//5. 展示书籍的函数
func showbook(){
	if len(allbooks) == 0 {
		fmt.Println("啥也么有")
	}
	for _, b := range allbooks {
		fmt.Printf("《%s》作者：%s 价格:%.2f 是否上架销售: %t\n",b.title,b.author,b.price,b.publish)
	}
}
//6. 退出 os.Exit(0)


func main(){
	for {
	    showmenu()
	    var option int
	    fmt.Scanln(&option)
	    switch option {
	    case 1:
		addbook()
	    case 2:
	    	updatebook()
	    case 3:
	    	showbook()
	    case 4:
	    	os.Exit(0)
		}
	}
}
