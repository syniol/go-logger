package sylog

type Logger interface {
	LogInfo(args ...interface{})
	LogDebug(args ...interface{})
	LogNotice(args ...interface{})
	LogWarning(args ...interface{})
	LogAlert(args ...interface{})
	LogEmergency(args ...interface{})
	LogError(args ...interface{})
	LogCritical(args ...interface{})
}

func New(facility string) Logger {
	return &logger{
		Facility: facility,
	}
}

// LogInfo logs messages where in output JSON key "level" is "info"
func (l *logger) LogInfo(args ...interface{}) {
	write(levelInfo, l.Facility, args)
}

// LogDebug logs messages where in output JSON key "level" is "debug"
func (l *logger) LogDebug(args ...interface{}) {
	write(levelDebug, l.Facility, args)
}

// LogNotice logs messages where in output JSON key "level" is "notice"
func (l *logger) LogNotice(args ...interface{}) {
	write(levelNotice, l.Facility, args)
}

// LogWarning logs messages where in output JSON key "level" is "warn"
func (l *logger) LogWarning(args ...interface{}) {
	write(levelWarning, l.Facility, args)
}

// LogAlert logs messages where in output JSON key "level" is "alert"
func (l *logger) LogAlert(args ...interface{}) {
	write(levelAlert, l.Facility, args)
}

// LogEmergency logs messages where in output JSON key "level" is "emergency"
func (l *logger) LogEmergency(args ...interface{}) {
	write(levelEmergency, l.Facility, args)
}

// LogError logs messages where in output JSON key "level" is "error"
func (l *logger) LogError(args ...interface{}) {
	write(levelError, l.Facility, args)
}

// LogCritical logs messages where in output JSON key "level" is "crit"
func (l *logger) LogCritical(args ...interface{}) {
	write(levelCritical, l.Facility, args)
}
