package main

import (
	"ginmiddleware/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
