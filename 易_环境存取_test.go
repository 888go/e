package e

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestE写环境变量(t *testing.T) {
	type args struct {
		环境变量名称 string
		欲写入内容  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test",
			args{
				"test",
				"123",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := X写环境变量(tt.args.环境变量名称, tt.args.欲写入内容); got != tt.want {
				t.Errorf("X写环境变量() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestE取命令行(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := X取命令行(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("X取命令行() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestE取运行目录(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := X取运行目录(); got != tt.want {
				t.Errorf("X取运行目录() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestE读环境变量(t *testing.T) {
	X写环境变量("test", "123")
	type args struct {
		环境变量名称 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test",
			args{
				"test",
			},
			"123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := X读环境变量(tt.args.环境变量名称); got != tt.want {
				t.Errorf("X读环境变量() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleE取命令行() {
	os.Args = []string{"example", "foo", "bar", "baz"}
	args := X取命令行()
	fmt.Println(args)
	// Output: [example foo bar baz]
}
func ExampleE取运行目录() {
	path := X取运行目录()
	fmt.Println(path)
}

func ExampleE读环境变量() {
	os.Setenv("MY_ENV_VAR", "my_value")
	value := X读环境变量("MY_ENV_VAR")
	fmt.Println(value)
	//  Output:  my_value
}

func ExampleE写环境变量() {
	result := X写环境变量("MY_ENV_VAR", "my_new_value")
	fmt.Println(result)
	//  Output:  true
}

func TestE设置命令行(t *testing.T) {
	X调试输出(X取命令行()[1])
	X调试输出(X取命令行()[2])
	X调试输出(X取命令行()[3])

	var a, b, c string
	X设置命令行("X父类", "X父类", "这是一个a参数", &a)
	X设置命令行("b", "b", "这是一个b参数", &b)
	X设置命令行("c", "c", "这是一个c参数", &c)
	X命令行解析()

	X调试输出(a, b, c)
}

func TestE加载环境变量文件(t *testing.T) {
	X加载环境变量_从文件("test.env")
	t.Log("SECRET_KEY", X读环境变量("S3_BUCKET"))
	t.Log("SECRET_KEY", X读环境变量("SECRET_KEY"))
	t.Log("default", X读环境变量("default", ""))
	t.Log("default", X读环境变量("default", "123"))

	env, _ := X加载环境变量_从内容到map(X到文本(X读入文件("test.env")))
	for k, v := range env {
		t.Log("env", k, v)
	}

}
