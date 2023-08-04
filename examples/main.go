package main

import "github.com/mriceman/go-uml/sequence"

func main() {
	userStartsChatting()
	userBIsNotPartOfChat()
	userBIsPartOfChat()

}

func userBIsPartOfChat() {
	d := sequence.NewDiagram("on_new_message_for_user_b")

	usrA := "User A"
	usrB := "User B"
	backend := "Backend"
	db := "DynamoDB"

	d.SetTitle("User A sends a message, User B is connected to the chat and receives it")

	d.AddParticipants(usrA, usrB, backend, db)

	d.AddDirectionalEdge(usrA, backend, "WS sendmessage/userB_id payload: hi")
	d.AddDirectionalEdge(backend, db, "get connection ID from userB")
	d.AddDirectionalEdge(backend, db, "store message")
	d.AddDirectionalEdge(backend, usrA, "new_message from: userA payload: hi")
	d.AddDirectionalEdge(backend, db, "set all messages read for usrA")
	d.Render()
}

func userBIsNotPartOfChat() {
	d := sequence.NewDiagram("on_new_message_for_user_b_not_read")

	usrA := "User A"
	usrB := "User B"
	backend := "Backend"
	db := "DynamoDB"

	d.SetTitle("User A sends a message, User B is not connected to the chat")

	d.AddParticipants(usrA, usrB, backend, db)

	d.AddDirectionalEdge(usrA, backend, "WS sendmessage/userB_id payload: hi")
	d.AddDirectionalEdge(backend, db, "get connection ID from userB")
	d.AddDirectionalEdge(backend, db, "store message")
	d.AddDirectionalEdge(backend, db, "increment_unread_count_for_user_b")
	d.AddDirectionalEdge(backend, backend, "create_push_for_user_b")
	d.AddDirectionalEdge(usrB, usrB, "receives a push notification")
	d.AddDirectionalEdge(usrB, usrB, "navigates to chats")
	d.AddDirectionalEdge(usrB, backend, "GET /chat")
	d.AddDirectionalEdge(backend, db, "Query * chats for userB")
	d.AddDirectionalEdge(backend, usrB, "return chats for userB")
	d.AddDirectionalEdge(
		usrB,
		usrB,
		"renders list of chats + unread count")
	d.Render()
}

func userStartsChatting() {
	d := sequence.NewDiagram("user_starts_chatting")

	client := "Client"
	backend := "Backend"
	db := "DynamoDB"

	d.SetTitle("User initiates chat with another User")

	d.AddParticipants(client)
	d.AddParticipants(backend)
	d.AddParticipants(db)

	d.AddDirectionalEdge(client, backend, "PUT /chat/user/<TO_ID>")
	d.AddDirectionalEdge(backend, db, "checks or create inbox for user")
	d.AddDirectionalEdge(backend, db, "set all unread messages to read if existing")
	d.AddDirectionalEdge(backend, client, "return chat meta data + chat secret")
	d.AddDirectionalEdge(client, backend, "WS /joinchat/<CHAT_ID> payload: secret")
	d.AddDirectionalEdge(backend, db, "compare inbox secret")
	d.AddDirectionalEdge(backend, db, "set all unread messages to read")
	d.AddDirectionalEdge(backend, db, "store connection ID for user")
	d.AddDirectionalEdge(backend, client, "202 - You're good to go")
	d.AddDirectionalEdge(client, backend, "WS /sendmessage/<CHAT_ID> payload: hello!")
	d.AddDirectionalEdge(backend, db, "get connection ID for <TO_ID>")
	d.AddDirectionalEdge(backend, backend, "notify <TO_ID> with connection ID if present")
	d.AddDirectionalEdge(backend, db, "store message in chat")
	d.AddDirectionalEdge(client, backend, "disconnect")
	d.AddDirectionalEdge(backend, db, "remove connection ID for user")

	d.Render()

}
