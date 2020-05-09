package main

import (
	"ginform/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
