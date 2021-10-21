package abstractfactory

import "fmt"

// OrderMainDAO is the interface for the main DAO
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO is the interface for the detail DAO
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory is the interface for the factory
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO     // CreateOrderMainDAO returns an OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO // CreateOrderDetailDAO returns an OrderDetailDAO
}

// RDBMainDAO is the concrete implementation of the main DAO
type RDBMainDAO struct{}

// SaveOrderMain saves the order main info in the database
func (*RDBMainDAO) SaveOrderMain() {
	fmt.Print("RDBMainDAO: Saving the order in the database...\n")
}

// RDBDetailDAO is the concrete implementation of the detail DAO
type RDBDetailDAO struct{}

// SaveOrderDetail saves the order detials info in the database
func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Print("RDBDetailDAO: Saving the order detail in the database...\n")
}

// RDBDAOFactory is the concrete implementation of the factory
type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// XMLMainDAO is the concrete implementation of the main DAO
type XMLMainDAO struct{}

// SaveOrderMain saves the order main info in the database
func (*XMLMainDAO) SaveOrderMain() {
	fmt.Print("XMLMainDAO: Saving the order in the database...\n")
}

// XMLDetailDAO is the concrete implementation of the detail DAO
type XMLDetailDAO struct{}

// SaveOrderDetail saves the order detials info in the database
func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Print("XMLDetailDAO: Saving the order detail in the database...\n")
}

// XMLDAOFactory is the concrete implementation of the factory
type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
