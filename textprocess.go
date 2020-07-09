package main

import (
	"fmt"
	"os"
)

/**
 文本处理示例.
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2019/8/20 9:17
 */
func xmlprocess(){

}
/**
func Contains(s,substr string) bool //字符串s中是否包含substr,返回bool值.
func Join(a []string,sep string) string//字符串链接,把slice a通过sep链接起来
func Index(s,sep string) int //在字符串s中查找sep所在位置,返回位置值,不存在返回-1
func Repeat(s string,count int) string //重复s字符串count次,然后返回重复的字符串
func Replace(s,old,new string,n int) string//在s字符串中把old字符串替换为new字符串,n表示替换次数,小于0表示全部替换
func Split(s,sep string) []string //把s字符串按照sep分割,返回slice
func Trim(s string,cutset string) string //在s字符串中去除cutset指定的字符串
func Fields(s string) []string //去除s字符串的空格符,并按空格分割返回slice
字符串转换在strconv中
Append
Format
Parse
*/
func stringprocess(){

}
/***
文件操作大多函数都是在os包里面,
func Mkdir(name string,perm FileMode) error
func MkdirAll(path string,perm FileMode) error
func Remove(name string) error
func RemoveAll(path string) error

建立文件
func Create(name string) (file *File,err Error) //根据提供的文件名创建新的文件,返回一个文件对象,默认权限是0666,返回对象是可读写的.
func NewFile(fd uintptr,name string) *File //根据文件描述符创建相应的文件,返回一个文件对象
打开文件.
func Open(name string) (file *File,err Error)//打开一个名称为name的文件,是以只读方式,内部实现调用了OpenFile.
func OpenFile(name string,flag int,perm uint32) (file *File,err Error)
打开名称为name的文件,flag是打开方式,只读,读写等,perm是权限
写文件:
func (file *File) Write(b []byte) (n int,err Error) //写入byte类型的信息到文件.
func (file *File) WriteAt(b []byte,off int64) (n int,err Error)//从指定位置开始写入byte类型的信息
func (file *File) WriteString(s string) (ret int,err Error)//写string信息到文件
读文件
func (file *File) Read(b []byte) (n int,err Error)//读取数据到b中
func (file *File) ReadAt(b []byte,off int64) (n int,err Error)//从off开始读取数据到b中
删除文件
删除文件和删除文件夹是同一个函数
func Remove(name string) Error
 */
func fileprocess(){
	dirname := "gisonwin"
	os.Mkdir(dirname,0777)
	os.MkdirAll("gisonwin/test1/test2",0777)

	err := os.Remove(dirname)
	if err != nil {
		fmt.Println("error occurs: ",err)
	}
	os.RemoveAll(dirname)
}

func main() {
	fileprocess()
}