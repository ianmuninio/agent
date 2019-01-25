// +build !windows

package process

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/buildkite/agent/logger"
)

func SetupProcessGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
		Pgid:    0,
	}
}

func TerminateProcessGroup(p *os.Process, l *logger.Logger) error {
	l.Debug("[Process] Sending signal SIGKILL to PGID: %d", p.Pid)
	return syscall.Kill(-p.Pid, syscall.SIGKILL)
}

func InterruptProcessGroup(p *os.Process, l *logger.Logger) error {
	l.Debug("[Process] Sending signal SIGTERM to PGID: %d", p.Pid)

	// TODO: this should be SIGINT, but will be a breaking change
	return syscall.Kill(-p.Pid, syscall.SIGTERM)
}

func GetPgid(pid int) (int, error) {
	return syscall.Getpgid(pid)
}

func SignalString(s syscall.Signal) string {
	switch int(s) {
	case 1:
		return "SIGHUP"
	case 2:
		return "SIGINT"
	case 3:
		return "SIGQUIT"
	case 4:
		return "SIGILL"
	case 5:
		return "SIGTRAP"
	case 6:
		return "SIGABRT"
	case 7:
		return "SIGEMT"
	case 8:
		return "SIGFPE"
	case 9:
		return "SIGKILL"
	case 10:
		return "SIGBUS"
	case 11:
		return "SIGSEGV"
	case 12:
		return "SIGSYS"
	case 13:
		return "SIGPIPE"
	case 14:
		return "SIGALRM"
	case 15:
		return "SIGTERM"
	case 16:
		return "SIGURG"
	case 17:
		return "SIGSTOP"
	case 18:
		return "SIGTSTP"
	case 19:
		return "SIGCONT"
	case 20:
		return "SIGCHLD"
	case 21:
		return "SIGTTIN"
	case 22:
		return "SIGTTOU"
	case 23:
		return "SIGIO"
	case 24:
		return "SIGXCPU"
	case 25:
		return "SIGXFSZ"
	case 26:
		return "SIGVTALRM"
	case 27:
		return "SIGPROF"
	case 28:
		return "SIGWINCH"
	case 29:
		return "SIGINFO"
	case 30:
		return "SIGUSR1"
	case 31:
		return "SIGUSR2"
	}
	return fmt.Sprintf("%d", int(s))
}
