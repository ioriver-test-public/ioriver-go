// increment_version.go
package ioriver

import (
    "fmt"
    "io/ioutil"
    "log"
    "regexp"
    "strconv"
)

func main() {
    versionFilePath := "version/version.go"
    data, err := ioutil.ReadFile(versionFilePath)
    if err != nil {
        log.Fatalf("Error reading version file: %v", err)
    }

    versionContent := string(data)

    // Define a regular expression to match the Minor version line.
    regex := regexp.MustCompile(`Minor\s+=\s+(\d+)`)

    // Find the current Minor version in the content.
    matches := regex.FindStringSubmatch(versionContent)
    if len(matches) != 2 {
        log.Fatal("Failed to extract the current Minor version")
    }

    // Extract the current Minor version and increment it.
    currentMinorVersion, err := strconv.Atoi(matches[1])
    if err != nil {
        log.Fatalf("Error parsing current Minor version: %v", err)
    }
    newMinorVersion := currentMinorVersion + 1

    // Replace the current Minor version with the new one in the content.
    updatedVersionContent := regex.ReplaceAllString(versionContent, fmt.Sprintf("Minor = %d", newMinorVersion))

    // Write the updated content back to the version file.
    err = ioutil.WriteFile(versionFilePath, []byte(updatedVersionContent), 0644)
    if err != nil {
        log.Fatalf("Error writing version file: %v", err)
    }

    fmt.Printf("Minor version updated to %d\n", newMinorVersion)
}
