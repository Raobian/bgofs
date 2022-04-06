package version

import (
	"fmt"
	"strings"
)

var (
	revision     = "$Format:%h$"
	revisionDate = "$Format:%as$"
	ver          = Semver{
		major:      1,
		minor:      0,
		patch:      0,
		preRelease: "dev",
		build:      fmt.Sprintf("%s.%s", revisionDate, revision),
	}
)

type Semver struct {
	major, minor, patch uint64
	preRelease, build   string
}

func Version() string {
	pr := ver.preRelease
	if pr != "" {
		pr = "-" + pr
	}
	if strings.Contains(ver.build, "Format") {
		ver.build = "unknown"
	}
	return fmt.Sprintf("%d.%d.%d%s+%s", ver.major, ver.minor, ver.patch, pr, ver.build)
}
