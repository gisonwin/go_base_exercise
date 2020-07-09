package main

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2019/8/19 16:45
 */
/***
对XSS最佳防护有两种方法:
1.验证所有输入数据,有效检测攻击
2.对所有输出数据进行适当处理,以防止任何已成功注入
的脚本在浏览器端运行.
Go里如何须知跨站脚本防护呢
Go的html/template里有以下向个函数可以帮助你
func HTMLEscape(w io.Writer,b []byte)//把b进行转义后写到w
func HTMLEscapeString(s string) string//转义s之后返回结果字符串
func HTMLEscaper(args ...interface{}) string//支持多个参数一起转义,返回结果字符串

 */
func anti_xss(){

}
/***
防止表单多次提交:
在表单中添加一个带有唯一值的隐藏字段.在验证表单时,先检查带有该唯一值的表单是否已经提交过了.
如果是,拒绝再次提交,如果不是,则按表单进行逻辑处理.
如果我们使用Ajax模式表单提交,当表单提交后,通过js来禁用表单的提交按钮.
 */
func main() {

}