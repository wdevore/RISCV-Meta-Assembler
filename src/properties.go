package src

type configJSON struct {
	BinaryName string
	Generate   string // "Binary", "Ascii"
}

type Properties struct {
	Config configJSON
	Source []string
}

func (p *Properties) BinaryName() string {
	return p.Config.BinaryName
}

func (p *Properties) Files() []string {
	return p.Source
}
