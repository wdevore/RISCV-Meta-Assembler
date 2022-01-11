package api

type IAssembler interface {
	//
	Configure(configRelPath string) error

	Properties() IProperties
	ConfigRelPath() string
	ReportLine(line int, message string)
	ReportWhere(line int, where, message string)
	ReportToken(token IToken, message string)
	ErrorOccurred() bool
	SetError(occurred bool)

	// The main process
	Run(source string) error

	// Print()
}
