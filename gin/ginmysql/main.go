package main

import (
	"ginmysql/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
