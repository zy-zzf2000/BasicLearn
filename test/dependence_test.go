package test

import (
	"Go/BasicLearn/channel"
	"testing"
)

//测试本地module的package导入
func TestHello(t *testing.T) {
	channel.Hello()
}
