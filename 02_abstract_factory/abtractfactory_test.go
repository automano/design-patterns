package abstractfactory

import (
	"log"
	"runtime"
)

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func getMainAndDetail(factory DAOFactory) {
	log.Printf("函数名: %s", runFuncName())
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func ExampleRDBDAOFactory() {
	log.Print("函数名: ", runFuncName())
	factory := &RDBDAOFactory{}
	getMainAndDetail(factory)
}

func ExampleXMLDAOFactory() {
	log.Print("函数名: ", runFuncName())
	factory := &XMLDAOFactory{}
	getMainAndDetail(factory)
}
