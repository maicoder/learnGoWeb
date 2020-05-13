package main

import (
	"gincookie/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
