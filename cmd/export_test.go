package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExportCommand(t *testing.T) {

	apiToken := os.Getenv("YUNDICT_API_TOKEN")
	testTeam := os.Getenv("YUNDICT_TEST_TEAM_NAME")
	testProject := os.Getenv("YUNDICT_TEST_PROJECT_NAME")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"export", "--team", testTeam, "--project", testProject, "--token", apiToken})
	rootCmd.Execute()

	expected := "This-is-command-a1"

	println(actual.String())
	assert.Equal(t, actual.String(), expected, "actual is not expected")
}
