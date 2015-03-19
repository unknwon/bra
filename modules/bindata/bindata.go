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

var _templates_default_bra_toml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x8f\x41\x4b\xc4\x30\x10\x85\xcf\xcd\xaf\x18\xe2\x55\xa5\x1e\x7a\x11\x3c\xa9\x47\x6f\x82\x87\x12\x4a\xb6\x19\xea\x40\x3a\x81\xc9\x74\xdd\xfd\xf7\x4e\xc3\x82\x22\xb2\x39\x85\x97\xef\xbd\xbc\x37\xca\xc6\xc1\x11\x93\x4e\xf3\x9a\x2a\x3c\xc1\xe8\xba\xd1\x2f\xc5\xdf\x82\x27\xae\x1a\x73\xf6\xc1\x85\xae\x9d\x1b\x78\x2e\xeb\x1a\xd9\x40\xf3\x01\x31\x18\x20\xea\xbe\xa2\xce\x9f\x93\xa1\xe6\x57\xd9\x70\x27\x3f\x76\x0d\x76\xad\x6e\x87\xbb\x44\x82\xb3\x16\x21\xac\x17\xda\x94\xf6\x5d\x68\xb9\x2f\x3f\xef\xa0\x05\x1a\x72\x01\xf1\xa4\x0d\xf4\xf7\xd6\x2a\x18\xfb\x7a\x52\xe4\x4a\x85\x7f\xa1\x87\x8d\x72\x9a\x12\xe6\x78\x36\xf6\x61\xe8\xfb\x3d\xf5\xcd\x86\xad\x31\x5b\x51\x45\x39\xda\xc5\xf8\x77\xa1\x65\x41\x81\xe6\x00\x3c\x22\xab\xbb\x32\xbd\xfb\x67\xbb\xa5\xd8\x7c\xe7\xc6\x7a\xe6\x39\xb8\x4c\xd5\x0a\x4d\x31\x25\xb1\x10\xff\x38\xf4\x43\xef\x9d\xe0\x5a\x14\xff\xa8\xdf\x01\x00\x00\xff\xff\x05\xa0\x75\x9b\x6f\x01\x00\x00")

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

	info := bindata_file_info{name: "templates/default.bra.toml", size: 367, mode: os.FileMode(436), modTime: time.Unix(1426746442, 0)}
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

