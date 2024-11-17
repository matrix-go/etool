package options

type ParserMode int

const (
	ParserModeUnSupport ParserMode = iota
	ParserModeDDLPath
	ParserModeDSNConnect
)

type Option struct {
	DDLPath string
	Prefix  string
	DSN     string
	Table   string
}

func (o *Option) Validate() ParserMode {
	if o.DDLPath != "" && o.Prefix != "" {
		return ParserModeDDLPath
	}
	if o.DSN != "" && o.Prefix != "" && o.Table != "" {
		return ParserModeDSNConnect
	}
	return ParserModeUnSupport
}

type ParserOption func(*Option)

func WithDDLPath(ddlPath string) func(o *Option) {
	return func(o *Option) {
		o.DDLPath = ddlPath
	}
}

func WithPrefix(prefix string) func(o *Option) {
	return func(o *Option) {
		o.Prefix = prefix
	}
}

func WithDsn(dsn string) func(o *Option) {
	return func(o *Option) {
		o.DSN = dsn
	}
}

func WithTable(table string) ParserOption {
	return func(o *Option) {
		o.Table = table
	}
}
