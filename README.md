# go-uml
Just a little tool to create UML Diagrams built in Go. You can use it already to build simple sequence diagrams, however, this project is still under development
and you won't find some functionalities yet such as

- provide conditional flows

The generated diagram is saved as .PNG file

So far, you can create only sequence diagrams and add Participants, directional and undirectional Edges, labels for Edges and set a Title for the diagram.
I'll be updating this repository whenever I need the tool to support more functionality, feel free to create an Issue with a feature request. Since I just started this project, contributing should also be quite easy (I appreciate any contribution).

You don't need to download any dependencies such as plantUML or Graphviz, which is what most of the tools out there require and what was also my motivation to start this project. go-uml is using a 2D graphics engine written 100% in Go https://github.com/fogleman/gg

# Example

```
 	d := sequence.NewDiagram("user_starts_chatting")

	client := "Client"
	backend := "Backend"
	db := "DynamoDB"

	d.SetTitle("User initiates chat with another User")
	d.AddParticipant(client)
	d.AddParticipant(backend)
	d.AddParticipant(db)

	d.AddDirectionalEdge(client, backend, "PUT /chat/user/<TO_ID>")
	d.AddDirectionalEdge(backend, db, "checks or create inbox for user")
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

```
Result PNG file: 
![image description](./examples/user_starts_chatting.png)
