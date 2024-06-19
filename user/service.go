package user

const (
	NoName = "Без имени"
)

type UserService struct {
	Repo UserRepoI
}

func (h *UserService) GetUserName(id int) (string, error) {
	user, err := h.Repo.GetUserByID(id)
	if err != nil {
		return "", err
	}

	if user == nil {
		return NoName, nil
	}

	return user.Name, nil
}
