package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/yyle88/erero"
)

func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func IsRootExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // 路径不存在
		}
		return false, erero.Wro(err) // 其他的错误
	}
	return info.IsDir(), nil
}

func IsFileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // 路径不存在
		}
		return false, erero.Wro(err) // 其他的错误
	}
	return !info.IsDir(), nil
}

func LsMapName2Path(root string) (map[string]string, error) {
	names, err := Ls(root)
	if err != nil {
		return nil, erero.WithMessage(err, "wrong")
	}
	var mp = make(map[string]string, len(names))
	for _, name := range names {
		mp[name] = filepath.Join(root, name)
	}
	return mp, nil
}

func Ls(root string) (names []string, err error) {
	infos, err := os.ReadDir(root)
	if err != nil {
		return nil, erero.WithMessage(err, "wrong")
	}
	names = make([]string, 0, len(infos))
	for _, info := range infos {
		names = append(names, info.Name())
	}
	return names, nil
}

func HasAnySuffix(s string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}
