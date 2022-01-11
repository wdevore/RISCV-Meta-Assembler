package api

type IReporter interface {
	ReportLine(line int, message string)
	ReportWhere(line int, where string, message string)
}
