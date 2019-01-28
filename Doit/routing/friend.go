package routing

import (
"github.com/go-ozzo/ozzo-routing"
	"Project/Doit/handler/friend"
)

func FriendRegisterRoutes(router *routing.RouteGroup) {
	router.Get("/users",friend.QueryUsers)					//查询用户
}
