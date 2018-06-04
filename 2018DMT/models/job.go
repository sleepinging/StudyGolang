package models

//除了id全部复制
func (this *Job) CopyFormEId(job *Job) {
	this.Name = job.Name
	this.Salary = job.Salary
	this.Time = job.Time
	this.Weekend = job.Weekend
	this.Pickup = job.Pickup
	this.Eat = job.Eat
	this.Live = job.Live
	this.WuXianYiJin = job.WuXianYiJin
	this.Place = job.Place
	this.LimPeople = job.LimPeople
	this.NowPeople = job.NowPeople
	this.Sex = job.Sex
	this.Phone = job.Phone
	this.Detail = job.Detail
}
