# expected 'package', found 'EOF'

**报错解释:这个错误通常出现在编程或脚本执行过程中，提示解析器期望在某个位置遇到一个 'package' 声明，
但是却遇到了文件的末尾（EOF, End Of File）。这通常发生在编写代码时忘记完成某个声明或语句，或者在处理配置文件时格式不正确。**

PS:遇到时的原因很奇怪，在一个主函数中写了很多程序段代码，然后在终端运行出错，但是重新移动package后程序又恢复正常

# package command-line-arguments is not a main package
**报错解释：
这个错误通常出现在尝试运行Go语言程序时。当你执行go run或go build命令时，如果Go编译器找不到包含main函数的包，它会报告这个错误。这意味着它没有找到声明为package main的主程序包。确保你的程序文件以package main开头**

PS：在调用不同包定义的函数时放在同一个文件中，其实应该在主目录创建子文件夹，在子文件夹写好程序，再调用
