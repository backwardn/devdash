package internal

import (
	"github.com/Phantas0s/devdash/internal/platform"
	"github.com/pkg/errors"
)

const (
	gitBranches = "git.table_branches"
)

type gitWidget struct {
	tui    *Tui
	client *platform.Git
}

func NewGitWidget(path string) *gitWidget {
	client := platform.NewGit(path)
	return &gitWidget{
		client: client,
	}
}

func (g gitWidget) CreateWidgets(widget Widget, tui *Tui) (f func() error, err error) {
	g.tui = tui

	switch widget.Name {
	case gitBranches:
		f, err = g.branches(widget)
	default:
		return nil, errors.Errorf("can't find the widget %s for service Git", widget.Name)
	}

	return
}

func (g gitWidget) branches(widget Widget) (f func() error, err error) {
	title := " Git Branches "
	if _, ok := widget.Options[optionTitle]; ok {
		title = widget.Options[optionTitle]
	}

	data, err := g.client.Branches()
	if err != nil {
		return nil, err
	}

	f = func() error {
		return g.tui.AddTable(data, title, widget.Options)
	}

	return
}
