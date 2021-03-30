// @Description: 单测
// @Author: Arvin
// @Date: 2021/3/17 5:25 下午
package database

import (
	"os"
	"testing"
)

// 测试从此方法进入，依次执行测试用例，最后从此方法退出
func TestMain(m *testing.M) {
	// TODO 测试初始化操作

	// 执行完之后退出
	os.Exit(m.Run())
}

// sql.Open 不会创建连接 ，只会创建一个DB实例
// 同时会创建一个go程来管理该DB实例的一个连接池(是长连接，但不是在Open的时候创建)
func TestDBModel_Connect(t *testing.T) {
	// 测试数据结构
	type data struct {
		name   string    // 测试说明
		param  *DBConfig // 测试参数
		result bool      // 预期结果
	}
	// 初始化测试数据
	datas := []data{
		{
			name: "测试连接方法",
			param: &DBConfig{
				DBType:   "mysql",
				Host:     "127.0.0.1:2306",
				DBName:   "test",
				UserName: "test",
				PassWord: "",
				CharSet:  "utf8mb4",
			},
			result: false,
		},
	}

	for _, d := range datas {
		t.Run(d.name, func(t *testing.T) {
			mod := NewDBModel(d.param)
			if err := mod.Connect(); (err != nil) != d.result {
				t.Errorf("connect db error:%s param:%+v", err, d.param)
			}
			t.Log(mod)
		})
	}
}
