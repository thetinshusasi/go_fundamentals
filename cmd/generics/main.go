package main

// import "fmt"

// type connect interface {
// 	myconnection()
// }

// type test interface {
// 	mytest(t connect)
// }

// type MySQLConnect struct{}

// func (m MySQLConnect) myconnection() {
// 	fmt.Println("MySQL connection established")
// }

// type MySQLTester struct{}

// func (t MySQLTester) mytest(c MySQLConnect) {
// 	c.myconnection()
// 	fmt.Println("Testing MySQL connection")
// }

// func main() {
// 	mysql := MySQLConnect{}
// 	tester := MySQLTester{}

// 	tester.mytest(mysql)
// }

//  generics
//  constraint
//  paramerterized constraint

import "fmt"

type connect interface {
	myconnection()
}

type MySQLConnect struct{}

func (m MySQLConnect) myconnection() {
	fmt.Println("MySQL connection established")
}

type PostgresConnect struct{}

func (p PostgresConnect) myconnection() {
	fmt.Println("Postgres connection established")
}

type test2[c connect] interface {
	mytest(c)
}

type MySQLTester struct{}

func (t MySQLTester) mytest(conn connect) {
	conn.myconnection()
	fmt.Println("Testing MySQL connection")
}

type PostgresTester struct{}

func (t PostgresTester) mytest(conn connect) {
	conn.myconnection()
	fmt.Println("Testing Postgres connection")
}

func main() {
	mysql := MySQLConnect{}
	postgres := PostgresConnect{}
	mysqlTester := MySQLTester{}
	postgresTester := PostgresTester{}

	mysqlTester.mytest(postgres)
	postgresTester.mytest(mysql)
}
