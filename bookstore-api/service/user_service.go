package service

type UserService struct {
}

const (
	query = `SELECT USERID FROM USERS WHERE UCIC=:1`
)

func NewUserService() *UserService {
	return &UserService{}
}

type User struct {
	UserId string `db:"USERID"`
}

func (us UserService) GetUserById(userId string) (string, error) {
	var users User
	//query := `SELECT USERID FROM USERS WHERE UCIC=:1`
	// err := us.db.Select(users, query, userId)
	// if err != nil {
	// 	return fmt.Sprintf("error : %s", err), err
	// }
	return users.UserId, nil
}
