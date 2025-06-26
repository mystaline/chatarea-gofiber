package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/controllers"
	"github.com/mystaline/chatarea-gofiber/internal/app/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/ping", controllers.Me)

	v1 := app.Group("/api/v1")
	auth := v1.Group("/auth")

	// Routes for authentication (login & register)
	auth.Post("/login", controllers.Login)       // login✅
	auth.Post("/register", controllers.Register) // register✅

	// Add middleware to route that has prefix other than login & register
	authenticated := v1.Group("", middleware.ValidateJWT()) // ValidateJWT✅

	// Define prefix for each similar route group
	me := authenticated.Group("/me")
	myProfile := me.Group("/profile")
	myRooms := me.Group("/rooms")
	userId := authenticated.Group("/user/:id")
	rooms := authenticated.Group("/rooms")
	roomChats := rooms.Group("/:id/chats")

	// Different group from login and register since this need user to be authenticated
	authenticated.Delete("/logout", controllers.Me) // logout

	me.Get("/blocked-users", controllers.Me) // getBlockedUsers

	// Routes for interact with authenticated user's profile (Read, Update, Delete)
	myProfile.Get("/", controllers.Me)               // me✅
	myProfile.Put("/", controllers.EditProfile)      // editProfile✅
	myProfile.Delete("/", controllers.DeleteAccount) // deleteAccount✅

	// Routes for interact with authenticated user's rooms (Read, Update, Delete)
	myRooms.Get("/", controllers.GetMyRooms)       // getMyRooms✅
	myRooms.Put("/:address", controllers.JoinRoom) // joinRoom
	myRooms.Delete("/:id", controllers.LeaveRoom)  // leaveRoom

	userId.Get("/", controllers.Me)        // getOtherUserInfo
	userId.Put("/block", controllers.Me)   // blockUser
	userId.Put("/unblock", controllers.Me) // unblockUser

	rooms.Post("/", controllers.CreateRoom)      // createRoom✅
	rooms.Put("/:id/invite", controllers.Me)     // inviteUsersToRoom
	rooms.Delete("/:id/kick", controllers.Me)    // kickUserFromRoom
	rooms.Get("/:id", controllers.GetRoomInfo)   // getRoomInfo✅
	rooms.Put("/:id", controllers.EditRoomInfo)  // editRoomInfo
	rooms.Delete("/:id", controllers.DeleteRoom) // deleteRoom

	roomChats.Get("/", controllers.Me)              // getChats
	roomChats.Post("/", controllers.Me)             // sendChat
	roomChats.Put("/:id/edit", controllers.Me)      // editChat
	roomChats.Put("/open-unread", controllers.Me)   // openUnred
	roomChats.Delete("/:id/unsend", controllers.Me) // unsendChat
}
