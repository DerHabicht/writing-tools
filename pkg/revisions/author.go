package revisions

type Author interface {
	FullName() string
	ShortName() string
	Initials() string
	DutyAssignment() string
	MarshalYAML() (interface{}, error)
	UnmarshalYAML(unmarshal func(interface{}) error) error 
}
