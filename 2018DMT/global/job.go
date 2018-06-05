package global

import (
	"../models"
)

var (
	//工作类型
	Jobtypedess models.JobTypeDess
)

//通过工作类型的ID找到工作类型
func FindJobTypeById(id int) (ok bool, jtd models.JobTypeDes, jt2 models.JobType2, jt models.JobType) {
	for _, jtd = range Jobtypedess.Types {
		for _, jt2 = range jtd.Types {
			for _, jt = range jt2.Types {
				if jt.JobId == id {
					ok = true
					return
				}
			}
		}
	}
	return
}

//通过工作类型的名字找到工作类型
func FindJobTypeByName(name string) (ok bool, jtd models.JobTypeDes, jt2 models.JobType2, jt models.JobType) {
	for _, jtd = range Jobtypedess.Types {
		for _, jt2 = range jtd.Types {
			for _, jt = range jt2.Types {
				if jt.JobName == name {
					ok = true
					return
				}
			}
		}
	}
	return
}