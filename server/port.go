package server

type ServerInterface interface {
	Run()
	GetServerInstance() interface{}
}
