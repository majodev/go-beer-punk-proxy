package test_test

import (
	"path/filepath"
	"testing"

	"github.com/majodev/go-beer-punk-proxy/internal/config"
	"github.com/majodev/go-beer-punk-proxy/internal/test"
	"github.com/majodev/go-beer-punk-proxy/internal/util"
	"github.com/stretchr/testify/assert"
)

// This test will run or skip depending upon if you currently have a `.env.local` in your project directory
// Note: We do not activate this test in the go-starter template, as it would always be skipped.
// func TestDotEnvLoadLocalOrSkipTest(t *testing.T) {
// 	test.DotEnvLoadLocalOrSkipTest(t)
// }

// This test will always run as the /internal/test/testdata/.env.test.local is checked into git.
func TestDotEnvLoadFileOrSkipTest(t *testing.T) {
	// explicitly load a (test) dotenv file before getting a new config (for a testserver)
	test.DotEnvLoadFileOrSkipTest(t, filepath.Join(util.GetProjectRootDir(), "/internal/test/testdata/.env.test.local"))
	c := config.DefaultServiceConfigFromEnv()
	assert.Equal(t, "http://overwritten.dotenv.frontend.url.tld:3000", c.Frontend.BaseURL)
}
