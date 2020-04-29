package reports

type Format int

const (
	UndefinedFormat Format = iota
	Json
	Html
)

func formatStrToEnum(formatStr string) Format {
	switch formatStr {
	case "json":
		return Json
	case "html":
		return Html
	}

	return Json
}
