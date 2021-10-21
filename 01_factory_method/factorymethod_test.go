package factorymethod

import "testing"

func TestPlusOperator(t *testing.T) {
	// 制定一个工厂 - plus operator的工厂
	factory := PlusOperatorFactory{}

	// 生产一个plus operator
	op := factory.Create()

	// 设置参数
	op.SetA(1)
	op.SetB(2)
	// 计算结果
	result := op.Result()

	if result != 3 {
		t.Fatal("error with plus operator factory method pattern")
	}
}

func TestMinusOperator(t *testing.T) {
	factory := MinusOperatorFactory{}
	op := factory.Create()

	op.SetA(1)
	op.SetB(2)
	result := op.Result()

	if result != -1 {
		t.Fatal("error with minus operator factory method pattern")
	}
}
