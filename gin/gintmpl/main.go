package main

import (
	"gintmpl/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
