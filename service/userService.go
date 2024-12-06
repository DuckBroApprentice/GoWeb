package service

type UserService struct{}

func (us *UserService) Login(name, password string) *models.User{
	ud := dao.UserDao{tool.DbEngine}
	user := ud.Query(name,password)
	if user.Id != 0 {
		return user
	}
	member := models.User{}
	member.UserName = name
	user.Password = tool.EncoderSha256(password)
	user.RegisterTime = time.Now().Unix()

	result := ud.InsertUsername(user)
	user.Id = result
	return &user
}