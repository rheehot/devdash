package internal

import (
	"github.com/Phantas0s/devdash/internal/platform"
)

type monitorServerWidget struct {
	tui    *Tui
	client platform.SSH
}

func NewMonitorServerWidget(username, addr string) (*monitorServerWidget, error) {
	sshClient, err := platform.NewSSHClient(username, addr)
	if err != nil {
		return nil, err
	}

	return &monitorServerWidget{
		client: sshClient,
	}, nil
}

func (ms *monitorServerWidget) CreateWidgets(widget Widget, tui *Tui) (f func() error, err error) {
	return nil, nil
}