package git

import (
	"strings"
)

type GitStatus struct {
    Staged []string
    Unstaged []string
    Untracked []string
}

func parseGitStatus(status_output string) (*GitStatus, error) {
    lines := strings.Split(status_output, "\n")

    gs := GitStatus {
        Staged: []string{},
        Unstaged: []string{},
        Untracked: []string{},
    }
    for _, v := range lines {
        lineSplit := strings.Split(strings.Trim(v, " "), " ")
        t := lineSplit[0]
        fp := lineSplit[1]

        switch t {
        case "A":
            gs.Staged = append(gs.Staged, fp)
        case "M":
            gs.Unstaged = append(gs.Unstaged, fp)
        case "MM":
            gs.Unstaged = append(gs.Unstaged, fp)
        case "??":
            gs.Untracked = append(gs.Untracked, fp)
        }
    }

    return &gs, nil
}
