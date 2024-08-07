package e

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type X配置项 struct {
	cfg     *ini.File
	节名区分大小写 bool
}

// X创建配置项 创建一个配置项
func X创建配置项(配置项数据 string, 是否区分大小写 bool) *X配置项 {
	o := new(X配置项)

	if 是否区分大小写 == false {
		o.X设置节名不区分大小写()
	}

	o.X加载配置项从内存(配置项数据)
	return o
}

// X设置节名不区分大小写
func (this *X配置项) X设置节名不区分大小写() {
	this.节名区分大小写 = true
}

// X加载配置项从文件
func (this *X配置项) X加载配置项从文件(fileapth string) bool {
	var err error

	this.cfg, err = ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment:      true,
		SpaceBeforeInlineComment: true,
		Insensitive:              this.节名区分大小写,
	}, fileapth)

	if err != nil {
		fmt.Println("配置项加载失败: ", err)
		return false
	}
	return true
}

// X加载配置项从内存
func (this *X配置项) X加载配置项从内存(v string) bool {
	var err error

	this.cfg, err = ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment:      true,
		SpaceBeforeInlineComment: true,
		Insensitive:              this.节名区分大小写,
	}, X到字节集(v))

	if err != nil {
		fmt.Println("配置项加载失败:", err)
		return false
	}
	return true
}

// E取配置项
func (this *X配置项) X读配置项(节名称 string, 配置项名称 string, 默认文本 string) string {
	v := this.cfg.Section(节名称).Key(配置项名称).String()
	if v == "" {
		return 默认文本
	}
	return v
}

// E取配置项 读取配置项
func (this *X配置项) X写配置项(节名称 string, 配置项名称 string, 设置文本 string) bool {
	this.cfg.Section(节名称).Key(配置项名称).SetValue(设置文本)
	return true
}

// X写到文件 写入到文件
func (this *X配置项) X写到文件(fileapth string) {
	this.cfg.SaveTo(fileapth)
}

// X取节名 读取节名
func (this *X配置项) X取节名() []string {
	names := this.cfg.SectionStrings()
	names = append(names[:0], names[1:]...)
	return names
}

// X取项名 读取项名
func (this *X配置项) X取项名(节名称 string) []string {
	names := this.cfg.Section(节名称).KeyStrings()
	return names
}
