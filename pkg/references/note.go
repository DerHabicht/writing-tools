package references

type Note struct {
	Citation string   `yaml:"citation"`
	Quote    string   `yaml:"quote"`
	Remarks  []string `yaml:"remarks"`
}
