package git

import "testing"

func Test(t *testing.T) {
    status_output := `A README.md
                      MM internal/git/gstatus.go
                      M internal/git/gstatus_test.go
                      ?? cmd/
                      ?? go.mod`

   gs, err := parseGitStatus(status_output)

   if err != nil {
       t.Fatalf("Expected err to be null, got %s\n", err)
   }

   if gs == nil {
       t.Fatalf("Expected gs to not be null\n")
   }

   if len(gs.Staged) != 2 {
       t.Fatalf("Expected to get %d staged files, got %d\n", 1, len(gs.Staged))
   }

   if len(gs.Unstaged) != 2 {
       t.Fatalf("Expected to get %d unstaged files, got %d\n", 2, len(gs.Unstaged))
   }

   if len(gs.Untracked) != 2 {
       t.Fatalf("Expected to get %d untracked files, got %d\n", 2, len(gs.Untracked))
   }
}

