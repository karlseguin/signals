// a thin wrapper for common tasks involving os/signal
package signals

import (
	"os"
	ossignal "os/signal"
	"syscall"
)

const (
	SIGABRT   = syscall.SIGABRT
	SIGALRM   = syscall.SIGALRM
	SIGBUS    = syscall.SIGBUS
	SIGCHLD   = syscall.SIGCHLD
	SIGCONT   = syscall.SIGCONT
	SIGFPE    = syscall.SIGFPE
	SIGHUP    = syscall.SIGHUP
	SIGILL    = syscall.SIGILL
	SIGINT    = syscall.SIGINT
	SIGIO     = syscall.SIGIO
	SIGIOT    = syscall.SIGIOT
	SIGKILL   = syscall.SIGKILL
	SIGPIPE   = syscall.SIGPIPE
	SIGPROF   = syscall.SIGPROF
	SIGQUIT   = syscall.SIGQUIT
	SIGSEGV   = syscall.SIGSEGV
	SIGSTOP   = syscall.SIGSTOP
	SIGSYS    = syscall.SIGSYS
	SIGTERM   = syscall.SIGTERM
	SIGTRAP   = syscall.SIGTRAP
	SIGTSTP   = syscall.SIGTSTP
	SIGTTIN   = syscall.SIGTTIN
	SIGTTOU   = syscall.SIGTTOU
	SIGURG    = syscall.SIGURG
	SIGUSR1   = syscall.SIGUSR1
	SIGUSR2   = syscall.SIGUSR2
	SIGVTALRM = syscall.SIGVTALRM
	SIGWINCH  = syscall.SIGWINCH
	SIGXCPU   = syscall.SIGXCPU
	SIGXFSZ   = syscall.SIGXFSZ
)

type Handler func() bool

func Handle(signal os.Signal, handler Handler) {
	channel := make(chan os.Signal, 1)
	ossignal.Notify(channel, signal)
	for {
		<-channel
		if handler() == false {
			break
		}
	}
}
