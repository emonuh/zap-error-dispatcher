package core

import "go.uber.org/zap/zapcore"

func NewErrorDispatcher(base zapcore.Core, err zapcore.Core) zapcore.Core {
	return &errorDispatcher{
		base: base,
		err:  err,
	}
}

type errorDispatcher struct {
	base zapcore.Core
	err  zapcore.Core
}

func (e *errorDispatcher) With(fields []zapcore.Field) zapcore.Core {
	clone := e.clone()
	clone.base = e.base.With(fields)
	clone.err = e.err.With(fields)
	return clone
}

func (e *errorDispatcher) Enabled(lvl zapcore.Level) bool {
	return e.base.Enabled(lvl)
}

func (e *errorDispatcher) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if ent.Level >= zapcore.ErrorLevel {
		return e.err.Check(ent, ce)
	}
	return e.base.Check(ent, ce)
}

func (e *errorDispatcher) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	if ent.Level >= zapcore.ErrorLevel {
		return e.err.Write(ent, fields)
	}
	return e.base.Write(ent, fields)
}

func (e *errorDispatcher) Sync() error {
	if err := e.err.Sync(); err != nil {
		return err
	}
	return e.base.Sync()
}

func (e *errorDispatcher) clone() *errorDispatcher {
	return &errorDispatcher{
		base: e.base,
		err:  e.err,
	}
}
