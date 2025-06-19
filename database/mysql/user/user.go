package mysql

func createUser(name, email, password string) {
	baseQuery := `INSERT INTO user (name,email,password) VALUES (@name,@email,@password)`

}
