package model

type StatusRegistration string

const (
	Pending StatusRegistration = "PENDING"
	Accepted StatusRegistration = "ACCEPTED"
	Rejected StatusRegistration = "REJECTED"
)