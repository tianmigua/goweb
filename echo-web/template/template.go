package template

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

// template_40x_tmpl reads file data from disk. It returns an error on failure.
func template_40x_tmpl() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\40x.tmpl",
		"template/40x.tmpl",
	)
}

// template_50x_tmpl reads file data from disk. It returns an error on failure.
func template_50x_tmpl() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\50x.tmpl",
		"template/50x.tmpl",
	)
}

// template_include_base_html reads file data from disk. It returns an error on failure.
func template_include_base_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\include\\base.html",
		"template/include/base.html",
	)
}

// template_layout_tmpl reads file data from disk. It returns an error on failure.
func template_layout_tmpl() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\layout.tmpl",
		"template/layout.tmpl",
	)
}

// template_layout_web_about_html reads file data from disk. It returns an error on failure.
func template_layout_web_about_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\layout\\web\\about.html",
		"template/layout/web/about.html",
	)
}

// template_layout_web_home_html reads file data from disk. It returns an error on failure.
func template_layout_web_home_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\layout\\web\\home.html",
		"template/layout/web/home.html",
	)
}

// template_layout_web_index_html reads file data from disk. It returns an error on failure.
func template_layout_web_index_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\layout\\web\\index.html",
		"template/layout/web/index.html",
	)
}

// template_layout_web_loginbak_html reads file data from disk. It returns an error on failure.
func template_layout_web_loginbak_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\layout\\web\\loginbak.html",
		"template/layout/web/loginbak.html",
	)
}

// template_layout_web_register_html reads file data from disk. It returns an error on failure.
func template_layout_web_register_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\layout\\web\\register.html",
		"template/layout/web/register.html",
	)
}

// template_layout_web_user_html reads file data from disk. It returns an error on failure.
func template_layout_web_user_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\layout\\web\\user.html",
		"template/layout/web/user.html",
	)
}

// template_pongo2_base_html reads file data from disk. It returns an error on failure.
func template_pongo2_base_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\base.html",
		"template/pongo2/base.html",
	)
}

// template_pongo2_basebak_html reads file data from disk. It returns an error on failure.
func template_pongo2_basebak_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\basebak.html",
		"template/pongo2/basebak.html",
	)
}

// template_pongo2_header2_html reads file data from disk. It returns an error on failure.
func template_pongo2_header2_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\header2.html",
		"template/pongo2/header2.html",
	)
}

// template_pongo2_web_about_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_about_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\about.html",
		"template/pongo2/web/about.html",
	)
}

// template_pongo2_web_add_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_add_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\add.html",
		"template/pongo2/web/add.html",
	)
}

// template_pongo2_web_dashboard_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_dashboard_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\dashboard.html",
		"template/pongo2/web/dashboard.html",
	)
}

// template_pongo2_web_edit_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_edit_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\edit.html",
		"template/pongo2/web/edit.html",
	)
}

// template_pongo2_web_edit1_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_edit1_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\edit1.html",
		"template/pongo2/web/edit1.html",
	)
}

// template_pongo2_web_home_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_home_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\home.html",
		"template/pongo2/web/home.html",
	)
}

// template_pongo2_web_homebak_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_homebak_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\homebak.html",
		"template/pongo2/web/homebak.html",
	)
}

// template_pongo2_web_index_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_index_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\index.html",
		"template/pongo2/web/index.html",
	)
}

// template_pongo2_web_jwt_tester_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_jwt_tester_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\jwt_tester.html",
		"template/pongo2/web/jwt_tester.html",
	)
}

// template_pongo2_web_login_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_login_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\login.html",
		"template/pongo2/web/login.html",
	)
}

// template_pongo2_web_login1_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_login1_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\login1.html",
		"template/pongo2/web/login1.html",
	)
}

// template_pongo2_web_register_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_register_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\register.html",
		"template/pongo2/web/register.html",
	)
}

// template_pongo2_web_user_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_user_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\user.html",
		"template/pongo2/web/user.html",
	)
}

// template_pongo2_web_user_add_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_user_add_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\user\\add.html",
		"template/pongo2/web/user/add.html",
	)
}

// template_pongo2_web_ws_html reads file data from disk. It returns an error on failure.
func template_pongo2_web_ws_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\template\\pongo2\\web\\ws.html",
		"template/pongo2/web/ws.html",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"template/40x.tmpl": template_40x_tmpl,
	"template/50x.tmpl": template_50x_tmpl,
	"template/include/base.html": template_include_base_html,
	"template/layout.tmpl": template_layout_tmpl,
	"template/layout/web/about.html": template_layout_web_about_html,
	"template/layout/web/home.html": template_layout_web_home_html,
	"template/layout/web/index.html": template_layout_web_index_html,
	"template/layout/web/loginbak.html": template_layout_web_loginbak_html,
	"template/layout/web/register.html": template_layout_web_register_html,
	"template/layout/web/user.html": template_layout_web_user_html,
	"template/pongo2/base.html": template_pongo2_base_html,
	"template/pongo2/basebak.html": template_pongo2_basebak_html,
	"template/pongo2/header2.html": template_pongo2_header2_html,
	"template/pongo2/web/about.html": template_pongo2_web_about_html,
	"template/pongo2/web/add.html": template_pongo2_web_add_html,
	"template/pongo2/web/dashboard.html": template_pongo2_web_dashboard_html,
	"template/pongo2/web/edit.html": template_pongo2_web_edit_html,
	"template/pongo2/web/edit1.html": template_pongo2_web_edit1_html,
	"template/pongo2/web/home.html": template_pongo2_web_home_html,
	"template/pongo2/web/homebak.html": template_pongo2_web_homebak_html,
	"template/pongo2/web/index.html": template_pongo2_web_index_html,
	"template/pongo2/web/jwt_tester.html": template_pongo2_web_jwt_tester_html,
	"template/pongo2/web/login.html": template_pongo2_web_login_html,
	"template/pongo2/web/login1.html": template_pongo2_web_login1_html,
	"template/pongo2/web/register.html": template_pongo2_web_register_html,
	"template/pongo2/web/user.html": template_pongo2_web_user_html,
	"template/pongo2/web/user/add.html": template_pongo2_web_user_add_html,
	"template/pongo2/web/ws.html": template_pongo2_web_ws_html,
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
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"template/40x.tmpl": &_bintree_t{template_40x_tmpl, map[string]*_bintree_t{
	}},
	"template/50x.tmpl": &_bintree_t{template_50x_tmpl, map[string]*_bintree_t{
	}},
	"template/include/base.html": &_bintree_t{template_include_base_html, map[string]*_bintree_t{
	}},
	"template/layout.tmpl": &_bintree_t{template_layout_tmpl, map[string]*_bintree_t{
	}},
	"template/layout/web/about.html": &_bintree_t{template_layout_web_about_html, map[string]*_bintree_t{
	}},
	"template/layout/web/home.html": &_bintree_t{template_layout_web_home_html, map[string]*_bintree_t{
	}},
	"template/layout/web/index.html": &_bintree_t{template_layout_web_index_html, map[string]*_bintree_t{
	}},
	"template/layout/web/loginbak.html": &_bintree_t{template_layout_web_loginbak_html, map[string]*_bintree_t{
	}},
	"template/layout/web/register.html": &_bintree_t{template_layout_web_register_html, map[string]*_bintree_t{
	}},
	"template/layout/web/user.html": &_bintree_t{template_layout_web_user_html, map[string]*_bintree_t{
	}},
	"template/pongo2/base.html": &_bintree_t{template_pongo2_base_html, map[string]*_bintree_t{
	}},
	"template/pongo2/basebak.html": &_bintree_t{template_pongo2_basebak_html, map[string]*_bintree_t{
	}},
	"template/pongo2/header2.html": &_bintree_t{template_pongo2_header2_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/about.html": &_bintree_t{template_pongo2_web_about_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/add.html": &_bintree_t{template_pongo2_web_add_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/dashboard.html": &_bintree_t{template_pongo2_web_dashboard_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/edit.html": &_bintree_t{template_pongo2_web_edit_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/edit1.html": &_bintree_t{template_pongo2_web_edit1_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/home.html": &_bintree_t{template_pongo2_web_home_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/homebak.html": &_bintree_t{template_pongo2_web_homebak_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/index.html": &_bintree_t{template_pongo2_web_index_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/jwt_tester.html": &_bintree_t{template_pongo2_web_jwt_tester_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/login.html": &_bintree_t{template_pongo2_web_login_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/login1.html": &_bintree_t{template_pongo2_web_login1_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/register.html": &_bintree_t{template_pongo2_web_register_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/user.html": &_bintree_t{template_pongo2_web_user_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/user/add.html": &_bintree_t{template_pongo2_web_user_add_html, map[string]*_bintree_t{
	}},
	"template/pongo2/web/ws.html": &_bintree_t{template_pongo2_web_ws_html, map[string]*_bintree_t{
	}},
}}
