package deps

import (
	auth "ispick-project-21022023/pkg/auth"
	"ispick-project-21022023/pkg/home"
	usr "ispick-project-21022023/pkg/user"
)

type Dependencies struct {
	Auth auth.AuthService
	User usr.UserService
	Home home.HomeService
}
