package main

import (
	"ginupload/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
