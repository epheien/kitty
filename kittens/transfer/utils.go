// License: GPLv3 Copyright: 2023, Kovid Goyal, <kovid at kovidgoyal.net>

package transfer

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"kitty/tools/utils"
)

var _ = fmt.Print

var global_cwd, global_home string

func cwd_path() string {
	if global_cwd == "" {
		ans, _ := os.Getwd()
		return ans
	}
	return global_cwd
}

func home_path() string {
	if global_home == "" {
		return utils.Expanduser("~")
	}
	return global_home
}

func abspath(path string, use_home ...bool) string {
	var base string
	if len(use_home) > 0 && use_home[0] {
		base = home_path()
	} else {
		base = cwd_path()
	}
	return filepath.Join(base, path)
}

func expand_home(path string) string {
	if strings.HasPrefix(path, "~"+string(os.PathSeparator)) {
		path = strings.TrimLeft(path[2:], string(os.PathSeparator))
		path = filepath.Join(home_path(), path)
	} else if path == "~" {
		path = home_path()
	}
	return path
}

func random_id() string {
	bytes := []byte{0, 0}
	rand.Read(bytes)
	return fmt.Sprintf("%x%s", os.Getpid(), hex.EncodeToString(bytes))
}

func run_with_paths(cwd, home string, f func()) {
	global_cwd, global_home = cwd, home
	defer func() { global_cwd, global_home = "", "" }()
	f()
}