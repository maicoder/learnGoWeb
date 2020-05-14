package main

import (
	"ginrestful/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
