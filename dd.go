package dd

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib
#include "dd.hpp"
*/
import "C"

func init() {
	ok := bool(C.init())
	if !ok {
		panic(`DD PACKAGE INITIALIZATION FAILED. 
		Please put the DD dll into the program directory.
		`)
	}
}

/*
功能： 模拟鼠标点击
参数： 1 =左键按下 ，2 =左键放开
4 =右键按下 ，8 =右键放开
16 =中键按下 ，32 =中键放开
64 =4键按下 ，128 =4键放开
256 =5键按下 ，512 =5键放开
*/
func MouseBtn(flag int32) int32{
	return int32(C.DD_btn(C.int32_t(flag)))
}

/*
功能： 模拟鼠标移动
参数： 参数x , 参数y 以屏幕左上角为原点。
例子： 把鼠标移动到分辨率1920*1080 的屏幕正中间，
int x = 1920/2 ; int y = 1080/2;
*/
func MouseMove(x int32, y int32) int32{
	return int32(C.DD_mov(C.int32_t(x), C.int32_t(y)))
}

/*
功能： 模拟鼠标相对移动
参数： 参数dx , 参数dy 以当前坐标为原点。
例子： 把鼠标向左移动10像素
*/
func MouseMoveRelative(x int32, y int32) int32{
	return int32(C.DD_movR(C.int32_t(x), C.int32_t(y)))
}

/*
功能: 模拟鼠标滚轮
参数: 1=前 , 2 = 后
*/
func MouseWheel(flag int32) int32{
	return int32(C.DD_whl(C.int32_t(flag)))
}

/*
功能： 模拟键盘按键
参数： 参数1 ，请查看[DD虚拟键盘码表]。
参数2，1=按下，2=放开
*/
func KeyboardKey(key int32, status int32) int32{
	return int32(C.DD_key(C.int32_t(key), C.int32_t(status)))
}

/*
功能： 模拟键盘按键，按下并放开
参数： 参数1 ，请查看[DD虚拟键盘码表]。
*/
func KeyboardPress(key int32) int32{
	if KeyboardKey(key, 1) == -1 {
		return -1
	}
	return KeyboardKey(key, 0)
}

/*
功能： 直接输入键盘上可见字符和空格
参数： 字符串, (注意，这个参数不是int32 类型)
*/
func KeyboardInput(str string) int32{
	return int32(C.DD_str(C.CString(str)))
}