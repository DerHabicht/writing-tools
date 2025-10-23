package references

type Reference struct {
	BibData map[string]string `yaml:"bibdata"`
	Notes   []Note            `yaml:"notes"`
}
