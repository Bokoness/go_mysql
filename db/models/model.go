package models

type Model interface {
	Save()
	Update()
	Destroy()
}
