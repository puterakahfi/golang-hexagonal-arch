package server

import "fmt"

func GetServer(serverAdapter string) (ServerInterface, error) {
	switch serverAdapter {
	case "gin":
		return InitGinServer(), nil
	}
	return nil, fmt.Errorf("server type " + serverAdapter + " not supported")

}
