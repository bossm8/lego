// Code generated by 'internal/releaser'; DO NOT EDIT.

package main

const defaultVersion = "v4.20.3+dev-release"

var version = ""

func getVersion() string {
	if version == "" {
		return defaultVersion
	}

	return version
}