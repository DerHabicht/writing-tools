package papers

type Reference struct {
	RefType string            `yaml:"type"`
	Title   string            `yaml:"title"`
	Author  string            `yaml:"author"`
	Date    string            `yaml:"date"`
	Bibdata map[string]string `yaml:"bibdata"`
	Notes   []Note            `yaml:"notes"`
}

type Note struct {
	Citation string   `yaml:"citation"`
	Quote    string   `yaml:"quote"`
	Remarks  []string `yaml:"remarks"`
}
