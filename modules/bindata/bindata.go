package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _templates_default_bra_toml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x90\xb1\x6a\xc3\x30\x10\x86\xe7\xe8\x29\x0e\xb5\x63\x9a\xba\x43\x96\x42\x87\xd0\x66\x4c\xc9\x50\xe8\x60\x84\x51\xac\xc3\x3d\x90\x25\x38\x9d\xd3\xe4\xed\x7b\x16\x86\x94\xce\xf5\x64\x7e\x7d\xfa\xee\xd7\xb5\x3c\x25\x67\x28\x91\x74\xfd\x18\x0a\xbc\x40\x6b\x56\xad\x1d\xb2\x5d\x83\xa5\x54\xc4\xc7\x68\xdd\xfa\x96\x9d\x26\x8a\x61\x49\x36\x8f\xf7\xbb\xe3\xb1\x7b\xdf\x1d\xf6\xd6\x19\xb7\xaa\xdf\x1d\xbc\xe6\x71\xf4\x49\x65\xea\x06\x4a\xa0\x12\x16\xf3\xed\xa5\xff\xea\x54\xa7\x33\x84\x27\x9c\xc9\xcf\x39\x83\x39\x2b\xd3\xe9\x21\x10\x63\x2f\x99\x09\xcb\x42\x6b\x52\x2b\xb9\xea\x7d\xbb\x9d\x83\x64\xa8\xc8\x02\xe2\x45\x2a\x68\x37\xda\xd2\x29\xbb\xbf\x08\xa6\x42\x39\xfd\x42\x6b\xf5\x2e\x60\xf4\x57\x65\x9f\xb6\x4d\x33\x5b\x0f\xfa\xf8\xd1\x47\x2d\x2a\xc8\x67\xfd\x51\xfe\x83\x69\x18\x90\xa1\xde\x00\x3c\x63\x12\xf3\xff\xeb\xd1\x41\xba\x21\x63\xda\x72\x4d\xbd\x33\x91\x8a\x76\xee\x7c\x08\xac\x73\xec\xf3\xb6\xd9\x36\xd6\x30\x8e\x59\xf0\x4f\xfa\x13\x00\x00\xff\xff\x8a\x14\x8a\x95\xb6\x01\x00\x00")

func templates_default_bra_toml_bytes() ([]byte, error) {
	return bindata_read(
		_templates_default_bra_toml,
		"templates/default.bra.toml",
	)
}

func templates_default_bra_toml() (*asset, error) {
	bytes, err := templates_default_bra_toml_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/default.bra.toml", size: 438, mode: os.FileMode(436), modTime: time.Unix(1426938511, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/default.bra.toml": templates_default_bra_toml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"default.bra.toml": &_bintree_t{templates_default_bra_toml, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

