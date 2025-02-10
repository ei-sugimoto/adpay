package vo

type ProjectName string

func NewProjectName(name string) ProjectName {
	return ProjectName(name)
}

func (n ProjectName) String() string {
	return string(n)
}
