// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build linux
// +build linux

package platform

import (
	"fmt"
	"os"
	"sync"
)

const ipv6SupportProcFile = "/proc/net/if_inet6"

var (
	ipv6SupportOnce  sync.Once
	ipv6Supported    bool
	ipv6SupportedErr error
)

func SupportsIPv6() (bool, error) {
	ipv6SupportOnce.Do(func() {
		ipv6Supported, ipv6SupportedErr = checkIfKernelSupportsIPv6()
	})
	return ipv6Supported, ipv6SupportedErr
}

func checkIfKernelSupportsIPv6() (bool, error) {
	_, err := os.Stat(ipv6SupportProcFile)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("error checking for ipv6 support file %s: %w", ipv6SupportProcFile, err)
	}

	return true, nil
}
