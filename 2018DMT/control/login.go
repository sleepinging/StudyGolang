package control

import "../dao"
import "../models"

func CheckLogin(email, pwd string) (res int) {
	res = dao.CheckLogin(&models.Login{Email: email, Password: pwd})
	return
}
