include $(GOROOT)/src/Make.inc

TARG=log4go.googlecode.com/svn/stable
GOFILES=\
	log4go.go\
	config.go\
  termlog.go\
	socklog.go\
	filelog.go\
	pattlog.go\
	xmllog.go\
  wrapper.go

DIRS=\

include $(GOROOT)/src/Make.pkg

elogtest :
	gotest -benchmarks=.* > /tmp/gotest && cat /tmp/gotest | grep -v WARN | grep -v INFO
