package models

import (
  "time"
)
type Logger struct{
  Printer PrinterStratedy
}

func NewLogger(printer PrinterStratedy) *Logger{
  return &Logger{
    Printer: printer,
  }
}

func(l *Logger)Debug(msg string){
  l.Printer.Log(l.formatMessage("DEBUG", msg))
}

func(l *Logger)Error(msg string){
  l.Printer.Log(l.formatMessage("ERROR", msg))
}

func(l *Logger)Warn(msg string){
  l.Printer.Log(l.formatMessage("WARN", msg))
}

func(l *Logger)Info(msg string){
  l.Printer.Log(l.formatMessage("INFO", msg))
}


func(l *Logger) formatMessage(level, msg string) string {
  return "[" + level + "] " + time.Now().Format("2006-01-02 15:04:05") + " " + msg
}