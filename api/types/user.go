package types

type User struct {
	ID                string `json:"id"`
	Email             string `json:"email"`
	Pfp               string `json:"pfp"`
	Banner            string `json:"banner"`
	ClubhouseTag      string `json:"clubhouse_tag"`
	Bot               bool   `json:"bot"`
	Username          string `json:"username"`
	GlobalDisplayName string `json:"global_display_name"`
}

type UserProfile struct {
	User
	Bio       string `json:"bio"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

func NewUser(id, email, pfp, banner, clubhouseTag string, bot bool, username, globalDisplayName string) User {
	return User{
		ID:                id,
		Email:             email,
		Pfp:               pfp,
		Banner:            banner,
		ClubhouseTag:      clubhouseTag,
		Bot:               bot,
		Username:          username,
		GlobalDisplayName: globalDisplayName,
	}
}

func NewUserProfile(user User, bio, status, createdAt string) UserProfile {
	return UserProfile{
		User:      user,
		Bio:       bio,
		Status:    status,
		CreatedAt: createdAt,
	}
}
