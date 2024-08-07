package e

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func ExampleWriteFile() {
	err := X写到文件("./1.txt", []byte("123"))
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleReadFile() {
	内容 := X读入文件("./1.txt")
	fmt.Print(内容)
}

func ExampleRemove() {
	err := X删除文件("./1.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func TestEWriteFile(t *testing.T) {
	err := X写到文件("./1.txt", []byte("123"))
	// 抛出错误
	if err != nil {
		t.Errorf("文件读写删() = %v, want %v", err, nil)
	}
	内容 := X读入文件("./1.txt")
	err = X删除文件("./1.txt")
	// 抛出错误
	if err != nil {
		t.Errorf("文件读写删() = %v, want %v", err, nil)
	}
	// 断言 期望值和实际值是否相等
	if !reflect.DeepEqual(内容, []byte("123")) {
		t.Errorf("文件读写删() = %v, want %v", 内容, []byte("123"))
	}
}
func ExampleGetwd() {
	当前目录 := X取当前目录()
	fmt.Print(当前目录)
}
func TestGetwd(t *testing.T) {
	当前目录 := X取当前目录()
	if 当前目录 == "" {
		t.Errorf("X取当前目录() = %v, want %v", 当前目录, "不为空")
	}
}

func ExampleMkdir() {
	欲创建的目录名称 := "./test"
	err := X创建目录(欲创建的目录名称)
	if err != nil {
		log.Fatal(err)
	}
}
func TestMkdir(t *testing.T) {
	欲创建的目录名称 := "./test"
	err := X创建目录(欲创建的目录名称)
	if err != nil {
		t.Errorf("X创建目录() = %v, want %v", err, nil)
	}
	X删除目录(欲创建的目录名称)

}

func ExampleRemoveAll() {
	欲删除的目录名称 := "./test"
	err := X删除目录(欲删除的目录名称)
	if err != nil {
		log.Fatal(err)
	}
}
func TestRemoveAll(t *testing.T) {
	欲删除的目录名称 := "./test"
	err := X删除目录(欲删除的目录名称)
	if err != nil {
		t.Errorf("X删除目录() = %v, want %v", err, nil)
	}
}

func ExampleCopy() {
	被复制的文件名 := "./1.txt"
	复制到的文件名 := "./2.txt"
	err := X复制文件(被复制的文件名, 复制到的文件名)
	if err != nil {
		log.Fatal(err)
	}
}
func TestCopy(t *testing.T) {
	被复制的文件名 := "./1.txt"
	复制到的文件名 := "./2.txt"
	X写到文件("./1.txt", []byte("123"))

	err := X复制文件(被复制的文件名, 复制到的文件名)
	if err != nil {
		t.Errorf("X复制文件() = %v, want %v", err, nil)
	}
	内容 := X读入文件("./2.txt")
	err = X删除文件("./1.txt")
	err = X删除文件("./2.txt")
	if !reflect.DeepEqual(内容, []byte("123")) {
		t.Errorf("X复制文件() = %v, want %v", 内容, []byte("123"))
	}
}

func ExampleCopyAll() {
	被复制的目录名 := "./test"
	复制到的目录名 := "./test2"
	err := X复制目录(被复制的目录名, 复制到的目录名)
	if err != nil {
		log.Fatal(err)
	}

}

func TestCopyAll(t *testing.T) {
	被复制的目录名 := "./test"
	复制到的目录名 := "./test2"
	X创建目录(被复制的目录名)
	X写到文件(被复制的目录名+"/1.txt", []byte("123"))

	err := X复制目录(被复制的目录名, 复制到的目录名)
	if err != nil {
		log.Fatal(err)
	}
	内容 := X读入文件(复制到的目录名 + "/1.txt")
	err = X删除目录(被复制的目录名)
	err = X删除目录(复制到的目录名)

	if !reflect.DeepEqual(内容, []byte("123")) {
		t.Errorf("X复制文件() = %v, want %v", 内容, []byte("123"))
	}
}

func ExampleRename() {
	被移动的文件名 := "./1.txt"
	移动到的文件名 := "./2.txt"
	err := X移动文件(被移动的文件名, 移动到的文件名)
	if err != nil {
		log.Fatal(err)
	}
}

func TestRename(t *testing.T) {
	被移动的文件名 := "./1.txt"
	移动到的文件名 := "./2.txt"
	X写到文件("./1.txt", []byte("123"))

	err := X移动文件(被移动的文件名, 移动到的文件名)
	if err != nil {
		t.Errorf("X移动文件() = %v, want %v", err, nil)
	}
	内容 := X读入文件("./2.txt")
	err = X删除文件("./2.txt")
	if !reflect.DeepEqual(内容, []byte("123")) {
		t.Errorf("X移动文件() = %v, want %v", 内容, []byte("123"))
	}

}

func ExampleRename2() {
	被更名的文件名 := "./1.txt"
	更名到的文件名 := "./2.txt"
	err := X文件更名(被更名的文件名, 更名到的文件名)
	if err != nil {
		log.Fatal(err)
	}
}

func TestRename2(t *testing.T) {
	被更名的文件名 := "./1.txt"
	更名到的文件名 := "./2.txt"
	X写到文件("./1.txt", []byte("123"))

	err := X文件更名(被更名的文件名, 更名到的文件名)
	if err != nil {
		t.Errorf("X文件更名() = %v, want %v", err, nil)
	}
	内容 := X读入文件("./2.txt")
	err = X删除文件("./2.txt")
	if !reflect.DeepEqual(内容, []byte("123")) {
		t.Errorf("X文件更名() = %v, want %v", 内容, []byte("123"))
	}
}

// X文件是否存在
func ExampleIsNotExist() {
	文件名 := "./1.txt"
	文件是否存在 := X文件是否存在(文件名)
	fmt.Println(文件是否存在)
	// Output:
	// false
}

func TestIsNotExist(t *testing.T) {
	文件名 := "./1.txt"
	文件是否存在 := X文件是否存在(文件名)
	if 文件是否存在 {
		t.Errorf("X文件是否存在() = %v, want %v", 文件是否存在, false)
	}

}

// X取文件尺寸
func ExampleSize() {
	文件名 := "./1.txt"
	文件尺寸 := X取文件尺寸(文件名)
	fmt.Println(文件尺寸)
}

func TestSize(t *testing.T) {
	文件名 := "./1.txt"
	X写到文件(文件名, []byte("123"))
	文件尺寸 := X取文件尺寸(文件名)
	if 文件尺寸 != 3 {
		t.Errorf("X取文件尺寸() = %v, want %v", 文件尺寸, 0)
	}
	X删除文件(文件名)
}

// X取临时文件名
func ExampleTempFile() {
	_, 临时文件名, _ := X取临时文件名("")
	fmt.Println(临时文件名)
}

func TestTempFile(t *testing.T) {
	_, 临时文件名, err := X取临时文件名("")
	if err != nil {
		t.Errorf("X取临时文件名() = %v, want %v", err, false)
	}
	X删除文件(临时文件名)
}

func TestListFiles(t *testing.T) {
	var files []string
	_ = X文件枚举("./", ".go", &files, false, true)
	for i, value := range files {
		fmt.Println(i, value)
	}
}

func TestEListSubdirectories(t *testing.T) {
	var files []string
	_ = X目录枚举子目录("./", &files, true, true)
	for i, value := range files {
		fmt.Println(i, value)
	}
}

func TestFileEnumeration(t *testing.T) {
	var files []string
	_ = X文件枚举("./", ".go", &files, true, true)
	for i, value := range files {
		fmt.Println(i, value, X文件取文件名(value, true))
	}
}

func TestDir(t *testing.T) {
	fmt.Println(" :", X文件取父目录(""))
	fmt.Println(". :", X文件取父目录("."))
	fmt.Println("/X父类 :", X文件取父目录("/X父类"))
	fmt.Println("/X父类/b :", X文件取父目录("/X父类/b"))
	fmt.Println("/X父类/b/ :", X文件取父目录("/X父类/b/"))
	fmt.Println("/////X父类, /b :", X文件路径合并处理("X父类", "b", "c"))
	fmt.Println("/////X父类, /b :", X文件路径合并处理("X父类/11////", "b", "c"))
}

func TestWriteFile(t *testing.T) {
	X调试输出(X取运行目录())
	X文件写出(X取运行目录()+"/aaa/ccc/ddd/1.txt", X到字节集("123"))
}

func TestEAddText(t *testing.T) {
	X调试输出(X取运行目录())
	X文件追加文本(X取运行目录()+"/aaa/ccc/ddd/2.txt", "123")
}

func TestFileSave(t *testing.T) {
	X调试输出(X取运行目录())
	X文件保存(X取运行目录()+"/aaa/ccc/ddd/2.txt", "1234")
}
