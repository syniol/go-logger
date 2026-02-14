package sylog

type Logger interface {
	LogInfo(args ...string)
	LogDebug(args ...string)
	LogNotice(args ...string)
	LogWarning(args ...string)
	LogAlert(args ...string)
	LogEmergency(args ...string)
	LogError(args ...string)
	LogCritical(args ...string)
}

func New(facility string, opt *LoggerOptions) Logger {
	return &logger{
		Facility:      facility,
		loggerOptions: opt,
	}
}

// LogInfo logs messages where in output JSON key "level" is "info"
func (l *logger) LogInfo(args ...string) {
	write(levelInfo, l.Facility, args)
}

// LogDebug logs messages where in output JSON key "level" is "debug"
func (l *logger) LogDebug(args ...string) {
	write(levelDebug, l.Facility, args)
}

// LogNotice logs messages where in output JSON key "level" is "notice"
func (l *logger) LogNotice(args ...string) {
	write(levelNotice, l.Facility, args)
}

// LogWarning logs messages where in output JSON key "level" is "warn"
func (l *logger) LogWarning(args ...string) {
	write(levelWarning, l.Facility, args)
}

// LogAlert logs messages where in output JSON key "level" is "alert"
func (l *logger) LogAlert(args ...string) {
	write(levelAlert, l.Facility, args)
}

// LogEmergency logs messages where in output JSON key "level" is "emergency"
func (l *logger) LogEmergency(args ...string) {
	write(levelEmergency, l.Facility, args)
}

// LogError logs messages where in output JSON key "level" is "error"
func (l *logger) LogError(args ...string) {
	write(levelError, l.Facility, args)
}

// LogCritical logs messages where in output JSON key "level" is "crit"
func (l *logger) LogCritical(args ...string) {
	write(levelCritical, l.Facility, args)
}
