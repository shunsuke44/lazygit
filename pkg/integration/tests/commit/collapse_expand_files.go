package commit

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var CollapseExpandFiles = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Collapsing and expanding all files of a commit",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.CreateFileAndAdd("dir/file-one", "original content\n")
		shell.CreateFileAndAdd("dir2/file-two", "original content\n")
		shell.Commit("initial commit")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Commits().
			Focus().
			PressEnter()

		t.Views().CommitFiles().
			IsFocused().
			Lines(
				Equals("▼ /").IsSelected(),
				Equals("  ▼ dir"),
				Equals("    A file-one"),
				Equals("  ▼ dir2"),
				Equals("    A file-two"),
			).
			Press(keys.Files.CollapseAll).
			Lines(
				Equals("▶ /"),
			).
			Press(keys.Files.ExpandAll).
			Lines(
				Equals("▼ /").IsSelected(),
				Equals("  ▼ dir"),
				Equals("    A file-one"),
				Equals("  ▼ dir2"),
				Equals("    A file-two"),
			)
	},
})
