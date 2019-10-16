package pojo


type Student struct {
	id int
	loginName string
	userEmail string
	userPassword string
}

type ResMap struct {
	val map[string]string
	err error
}