//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// Contains common helpers for TESTS ONLY
package testcommon

import (
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azfile/fileerror"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

const (
	SharePrefix     = "gos"
	DirectoryPrefix = "godir"
	FilePrefix      = "gotestfile"
)

func GenerateShareName(testName string) string {
	return SharePrefix + GenerateEntityName(testName)
}

func GenerateEntityName(testName string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(testName), "/", ""), "test", "")
}

func GenerateDirectoryName(testName string) string {
	return DirectoryPrefix + GenerateEntityName(testName)
}

func GenerateFileName(testName string) string {
	return FilePrefix + GenerateEntityName(testName)
}

func ValidateFileErrorCode(_require *require.Assertions, err error, code fileerror.Code) {
	_require.Error(err)
	var responseErr *azcore.ResponseError
	errors.As(err, &responseErr)
	if responseErr != nil {
		_require.Equal(string(code), responseErr.ErrorCode)
	} else {
		_require.Contains(err.Error(), code)
	}
}

// GetRequiredEnv gets an environment variable by name and returns an error if it is not found
func GetRequiredEnv(name string) (string, error) {
	env, ok := os.LookupEnv(name)
	if ok {
		return env, nil
	} else {
		return "", errors.New("Required environment variable not set: " + name)
	}
}

func BeforeTest(t *testing.T, suite string, test string) {
	const urlRegex = `https://\S+\.file\.core\.windows\.net`
	require.NoError(t, recording.AddURISanitizer(FakeStorageURL, urlRegex, nil))
	require.NoError(t, recording.AddHeaderRegexSanitizer("x-ms-copy-source", FakeStorageURL, urlRegex, nil))
	// we freeze request IDs and timestamps to avoid creating noisy diffs
	// NOTE: we can't freeze time stamps as that breaks some tests that use if-modified-since etc (maybe it can be fixed?)
	//testframework.AddHeaderRegexSanitizer("X-Ms-Date", "Wed, 10 Aug 2022 23:34:14 GMT", "", nil)
	require.NoError(t, recording.AddHeaderRegexSanitizer("x-ms-request-id", "00000000-0000-0000-0000-000000000000", "", nil))
	//testframework.AddHeaderRegexSanitizer("Date", "Wed, 10 Aug 2022 23:34:14 GMT", "", nil)
	// TODO: more freezing
	//testframework.AddBodyRegexSanitizer("RequestId:00000000-0000-0000-0000-000000000000", `RequestId:\w{8}-\w{4}-\w{4}-\w{4}-\w{12}`, nil)
	//testframework.AddBodyRegexSanitizer("Time:2022-08-11T00:21:56.4562741Z", `Time:\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d*)?Z`, nil)
	require.NoError(t, recording.Start(t, "sdk/storage/azfile/testdata", nil))
}

func AfterTest(t *testing.T, suite string, test string) {
	require.NoError(t, recording.Stop(t, nil))
}
