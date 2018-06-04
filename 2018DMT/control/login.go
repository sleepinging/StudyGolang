package control

import "../dao"
import "../models"

func CheckLogin(email, pwd string) (res int, user *models.User, err error) {
	res = dao.CheckLogin(&models.Login{Email: email, Password: pwd})
	if res == 1 {
		user, err = dao.GetUserByEmail(email)
	}
	return
}
