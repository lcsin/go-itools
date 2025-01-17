package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Exists 判断文件或目录是否存在
func Exists(src string) bool {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return false
	}
	return true
}

// CopyFile 拷贝文件
func CopyFile(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	buf, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, buf, info.Mode())
}

// CopyDir 拷贝目录
func CopyDir(src, dst string, ignores []string) error {
	srcinfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dst, srcinfo.Mode())
	if err != nil {
		return err
	}

	fds, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	hasIgnores := func(name string, ignores []string) bool {
		for _, v := range ignores {
			if v == name {
				return true
			}
		}
		return false
	}

	for _, fd := range fds {
		if hasIgnores(fd.Name(), ignores) {
			continue
		}
		srcfp := filepath.Join(src, fd.Name())
		dstfp := filepath.Join(dst, fd.Name())
		if fd.IsDir() {
			err = CopyDir(srcfp, dstfp, ignores)
		} else {
			err = CopyFile(srcfp, dstfp)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// FileMD5 获取文件MD5
func FileMD5(fp string) (string, error) {
	file, err := os.Open(fp)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil)[:16])), nil
}
