// version/version.go
package version

import (
    "fmt"
)

var (
    Major          = 1
    Minor          = 0
    Patch          = 0
    CommitHash     = "n/a"
    BuildTimestamp = "n/a"
)

func BuildVersion() string {
    return fmt.Sprintf("%d.%d.%d-%s (%s)", Major, Minor, Patch, CommitHash, BuildTimestamp)
}

func IncrementMinorVersion() {
    Minor++
}
