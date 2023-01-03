package app

import (
	"github.com/an1l4/bookstore_users_api/controllers/ping"
	"github.com/an1l4/bookstore_users_api/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/user", user.CreateUser)

	router.GET("/user/:user_id", user.GetUser)

	router.GET("/user/search", user.SearchUser)

}
