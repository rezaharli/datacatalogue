package models

type LinkRolePeople struct {
	ID          int
	Role_ID     Role
	People_ID   People
	Object_Type string
	Object_ID   int
}
