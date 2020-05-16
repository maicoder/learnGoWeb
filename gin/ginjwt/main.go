package main

import (
	"ginjwt/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
