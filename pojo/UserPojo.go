package pojo


type Student struct {
	Id string
	LoginName string
	UserEmail string
	UserPassword string
}

type ResMap struct {
	val map[string]string
	err error
}