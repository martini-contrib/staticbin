package staticbin

// Options represents configuration options for the staticbin.Static middleware.
type Options struct {
	// SkipLogging will disable [Static] log messages when a static file is served.
	SkipLogging bool
}

// retrieveOptions retrieves an options from the array of options.
func retrieveOptions(options []Options) Options {
	var opt Options

	if len(options) > 0 {
		opt = options[0]
	}

	return opt
}
