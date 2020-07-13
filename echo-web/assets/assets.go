package assets

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

// assets_bootstrap_css_bootstrap_min_css reads file data from disk. It returns an error on failure.
func assets_bootstrap_css_bootstrap_min_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\bootstrap\\css\\bootstrap.min.css",
		"assets/bootstrap/css/bootstrap.min.css",
	)
}

// assets_bootstrap_css_bootstrap_min_css_map reads file data from disk. It returns an error on failure.
func assets_bootstrap_css_bootstrap_min_css_map() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\bootstrap\\css\\bootstrap.min.css.map",
		"assets/bootstrap/css/bootstrap.min.css.map",
	)
}

// assets_bootstrap_fonts_glyphicons_halflings_regular_eot reads file data from disk. It returns an error on failure.
func assets_bootstrap_fonts_glyphicons_halflings_regular_eot() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\bootstrap\\fonts\\glyphicons-halflings-regular.eot",
		"assets/bootstrap/fonts/glyphicons-halflings-regular.eot",
	)
}

// assets_bootstrap_fonts_glyphicons_halflings_regular_svg reads file data from disk. It returns an error on failure.
func assets_bootstrap_fonts_glyphicons_halflings_regular_svg() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\bootstrap\\fonts\\glyphicons-halflings-regular.svg",
		"assets/bootstrap/fonts/glyphicons-halflings-regular.svg",
	)
}

// assets_bootstrap_fonts_glyphicons_halflings_regular_ttf reads file data from disk. It returns an error on failure.
func assets_bootstrap_fonts_glyphicons_halflings_regular_ttf() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\bootstrap\\fonts\\glyphicons-halflings-regular.ttf",
		"assets/bootstrap/fonts/glyphicons-halflings-regular.ttf",
	)
}

// assets_bootstrap_fonts_glyphicons_halflings_regular_woff reads file data from disk. It returns an error on failure.
func assets_bootstrap_fonts_glyphicons_halflings_regular_woff() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\bootstrap\\fonts\\glyphicons-halflings-regular.woff",
		"assets/bootstrap/fonts/glyphicons-halflings-regular.woff",
	)
}

// assets_bootstrap_fonts_glyphicons_halflings_regular_woff2 reads file data from disk. It returns an error on failure.
func assets_bootstrap_fonts_glyphicons_halflings_regular_woff2() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\bootstrap\\fonts\\glyphicons-halflings-regular.woff2",
		"assets/bootstrap/fonts/glyphicons-halflings-regular.woff2",
	)
}

// assets_bootstrap_js_bootstrap_min_js reads file data from disk. It returns an error on failure.
func assets_bootstrap_js_bootstrap_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\bootstrap\\js\\bootstrap.min.js",
		"assets/bootstrap/js/bootstrap.min.js",
	)
}

// assets_css_admin_login_css reads file data from disk. It returns an error on failure.
func assets_css_admin_login_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\css\\admin-login.css",
		"assets/css/admin-login.css",
	)
}

// assets_css_admin_css reads file data from disk. It returns an error on failure.
func assets_css_admin_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\css\\admin.css",
		"assets/css/admin.css",
	)
}

// assets_css_animate_css reads file data from disk. It returns an error on failure.
func assets_css_animate_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\css\\animate.css",
		"assets/css/animate.css",
	)
}

// assets_css_bootstrap_combined_min_css reads file data from disk. It returns an error on failure.
func assets_css_bootstrap_combined_min_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\css\\bootstrap-combined.min.css",
		"assets/css/bootstrap-combined.min.css",
	)
}

// assets_css_bootstrap_theme_css reads file data from disk. It returns an error on failure.
func assets_css_bootstrap_theme_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\css\\bootstrap-theme.css",
		"assets/css/bootstrap-theme.css",
	)
}

// assets_css_bootstrap_theme_css_map reads file data from disk. It returns an error on failure.
func assets_css_bootstrap_theme_css_map() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\css\\bootstrap-theme.css.map",
		"assets/css/bootstrap-theme.css.map",
	)
}

// assets_css_bootstrap_theme_min_css reads file data from disk. It returns an error on failure.
func assets_css_bootstrap_theme_min_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\css\\bootstrap-theme.min.css",
		"assets/css/bootstrap-theme.min.css",
	)
}

// assets_css_bootstrap_css reads file data from disk. It returns an error on failure.
func assets_css_bootstrap_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\css\\bootstrap.css",
		"assets/css/bootstrap.css",
	)
}

// assets_css_bootstrap_css_map reads file data from disk. It returns an error on failure.
func assets_css_bootstrap_css_map() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\css\\bootstrap.css.map",
		"assets/css/bootstrap.css.map",
	)
}

// assets_css_bootstrap_min_css reads file data from disk. It returns an error on failure.
func assets_css_bootstrap_min_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\css\\bootstrap.min.css",
		"assets/css/bootstrap.min.css",
	)
}

// assets_css_jquery_treetable_css reads file data from disk. It returns an error on failure.
func assets_css_jquery_treetable_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\css\\jquery.treetable.css",
		"assets/css/jquery.treetable.css",
	)
}

// assets_css_jquery_treetable_theme_default_css reads file data from disk. It returns an error on failure.
func assets_css_jquery_treetable_theme_default_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\css\\jquery.treetable.theme.default.css",
		"assets/css/jquery.treetable.theme.default.css",
	)
}

// assets_font_awesome_4_7_0_help_us_out_txt reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_help_us_out_txt() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\HELP-US-OUT.txt",
		"assets/font-awesome-4.7.0/HELP-US-OUT.txt",
	)
}

// assets_font_awesome_4_7_0_css_font_awesome_css reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_css_font_awesome_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\css\\font-awesome.css",
		"assets/font-awesome-4.7.0/css/font-awesome.css",
	)
}

// assets_font_awesome_4_7_0_css_font_awesome_min_css reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_css_font_awesome_min_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\css\\font-awesome.min.css",
		"assets/font-awesome-4.7.0/css/font-awesome.min.css",
	)
}

// assets_font_awesome_4_7_0_fonts_fontawesome_otf reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_fonts_fontawesome_otf() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\fonts\\FontAwesome.otf",
		"assets/font-awesome-4.7.0/fonts/FontAwesome.otf",
	)
}

// assets_font_awesome_4_7_0_fonts_fontawesome_webfont_eot reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_fonts_fontawesome_webfont_eot() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\fonts\\fontawesome-webfont.eot",
		"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.eot",
	)
}

// assets_font_awesome_4_7_0_fonts_fontawesome_webfont_svg reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_fonts_fontawesome_webfont_svg() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\fonts\\fontawesome-webfont.svg",
		"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.svg",
	)
}

// assets_font_awesome_4_7_0_fonts_fontawesome_webfont_ttf reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_fonts_fontawesome_webfont_ttf() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\fonts\\fontawesome-webfont.ttf",
		"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.ttf",
	)
}

// assets_font_awesome_4_7_0_fonts_fontawesome_webfont_woff reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_fonts_fontawesome_webfont_woff() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\fonts\\fontawesome-webfont.woff",
		"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.woff",
	)
}

// assets_font_awesome_4_7_0_fonts_fontawesome_webfont_woff2 reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_fonts_fontawesome_webfont_woff2() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\fonts\\fontawesome-webfont.woff2",
		"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.woff2",
	)
}

// assets_font_awesome_4_7_0_less_animated_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_animated_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\animated.less",
		"assets/font-awesome-4.7.0/less/animated.less",
	)
}

// assets_font_awesome_4_7_0_less_bordered_pulled_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_bordered_pulled_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\bordered-pulled.less",
		"assets/font-awesome-4.7.0/less/bordered-pulled.less",
	)
}

// assets_font_awesome_4_7_0_less_core_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_core_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\core.less",
		"assets/font-awesome-4.7.0/less/core.less",
	)
}

// assets_font_awesome_4_7_0_less_fixed_width_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_fixed_width_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\fixed-width.less",
		"assets/font-awesome-4.7.0/less/fixed-width.less",
	)
}

// assets_font_awesome_4_7_0_less_font_awesome_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_font_awesome_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\font-awesome.less",
		"assets/font-awesome-4.7.0/less/font-awesome.less",
	)
}

// assets_font_awesome_4_7_0_less_icons_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_icons_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\icons.less",
		"assets/font-awesome-4.7.0/less/icons.less",
	)
}

// assets_font_awesome_4_7_0_less_larger_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_larger_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\larger.less",
		"assets/font-awesome-4.7.0/less/larger.less",
	)
}

// assets_font_awesome_4_7_0_less_list_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_list_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\list.less",
		"assets/font-awesome-4.7.0/less/list.less",
	)
}

// assets_font_awesome_4_7_0_less_mixins_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_mixins_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\mixins.less",
		"assets/font-awesome-4.7.0/less/mixins.less",
	)
}

// assets_font_awesome_4_7_0_less_path_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_path_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\path.less",
		"assets/font-awesome-4.7.0/less/path.less",
	)
}

// assets_font_awesome_4_7_0_less_rotated_flipped_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_rotated_flipped_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\rotated-flipped.less",
		"assets/font-awesome-4.7.0/less/rotated-flipped.less",
	)
}

// assets_font_awesome_4_7_0_less_screen_reader_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_screen_reader_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\screen-reader.less",
		"assets/font-awesome-4.7.0/less/screen-reader.less",
	)
}

// assets_font_awesome_4_7_0_less_stacked_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_stacked_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\stacked.less",
		"assets/font-awesome-4.7.0/less/stacked.less",
	)
}

// assets_font_awesome_4_7_0_less_variables_less reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_less_variables_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\less\\variables.less",
		"assets/font-awesome-4.7.0/less/variables.less",
	)
}

// assets_font_awesome_4_7_0_scss_animated_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_animated_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\_animated.scss",
		"assets/font-awesome-4.7.0/scss/_animated.scss",
	)
}

// assets_font_awesome_4_7_0_scss_bordered_pulled_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_bordered_pulled_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\_bordered-pulled.scss",
		"assets/font-awesome-4.7.0/scss/_bordered-pulled.scss",
	)
}

// assets_font_awesome_4_7_0_scss_core_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_core_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\_core.scss",
		"assets/font-awesome-4.7.0/scss/_core.scss",
	)
}

// assets_font_awesome_4_7_0_scss_fixed_width_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_fixed_width_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\_fixed-width.scss",
		"assets/font-awesome-4.7.0/scss/_fixed-width.scss",
	)
}

// assets_font_awesome_4_7_0_scss_icons_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_icons_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\_icons.scss",
		"assets/font-awesome-4.7.0/scss/_icons.scss",
	)
}

// assets_font_awesome_4_7_0_scss_larger_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_larger_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\_larger.scss",
		"assets/font-awesome-4.7.0/scss/_larger.scss",
	)
}

// assets_font_awesome_4_7_0_scss_list_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_list_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\_list.scss",
		"assets/font-awesome-4.7.0/scss/_list.scss",
	)
}

// assets_font_awesome_4_7_0_scss_mixins_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_mixins_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\_mixins.scss",
		"assets/font-awesome-4.7.0/scss/_mixins.scss",
	)
}

// assets_font_awesome_4_7_0_scss_path_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_path_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\_path.scss",
		"assets/font-awesome-4.7.0/scss/_path.scss",
	)
}

// assets_font_awesome_4_7_0_scss_rotated_flipped_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_rotated_flipped_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\_rotated-flipped.scss",
		"assets/font-awesome-4.7.0/scss/_rotated-flipped.scss",
	)
}

// assets_font_awesome_4_7_0_scss_screen_reader_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_screen_reader_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\_screen-reader.scss",
		"assets/font-awesome-4.7.0/scss/_screen-reader.scss",
	)
}

// assets_font_awesome_4_7_0_scss_stacked_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_stacked_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\_stacked.scss",
		"assets/font-awesome-4.7.0/scss/_stacked.scss",
	)
}

// assets_font_awesome_4_7_0_scss_variables_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_variables_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\_variables.scss",
		"assets/font-awesome-4.7.0/scss/_variables.scss",
	)
}

// assets_font_awesome_4_7_0_scss_font_awesome_scss reads file data from disk. It returns an error on failure.
func assets_font_awesome_4_7_0_scss_font_awesome_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\font-awesome-4.7.0\\scss\\font-awesome.scss",
		"assets/font-awesome-4.7.0/scss/font-awesome.scss",
	)
}

// assets_fonts_glyphicons_halflings_regular_eot reads file data from disk. It returns an error on failure.
func assets_fonts_glyphicons_halflings_regular_eot() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\fonts\\glyphicons-halflings-regular.eot",
		"assets/fonts/glyphicons-halflings-regular.eot",
	)
}

// assets_fonts_glyphicons_halflings_regular_svg reads file data from disk. It returns an error on failure.
func assets_fonts_glyphicons_halflings_regular_svg() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\fonts\\glyphicons-halflings-regular.svg",
		"assets/fonts/glyphicons-halflings-regular.svg",
	)
}

// assets_fonts_glyphicons_halflings_regular_ttf reads file data from disk. It returns an error on failure.
func assets_fonts_glyphicons_halflings_regular_ttf() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\fonts\\glyphicons-halflings-regular.ttf",
		"assets/fonts/glyphicons-halflings-regular.ttf",
	)
}

// assets_fonts_glyphicons_halflings_regular_woff reads file data from disk. It returns an error on failure.
func assets_fonts_glyphicons_halflings_regular_woff() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\fonts\\glyphicons-halflings-regular.woff",
		"assets/fonts/glyphicons-halflings-regular.woff",
	)
}

// assets_fonts_glyphicons_halflings_regular_woff2 reads file data from disk. It returns an error on failure.
func assets_fonts_glyphicons_halflings_regular_woff2() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\fonts\\glyphicons-halflings-regular.woff2",
		"assets/fonts/glyphicons-halflings-regular.woff2",
	)
}

// assets_images_bg_2_jpg reads file data from disk. It returns an error on failure.
func assets_images_bg_2_jpg() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\images\\bg_2.jpg",
		"assets/images/bg_2.jpg",
	)
}

// assets_images_chahao_png reads file data from disk. It returns an error on failure.
func assets_images_chahao_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\images\\chahao.png",
		"assets/images/chahao.png",
	)
}

// assets_images_duihao_png reads file data from disk. It returns an error on failure.
func assets_images_duihao_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\images\\duihao.png",
		"assets/images/duihao.png",
	)
}

// assets_images_geometry2_png reads file data from disk. It returns an error on failure.
func assets_images_geometry2_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\images\\geometry2.png",
		"assets/images/geometry2.png",
	)
}

// assets_images_hevr_png reads file data from disk. It returns an error on failure.
func assets_images_hevr_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\images\\hevr.png",
		"assets/images/hevr.png",
	)
}

// assets_images_logo_png reads file data from disk. It returns an error on failure.
func assets_images_logo_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\images\\logo.png",
		"assets/images/logo.png",
	)
}

// assets_images_qrcode_jpg reads file data from disk. It returns an error on failure.
func assets_images_qrcode_jpg() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\images\\qrcode.jpg",
		"assets/images/qrcode.jpg",
	)
}

// assets_images_success_png reads file data from disk. It returns an error on failure.
func assets_images_success_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\images\\success.png",
		"assets/images/success.png",
	)
}

// assets_images_upload_jpg reads file data from disk. It returns an error on failure.
func assets_images_upload_jpg() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\images\\upload.jpg",
		"assets/images/upload.jpg",
	)
}

// assets_img_favicon_ico reads file data from disk. It returns an error on failure.
func assets_img_favicon_ico() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\img\\favicon.ico",
		"assets/img/favicon.ico",
	)
}

// assets_img_icon_brand_png reads file data from disk. It returns an error on failure.
func assets_img_icon_brand_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\img\\icon_brand.png",
		"assets/img/icon_brand.png",
	)
}

// assets_js_admin_login_js reads file data from disk. It returns an error on failure.
func assets_js_admin_login_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\admin-login.js",
		"assets/js/admin-login.js",
	)
}

// assets_js_admin_js reads file data from disk. It returns an error on failure.
func assets_js_admin_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\admin.js",
		"assets/js/admin.js",
	)
}

// assets_js_app_jwt_jwt_js reads file data from disk. It returns an error on failure.
func assets_js_app_jwt_jwt_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\app\\jwt\\jwt.js",
		"assets/js/app/jwt/jwt.js",
	)
}

// assets_js_app_treetable_treetable_page_js reads file data from disk. It returns an error on failure.
func assets_js_app_treetable_treetable_page_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\app\\treetable\\treetable-page.js",
		"assets/js/app/treetable/treetable-page.js",
	)
}

// assets_js_common_js reads file data from disk. It returns an error on failure.
func assets_js_common_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\common.js",
		"assets/js/common.js",
	)
}

// assets_js_echarts_min_js reads file data from disk. It returns an error on failure.
func assets_js_echarts_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\echarts.min.js",
		"assets/js/echarts.min.js",
	)
}

// assets_js_jquery_min_js reads file data from disk. It returns an error on failure.
func assets_js_jquery_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\jquery.min.js",
		"assets/js/jquery.min.js",
	)
}

// assets_js_jquery_placeholder_min_js reads file data from disk. It returns an error on failure.
func assets_js_jquery_placeholder_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\jquery.placeholder.min.js",
		"assets/js/jquery.placeholder.min.js",
	)
}

// assets_js_jquery_waypoints_min_js reads file data from disk. It returns an error on failure.
func assets_js_jquery_waypoints_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\jquery.waypoints.min.js",
		"assets/js/jquery.waypoints.min.js",
	)
}

// assets_js_lib_bootstrap_bootstrap_js reads file data from disk. It returns an error on failure.
func assets_js_lib_bootstrap_bootstrap_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\lib\\bootstrap\\bootstrap.js",
		"assets/js/lib/bootstrap/bootstrap.js",
	)
}

// assets_js_lib_bootstrap_bootstrap_min_js reads file data from disk. It returns an error on failure.
func assets_js_lib_bootstrap_bootstrap_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\lib\\bootstrap\\bootstrap.min.js",
		"assets/js/lib/bootstrap/bootstrap.min.js",
	)
}

// assets_js_lib_bootstrap_npm_js reads file data from disk. It returns an error on failure.
func assets_js_lib_bootstrap_npm_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\lib\\bootstrap\\npm.js",
		"assets/js/lib/bootstrap/npm.js",
	)
}

// assets_js_lib_jquery_form_jquery_form_3_51_js reads file data from disk. It returns an error on failure.
func assets_js_lib_jquery_form_jquery_form_3_51_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\lib\\jquery-form\\jquery.form.3.51.js",
		"assets/js/lib/jquery-form/jquery.form.3.51.js",
	)
}

// assets_js_lib_jquery_treetable_jquery_treetable_js reads file data from disk. It returns an error on failure.
func assets_js_lib_jquery_treetable_jquery_treetable_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\lib\\jquery-treetable\\jquery.treetable.js",
		"assets/js/lib/jquery-treetable/jquery.treetable.js",
	)
}

// assets_js_lib_jquery_jquery_2_1_3_js reads file data from disk. It returns an error on failure.
func assets_js_lib_jquery_jquery_2_1_3_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\lib\\jquery\\jquery-2.1.3.js",
		"assets/js/lib/jquery/jquery-2.1.3.js",
	)
}

// assets_js_lib_jquery_jquery_2_1_3_min_js reads file data from disk. It returns an error on failure.
func assets_js_lib_jquery_jquery_2_1_3_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\lib\\jquery\\jquery-2.1.3.min.js",
		"assets/js/lib/jquery/jquery-2.1.3.min.js",
	)
}

// assets_js_lib_require_setup_js reads file data from disk. It returns an error on failure.
func assets_js_lib_require_setup_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\lib\\require-setup.js",
		"assets/js/lib/require-setup.js",
	)
}

// assets_js_lib_require_js reads file data from disk. It returns an error on failure.
func assets_js_lib_require_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\lib\\require.js",
		"assets/js/lib/require.js",
	)
}

// assets_js_modernizr_2_6_2_min_js reads file data from disk. It returns an error on failure.
func assets_js_modernizr_2_6_2_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\modernizr-2.6.2.min.js",
		"assets/js/modernizr-2.6.2.min.js",
	)
}

// assets_js_modernizr_custom_v2_7_1_min_js reads file data from disk. It returns an error on failure.
func assets_js_modernizr_custom_v2_7_1_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\modernizr-custom-v2.7.1.min.js",
		"assets/js/modernizr-custom-v2.7.1.min.js",
	)
}

// assets_js_tools_r_js reads file data from disk. It returns an error on failure.
func assets_js_tools_r_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\js\\tools\\r.js",
		"assets/js/tools/r.js",
	)
}

// assets_layui_css_layui_css reads file data from disk. It returns an error on failure.
func assets_layui_css_layui_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\css\\layui.css",
		"assets/layui/css/layui.css",
	)
}

// assets_layui_css_layui_mobile_css reads file data from disk. It returns an error on failure.
func assets_layui_css_layui_mobile_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\css\\layui.mobile.css",
		"assets/layui/css/layui.mobile.css",
	)
}

// assets_layui_css_modules_code_css reads file data from disk. It returns an error on failure.
func assets_layui_css_modules_code_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\css\\modules\\code.css",
		"assets/layui/css/modules/code.css",
	)
}

// assets_layui_css_modules_laydate_default_laydate_css reads file data from disk. It returns an error on failure.
func assets_layui_css_modules_laydate_default_laydate_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\css\\modules\\laydate\\default\\laydate.css",
		"assets/layui/css/modules/laydate/default/laydate.css",
	)
}

// assets_layui_css_modules_layer_default_icon_ext_png reads file data from disk. It returns an error on failure.
func assets_layui_css_modules_layer_default_icon_ext_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\css\\modules\\layer\\default\\icon-ext.png",
		"assets/layui/css/modules/layer/default/icon-ext.png",
	)
}

// assets_layui_css_modules_layer_default_icon_png reads file data from disk. It returns an error on failure.
func assets_layui_css_modules_layer_default_icon_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\css\\modules\\layer\\default\\icon.png",
		"assets/layui/css/modules/layer/default/icon.png",
	)
}

// assets_layui_css_modules_layer_default_layer_css reads file data from disk. It returns an error on failure.
func assets_layui_css_modules_layer_default_layer_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\css\\modules\\layer\\default\\layer.css",
		"assets/layui/css/modules/layer/default/layer.css",
	)
}

// assets_layui_css_modules_layer_default_loading_0_gif reads file data from disk. It returns an error on failure.
func assets_layui_css_modules_layer_default_loading_0_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\css\\modules\\layer\\default\\loading-0.gif",
		"assets/layui/css/modules/layer/default/loading-0.gif",
	)
}

// assets_layui_css_modules_layer_default_loading_1_gif reads file data from disk. It returns an error on failure.
func assets_layui_css_modules_layer_default_loading_1_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\css\\modules\\layer\\default\\loading-1.gif",
		"assets/layui/css/modules/layer/default/loading-1.gif",
	)
}

// assets_layui_css_modules_layer_default_loading_2_gif reads file data from disk. It returns an error on failure.
func assets_layui_css_modules_layer_default_loading_2_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\css\\modules\\layer\\default\\loading-2.gif",
		"assets/layui/css/modules/layer/default/loading-2.gif",
	)
}

// assets_layui_font_iconfont_eot reads file data from disk. It returns an error on failure.
func assets_layui_font_iconfont_eot() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\font\\iconfont.eot",
		"assets/layui/font/iconfont.eot",
	)
}

// assets_layui_font_iconfont_svg reads file data from disk. It returns an error on failure.
func assets_layui_font_iconfont_svg() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\font\\iconfont.svg",
		"assets/layui/font/iconfont.svg",
	)
}

// assets_layui_font_iconfont_ttf reads file data from disk. It returns an error on failure.
func assets_layui_font_iconfont_ttf() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\font\\iconfont.ttf",
		"assets/layui/font/iconfont.ttf",
	)
}

// assets_layui_font_iconfont_woff reads file data from disk. It returns an error on failure.
func assets_layui_font_iconfont_woff() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\font\\iconfont.woff",
		"assets/layui/font/iconfont.woff",
	)
}

// assets_layui_font_iconfont_woff2 reads file data from disk. It returns an error on failure.
func assets_layui_font_iconfont_woff2() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\font\\iconfont.woff2",
		"assets/layui/font/iconfont.woff2",
	)
}

// assets_layui_images_face_0_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_0_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\0.gif",
		"assets/layui/images/face/0.gif",
	)
}

// assets_layui_images_face_1_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_1_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\1.gif",
		"assets/layui/images/face/1.gif",
	)
}

// assets_layui_images_face_10_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_10_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\10.gif",
		"assets/layui/images/face/10.gif",
	)
}

// assets_layui_images_face_11_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_11_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\11.gif",
		"assets/layui/images/face/11.gif",
	)
}

// assets_layui_images_face_12_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_12_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\12.gif",
		"assets/layui/images/face/12.gif",
	)
}

// assets_layui_images_face_13_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_13_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\13.gif",
		"assets/layui/images/face/13.gif",
	)
}

// assets_layui_images_face_14_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_14_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\14.gif",
		"assets/layui/images/face/14.gif",
	)
}

// assets_layui_images_face_15_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_15_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\15.gif",
		"assets/layui/images/face/15.gif",
	)
}

// assets_layui_images_face_16_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_16_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\16.gif",
		"assets/layui/images/face/16.gif",
	)
}

// assets_layui_images_face_17_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_17_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\17.gif",
		"assets/layui/images/face/17.gif",
	)
}

// assets_layui_images_face_18_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_18_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\18.gif",
		"assets/layui/images/face/18.gif",
	)
}

// assets_layui_images_face_19_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_19_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\19.gif",
		"assets/layui/images/face/19.gif",
	)
}

// assets_layui_images_face_2_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_2_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\2.gif",
		"assets/layui/images/face/2.gif",
	)
}

// assets_layui_images_face_20_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_20_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\20.gif",
		"assets/layui/images/face/20.gif",
	)
}

// assets_layui_images_face_21_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_21_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\21.gif",
		"assets/layui/images/face/21.gif",
	)
}

// assets_layui_images_face_22_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_22_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\22.gif",
		"assets/layui/images/face/22.gif",
	)
}

// assets_layui_images_face_23_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_23_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\23.gif",
		"assets/layui/images/face/23.gif",
	)
}

// assets_layui_images_face_24_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_24_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\24.gif",
		"assets/layui/images/face/24.gif",
	)
}

// assets_layui_images_face_25_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_25_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\25.gif",
		"assets/layui/images/face/25.gif",
	)
}

// assets_layui_images_face_26_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_26_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\26.gif",
		"assets/layui/images/face/26.gif",
	)
}

// assets_layui_images_face_27_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_27_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\27.gif",
		"assets/layui/images/face/27.gif",
	)
}

// assets_layui_images_face_28_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_28_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\28.gif",
		"assets/layui/images/face/28.gif",
	)
}

// assets_layui_images_face_29_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_29_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\29.gif",
		"assets/layui/images/face/29.gif",
	)
}

// assets_layui_images_face_3_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_3_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\3.gif",
		"assets/layui/images/face/3.gif",
	)
}

// assets_layui_images_face_30_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_30_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\30.gif",
		"assets/layui/images/face/30.gif",
	)
}

// assets_layui_images_face_31_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_31_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\31.gif",
		"assets/layui/images/face/31.gif",
	)
}

// assets_layui_images_face_32_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_32_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\32.gif",
		"assets/layui/images/face/32.gif",
	)
}

// assets_layui_images_face_33_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_33_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\33.gif",
		"assets/layui/images/face/33.gif",
	)
}

// assets_layui_images_face_34_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_34_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\34.gif",
		"assets/layui/images/face/34.gif",
	)
}

// assets_layui_images_face_35_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_35_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\35.gif",
		"assets/layui/images/face/35.gif",
	)
}

// assets_layui_images_face_36_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_36_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\36.gif",
		"assets/layui/images/face/36.gif",
	)
}

// assets_layui_images_face_37_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_37_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\37.gif",
		"assets/layui/images/face/37.gif",
	)
}

// assets_layui_images_face_38_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_38_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\38.gif",
		"assets/layui/images/face/38.gif",
	)
}

// assets_layui_images_face_39_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_39_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\39.gif",
		"assets/layui/images/face/39.gif",
	)
}

// assets_layui_images_face_4_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_4_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\4.gif",
		"assets/layui/images/face/4.gif",
	)
}

// assets_layui_images_face_40_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_40_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\40.gif",
		"assets/layui/images/face/40.gif",
	)
}

// assets_layui_images_face_41_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_41_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\41.gif",
		"assets/layui/images/face/41.gif",
	)
}

// assets_layui_images_face_42_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_42_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\42.gif",
		"assets/layui/images/face/42.gif",
	)
}

// assets_layui_images_face_43_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_43_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\43.gif",
		"assets/layui/images/face/43.gif",
	)
}

// assets_layui_images_face_44_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_44_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\44.gif",
		"assets/layui/images/face/44.gif",
	)
}

// assets_layui_images_face_45_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_45_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\45.gif",
		"assets/layui/images/face/45.gif",
	)
}

// assets_layui_images_face_46_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_46_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\46.gif",
		"assets/layui/images/face/46.gif",
	)
}

// assets_layui_images_face_47_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_47_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\47.gif",
		"assets/layui/images/face/47.gif",
	)
}

// assets_layui_images_face_48_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_48_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\48.gif",
		"assets/layui/images/face/48.gif",
	)
}

// assets_layui_images_face_49_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_49_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\49.gif",
		"assets/layui/images/face/49.gif",
	)
}

// assets_layui_images_face_5_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_5_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\5.gif",
		"assets/layui/images/face/5.gif",
	)
}

// assets_layui_images_face_50_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_50_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\50.gif",
		"assets/layui/images/face/50.gif",
	)
}

// assets_layui_images_face_51_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_51_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\51.gif",
		"assets/layui/images/face/51.gif",
	)
}

// assets_layui_images_face_52_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_52_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\52.gif",
		"assets/layui/images/face/52.gif",
	)
}

// assets_layui_images_face_53_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_53_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\53.gif",
		"assets/layui/images/face/53.gif",
	)
}

// assets_layui_images_face_54_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_54_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\54.gif",
		"assets/layui/images/face/54.gif",
	)
}

// assets_layui_images_face_55_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_55_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\55.gif",
		"assets/layui/images/face/55.gif",
	)
}

// assets_layui_images_face_56_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_56_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\56.gif",
		"assets/layui/images/face/56.gif",
	)
}

// assets_layui_images_face_57_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_57_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\57.gif",
		"assets/layui/images/face/57.gif",
	)
}

// assets_layui_images_face_58_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_58_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\58.gif",
		"assets/layui/images/face/58.gif",
	)
}

// assets_layui_images_face_59_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_59_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\59.gif",
		"assets/layui/images/face/59.gif",
	)
}

// assets_layui_images_face_6_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_6_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\6.gif",
		"assets/layui/images/face/6.gif",
	)
}

// assets_layui_images_face_60_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_60_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\60.gif",
		"assets/layui/images/face/60.gif",
	)
}

// assets_layui_images_face_61_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_61_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\61.gif",
		"assets/layui/images/face/61.gif",
	)
}

// assets_layui_images_face_62_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_62_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\62.gif",
		"assets/layui/images/face/62.gif",
	)
}

// assets_layui_images_face_63_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_63_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\63.gif",
		"assets/layui/images/face/63.gif",
	)
}

// assets_layui_images_face_64_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_64_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\64.gif",
		"assets/layui/images/face/64.gif",
	)
}

// assets_layui_images_face_65_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_65_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\65.gif",
		"assets/layui/images/face/65.gif",
	)
}

// assets_layui_images_face_66_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_66_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\66.gif",
		"assets/layui/images/face/66.gif",
	)
}

// assets_layui_images_face_67_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_67_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\67.gif",
		"assets/layui/images/face/67.gif",
	)
}

// assets_layui_images_face_68_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_68_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\68.gif",
		"assets/layui/images/face/68.gif",
	)
}

// assets_layui_images_face_69_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_69_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\69.gif",
		"assets/layui/images/face/69.gif",
	)
}

// assets_layui_images_face_7_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_7_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\7.gif",
		"assets/layui/images/face/7.gif",
	)
}

// assets_layui_images_face_70_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_70_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\70.gif",
		"assets/layui/images/face/70.gif",
	)
}

// assets_layui_images_face_71_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_71_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\71.gif",
		"assets/layui/images/face/71.gif",
	)
}

// assets_layui_images_face_8_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_8_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\8.gif",
		"assets/layui/images/face/8.gif",
	)
}

// assets_layui_images_face_9_gif reads file data from disk. It returns an error on failure.
func assets_layui_images_face_9_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\images\\face\\9.gif",
		"assets/layui/images/face/9.gif",
	)
}

// assets_layui_lay_modules_carousel_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_carousel_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\carousel.js",
		"assets/layui/lay/modules/carousel.js",
	)
}

// assets_layui_lay_modules_code_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_code_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\code.js",
		"assets/layui/lay/modules/code.js",
	)
}

// assets_layui_lay_modules_colorpicker_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_colorpicker_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\colorpicker.js",
		"assets/layui/lay/modules/colorpicker.js",
	)
}

// assets_layui_lay_modules_element_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_element_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\element.js",
		"assets/layui/lay/modules/element.js",
	)
}

// assets_layui_lay_modules_flow_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_flow_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\flow.js",
		"assets/layui/lay/modules/flow.js",
	)
}

// assets_layui_lay_modules_form_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_form_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\form.js",
		"assets/layui/lay/modules/form.js",
	)
}

// assets_layui_lay_modules_jquery_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_jquery_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\jquery.js",
		"assets/layui/lay/modules/jquery.js",
	)
}

// assets_layui_lay_modules_laydate_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_laydate_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\laydate.js",
		"assets/layui/lay/modules/laydate.js",
	)
}

// assets_layui_lay_modules_layedit_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_layedit_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\layedit.js",
		"assets/layui/lay/modules/layedit.js",
	)
}

// assets_layui_lay_modules_layer_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_layer_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\layer.js",
		"assets/layui/lay/modules/layer.js",
	)
}

// assets_layui_lay_modules_laypage_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_laypage_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\laypage.js",
		"assets/layui/lay/modules/laypage.js",
	)
}

// assets_layui_lay_modules_laytpl_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_laytpl_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\laytpl.js",
		"assets/layui/lay/modules/laytpl.js",
	)
}

// assets_layui_lay_modules_mobile_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_mobile_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\mobile.js",
		"assets/layui/lay/modules/mobile.js",
	)
}

// assets_layui_lay_modules_rate_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_rate_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\rate.js",
		"assets/layui/lay/modules/rate.js",
	)
}

// assets_layui_lay_modules_slider_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_slider_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\slider.js",
		"assets/layui/lay/modules/slider.js",
	)
}

// assets_layui_lay_modules_table_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_table_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\table.js",
		"assets/layui/lay/modules/table.js",
	)
}

// assets_layui_lay_modules_transfer_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_transfer_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\transfer.js",
		"assets/layui/lay/modules/transfer.js",
	)
}

// assets_layui_lay_modules_tree_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_tree_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\tree.js",
		"assets/layui/lay/modules/tree.js",
	)
}

// assets_layui_lay_modules_upload_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_upload_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\upload.js",
		"assets/layui/lay/modules/upload.js",
	)
}

// assets_layui_lay_modules_util_js reads file data from disk. It returns an error on failure.
func assets_layui_lay_modules_util_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\lay\\modules\\util.js",
		"assets/layui/lay/modules/util.js",
	)
}

// assets_layui_layui_all_js reads file data from disk. It returns an error on failure.
func assets_layui_layui_all_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\layui.all.js",
		"assets/layui/layui.all.js",
	)
}

// assets_layui_layui_js reads file data from disk. It returns an error on failure.
func assets_layui_layui_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\layui\\layui.js",
		"assets/layui/layui.js",
	)
}

// assets_login_html reads file data from disk. It returns an error on failure.
func assets_login_html() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\login.html",
		"assets/login.html",
	)
}

// assets_static_bootstrap_css_bootstrap_min_css reads file data from disk. It returns an error on failure.
func assets_static_bootstrap_css_bootstrap_min_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\bootstrap\\css\\bootstrap.min.css",
		"assets/static/bootstrap/css/bootstrap.min.css",
	)
}

// assets_static_bootstrap_css_bootstrap_min_css_map reads file data from disk. It returns an error on failure.
func assets_static_bootstrap_css_bootstrap_min_css_map() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\bootstrap\\css\\bootstrap.min.css.map",
		"assets/static/bootstrap/css/bootstrap.min.css.map",
	)
}

// assets_static_bootstrap_fonts_glyphicons_halflings_regular_eot reads file data from disk. It returns an error on failure.
func assets_static_bootstrap_fonts_glyphicons_halflings_regular_eot() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\bootstrap\\fonts\\glyphicons-halflings-regular.eot",
		"assets/static/bootstrap/fonts/glyphicons-halflings-regular.eot",
	)
}

// assets_static_bootstrap_fonts_glyphicons_halflings_regular_svg reads file data from disk. It returns an error on failure.
func assets_static_bootstrap_fonts_glyphicons_halflings_regular_svg() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\bootstrap\\fonts\\glyphicons-halflings-regular.svg",
		"assets/static/bootstrap/fonts/glyphicons-halflings-regular.svg",
	)
}

// assets_static_bootstrap_fonts_glyphicons_halflings_regular_ttf reads file data from disk. It returns an error on failure.
func assets_static_bootstrap_fonts_glyphicons_halflings_regular_ttf() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\bootstrap\\fonts\\glyphicons-halflings-regular.ttf",
		"assets/static/bootstrap/fonts/glyphicons-halflings-regular.ttf",
	)
}

// assets_static_bootstrap_fonts_glyphicons_halflings_regular_woff reads file data from disk. It returns an error on failure.
func assets_static_bootstrap_fonts_glyphicons_halflings_regular_woff() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\bootstrap\\fonts\\glyphicons-halflings-regular.woff",
		"assets/static/bootstrap/fonts/glyphicons-halflings-regular.woff",
	)
}

// assets_static_bootstrap_fonts_glyphicons_halflings_regular_woff2 reads file data from disk. It returns an error on failure.
func assets_static_bootstrap_fonts_glyphicons_halflings_regular_woff2() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\bootstrap\\fonts\\glyphicons-halflings-regular.woff2",
		"assets/static/bootstrap/fonts/glyphicons-halflings-regular.woff2",
	)
}

// assets_static_bootstrap_js_bootstrap_min_js reads file data from disk. It returns an error on failure.
func assets_static_bootstrap_js_bootstrap_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\bootstrap\\js\\bootstrap.min.js",
		"assets/static/bootstrap/js/bootstrap.min.js",
	)
}

// assets_static_css_admin_login_css reads file data from disk. It returns an error on failure.
func assets_static_css_admin_login_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\css\\admin-login.css",
		"assets/static/css/admin-login.css",
	)
}

// assets_static_css_admin_css reads file data from disk. It returns an error on failure.
func assets_static_css_admin_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\css\\admin.css",
		"assets/static/css/admin.css",
	)
}

// assets_static_css_animate_css reads file data from disk. It returns an error on failure.
func assets_static_css_animate_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\css\\animate.css",
		"assets/static/css/animate.css",
	)
}

// assets_static_font_awesome_4_7_0_help_us_out_txt reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_help_us_out_txt() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\HELP-US-OUT.txt",
		"assets/static/font-awesome-4.7.0/HELP-US-OUT.txt",
	)
}

// assets_static_font_awesome_4_7_0_css_font_awesome_css reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_css_font_awesome_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\css\\font-awesome.css",
		"assets/static/font-awesome-4.7.0/css/font-awesome.css",
	)
}

// assets_static_font_awesome_4_7_0_css_font_awesome_min_css reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_css_font_awesome_min_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\css\\font-awesome.min.css",
		"assets/static/font-awesome-4.7.0/css/font-awesome.min.css",
	)
}

// assets_static_font_awesome_4_7_0_fonts_fontawesome_otf reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_fonts_fontawesome_otf() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\fonts\\FontAwesome.otf",
		"assets/static/font-awesome-4.7.0/fonts/FontAwesome.otf",
	)
}

// assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_eot reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_eot() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\fonts\\fontawesome-webfont.eot",
		"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.eot",
	)
}

// assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_svg reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_svg() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\fonts\\fontawesome-webfont.svg",
		"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.svg",
	)
}

// assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_ttf reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_ttf() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\fonts\\fontawesome-webfont.ttf",
		"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.ttf",
	)
}

// assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_woff reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_woff() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\fonts\\fontawesome-webfont.woff",
		"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.woff",
	)
}

// assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_woff2 reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_woff2() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\fonts\\fontawesome-webfont.woff2",
		"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.woff2",
	)
}

// assets_static_font_awesome_4_7_0_less_animated_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_animated_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\animated.less",
		"assets/static/font-awesome-4.7.0/less/animated.less",
	)
}

// assets_static_font_awesome_4_7_0_less_bordered_pulled_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_bordered_pulled_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\bordered-pulled.less",
		"assets/static/font-awesome-4.7.0/less/bordered-pulled.less",
	)
}

// assets_static_font_awesome_4_7_0_less_core_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_core_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\core.less",
		"assets/static/font-awesome-4.7.0/less/core.less",
	)
}

// assets_static_font_awesome_4_7_0_less_fixed_width_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_fixed_width_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\fixed-width.less",
		"assets/static/font-awesome-4.7.0/less/fixed-width.less",
	)
}

// assets_static_font_awesome_4_7_0_less_font_awesome_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_font_awesome_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\font-awesome.less",
		"assets/static/font-awesome-4.7.0/less/font-awesome.less",
	)
}

// assets_static_font_awesome_4_7_0_less_icons_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_icons_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\icons.less",
		"assets/static/font-awesome-4.7.0/less/icons.less",
	)
}

// assets_static_font_awesome_4_7_0_less_larger_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_larger_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\larger.less",
		"assets/static/font-awesome-4.7.0/less/larger.less",
	)
}

// assets_static_font_awesome_4_7_0_less_list_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_list_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\list.less",
		"assets/static/font-awesome-4.7.0/less/list.less",
	)
}

// assets_static_font_awesome_4_7_0_less_mixins_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_mixins_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\mixins.less",
		"assets/static/font-awesome-4.7.0/less/mixins.less",
	)
}

// assets_static_font_awesome_4_7_0_less_path_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_path_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\path.less",
		"assets/static/font-awesome-4.7.0/less/path.less",
	)
}

// assets_static_font_awesome_4_7_0_less_rotated_flipped_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_rotated_flipped_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\rotated-flipped.less",
		"assets/static/font-awesome-4.7.0/less/rotated-flipped.less",
	)
}

// assets_static_font_awesome_4_7_0_less_screen_reader_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_screen_reader_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\screen-reader.less",
		"assets/static/font-awesome-4.7.0/less/screen-reader.less",
	)
}

// assets_static_font_awesome_4_7_0_less_stacked_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_stacked_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\stacked.less",
		"assets/static/font-awesome-4.7.0/less/stacked.less",
	)
}

// assets_static_font_awesome_4_7_0_less_variables_less reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_less_variables_less() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\less\\variables.less",
		"assets/static/font-awesome-4.7.0/less/variables.less",
	)
}

// assets_static_font_awesome_4_7_0_scss_animated_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_animated_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\_animated.scss",
		"assets/static/font-awesome-4.7.0/scss/_animated.scss",
	)
}

// assets_static_font_awesome_4_7_0_scss_bordered_pulled_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_bordered_pulled_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\_bordered-pulled.scss",
		"assets/static/font-awesome-4.7.0/scss/_bordered-pulled.scss",
	)
}

// assets_static_font_awesome_4_7_0_scss_core_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_core_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\_core.scss",
		"assets/static/font-awesome-4.7.0/scss/_core.scss",
	)
}

// assets_static_font_awesome_4_7_0_scss_fixed_width_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_fixed_width_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\_fixed-width.scss",
		"assets/static/font-awesome-4.7.0/scss/_fixed-width.scss",
	)
}

// assets_static_font_awesome_4_7_0_scss_icons_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_icons_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\_icons.scss",
		"assets/static/font-awesome-4.7.0/scss/_icons.scss",
	)
}

// assets_static_font_awesome_4_7_0_scss_larger_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_larger_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\_larger.scss",
		"assets/static/font-awesome-4.7.0/scss/_larger.scss",
	)
}

// assets_static_font_awesome_4_7_0_scss_list_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_list_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\_list.scss",
		"assets/static/font-awesome-4.7.0/scss/_list.scss",
	)
}

// assets_static_font_awesome_4_7_0_scss_mixins_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_mixins_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\_mixins.scss",
		"assets/static/font-awesome-4.7.0/scss/_mixins.scss",
	)
}

// assets_static_font_awesome_4_7_0_scss_path_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_path_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\_path.scss",
		"assets/static/font-awesome-4.7.0/scss/_path.scss",
	)
}

// assets_static_font_awesome_4_7_0_scss_rotated_flipped_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_rotated_flipped_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\_rotated-flipped.scss",
		"assets/static/font-awesome-4.7.0/scss/_rotated-flipped.scss",
	)
}

// assets_static_font_awesome_4_7_0_scss_screen_reader_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_screen_reader_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\_screen-reader.scss",
		"assets/static/font-awesome-4.7.0/scss/_screen-reader.scss",
	)
}

// assets_static_font_awesome_4_7_0_scss_stacked_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_stacked_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\_stacked.scss",
		"assets/static/font-awesome-4.7.0/scss/_stacked.scss",
	)
}

// assets_static_font_awesome_4_7_0_scss_variables_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_variables_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\_variables.scss",
		"assets/static/font-awesome-4.7.0/scss/_variables.scss",
	)
}

// assets_static_font_awesome_4_7_0_scss_font_awesome_scss reads file data from disk. It returns an error on failure.
func assets_static_font_awesome_4_7_0_scss_font_awesome_scss() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\font-awesome-4.7.0\\scss\\font-awesome.scss",
		"assets/static/font-awesome-4.7.0/scss/font-awesome.scss",
	)
}

// assets_static_images_bg_2_jpg reads file data from disk. It returns an error on failure.
func assets_static_images_bg_2_jpg() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\images\\bg_2.jpg",
		"assets/static/images/bg_2.jpg",
	)
}

// assets_static_images_chahao_png reads file data from disk. It returns an error on failure.
func assets_static_images_chahao_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\images\\chahao.png",
		"assets/static/images/chahao.png",
	)
}

// assets_static_images_duihao_png reads file data from disk. It returns an error on failure.
func assets_static_images_duihao_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\images\\duihao.png",
		"assets/static/images/duihao.png",
	)
}

// assets_static_images_geometry2_png reads file data from disk. It returns an error on failure.
func assets_static_images_geometry2_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\images\\geometry2.png",
		"assets/static/images/geometry2.png",
	)
}

// assets_static_images_hevr_png reads file data from disk. It returns an error on failure.
func assets_static_images_hevr_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\images\\hevr.png",
		"assets/static/images/hevr.png",
	)
}

// assets_static_images_logo_png reads file data from disk. It returns an error on failure.
func assets_static_images_logo_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\images\\logo.png",
		"assets/static/images/logo.png",
	)
}

// assets_static_images_qrcode_jpg reads file data from disk. It returns an error on failure.
func assets_static_images_qrcode_jpg() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\images\\qrcode.jpg",
		"assets/static/images/qrcode.jpg",
	)
}

// assets_static_images_success_png reads file data from disk. It returns an error on failure.
func assets_static_images_success_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\images\\success.png",
		"assets/static/images/success.png",
	)
}

// assets_static_images_upload_jpg reads file data from disk. It returns an error on failure.
func assets_static_images_upload_jpg() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\images\\upload.jpg",
		"assets/static/images/upload.jpg",
	)
}

// assets_static_js_admin_login_js reads file data from disk. It returns an error on failure.
func assets_static_js_admin_login_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\js\\admin-login.js",
		"assets/static/js/admin-login.js",
	)
}

// assets_static_js_admin_js reads file data from disk. It returns an error on failure.
func assets_static_js_admin_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\js\\admin.js",
		"assets/static/js/admin.js",
	)
}

// assets_static_js_common_js reads file data from disk. It returns an error on failure.
func assets_static_js_common_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\js\\common.js",
		"assets/static/js/common.js",
	)
}

// assets_static_js_echarts_min_js reads file data from disk. It returns an error on failure.
func assets_static_js_echarts_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\js\\echarts.min.js",
		"assets/static/js/echarts.min.js",
	)
}

// assets_static_js_jquery_min_js reads file data from disk. It returns an error on failure.
func assets_static_js_jquery_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\js\\jquery.min.js",
		"assets/static/js/jquery.min.js",
	)
}

// assets_static_js_jquery_placeholder_min_js reads file data from disk. It returns an error on failure.
func assets_static_js_jquery_placeholder_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\js\\jquery.placeholder.min.js",
		"assets/static/js/jquery.placeholder.min.js",
	)
}

// assets_static_js_jquery_waypoints_min_js reads file data from disk. It returns an error on failure.
func assets_static_js_jquery_waypoints_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\js\\jquery.waypoints.min.js",
		"assets/static/js/jquery.waypoints.min.js",
	)
}

// assets_static_js_modernizr_2_6_2_min_js reads file data from disk. It returns an error on failure.
func assets_static_js_modernizr_2_6_2_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\js\\modernizr-2.6.2.min.js",
		"assets/static/js/modernizr-2.6.2.min.js",
	)
}

// assets_static_js_modernizr_custom_v2_7_1_min_js reads file data from disk. It returns an error on failure.
func assets_static_js_modernizr_custom_v2_7_1_min_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\js\\modernizr-custom-v2.7.1.min.js",
		"assets/static/js/modernizr-custom-v2.7.1.min.js",
	)
}

// assets_static_layui_css_formselects_v4_css reads file data from disk. It returns an error on failure.
func assets_static_layui_css_formselects_v4_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\css\\formSelects-v4.css",
		"assets/static/layui/css/formSelects-v4.css",
	)
}

// assets_static_layui_css_layui_css reads file data from disk. It returns an error on failure.
func assets_static_layui_css_layui_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\css\\layui.css",
		"assets/static/layui/css/layui.css",
	)
}

// assets_static_layui_css_layui_mobile_css reads file data from disk. It returns an error on failure.
func assets_static_layui_css_layui_mobile_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\css\\layui.mobile.css",
		"assets/static/layui/css/layui.mobile.css",
	)
}

// assets_static_layui_css_modules_code_css reads file data from disk. It returns an error on failure.
func assets_static_layui_css_modules_code_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\css\\modules\\code.css",
		"assets/static/layui/css/modules/code.css",
	)
}

// assets_static_layui_css_modules_laydate_default_laydate_css reads file data from disk. It returns an error on failure.
func assets_static_layui_css_modules_laydate_default_laydate_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\css\\modules\\laydate\\default\\laydate.css",
		"assets/static/layui/css/modules/laydate/default/laydate.css",
	)
}

// assets_static_layui_css_modules_layer_default_icon_ext_png reads file data from disk. It returns an error on failure.
func assets_static_layui_css_modules_layer_default_icon_ext_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\css\\modules\\layer\\default\\icon-ext.png",
		"assets/static/layui/css/modules/layer/default/icon-ext.png",
	)
}

// assets_static_layui_css_modules_layer_default_icon_png reads file data from disk. It returns an error on failure.
func assets_static_layui_css_modules_layer_default_icon_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\css\\modules\\layer\\default\\icon.png",
		"assets/static/layui/css/modules/layer/default/icon.png",
	)
}

// assets_static_layui_css_modules_layer_default_layer_css reads file data from disk. It returns an error on failure.
func assets_static_layui_css_modules_layer_default_layer_css() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\css\\modules\\layer\\default\\layer.css",
		"assets/static/layui/css/modules/layer/default/layer.css",
	)
}

// assets_static_layui_css_modules_layer_default_loading_0_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_css_modules_layer_default_loading_0_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\css\\modules\\layer\\default\\loading-0.gif",
		"assets/static/layui/css/modules/layer/default/loading-0.gif",
	)
}

// assets_static_layui_css_modules_layer_default_loading_1_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_css_modules_layer_default_loading_1_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\css\\modules\\layer\\default\\loading-1.gif",
		"assets/static/layui/css/modules/layer/default/loading-1.gif",
	)
}

// assets_static_layui_css_modules_layer_default_loading_2_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_css_modules_layer_default_loading_2_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\css\\modules\\layer\\default\\loading-2.gif",
		"assets/static/layui/css/modules/layer/default/loading-2.gif",
	)
}

// assets_static_layui_font_iconfont_eot reads file data from disk. It returns an error on failure.
func assets_static_layui_font_iconfont_eot() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\font\\iconfont.eot",
		"assets/static/layui/font/iconfont.eot",
	)
}

// assets_static_layui_font_iconfont_svg reads file data from disk. It returns an error on failure.
func assets_static_layui_font_iconfont_svg() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\font\\iconfont.svg",
		"assets/static/layui/font/iconfont.svg",
	)
}

// assets_static_layui_font_iconfont_ttf reads file data from disk. It returns an error on failure.
func assets_static_layui_font_iconfont_ttf() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\font\\iconfont.ttf",
		"assets/static/layui/font/iconfont.ttf",
	)
}

// assets_static_layui_font_iconfont_woff reads file data from disk. It returns an error on failure.
func assets_static_layui_font_iconfont_woff() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\font\\iconfont.woff",
		"assets/static/layui/font/iconfont.woff",
	)
}

// assets_static_layui_font_iconfont_woff2 reads file data from disk. It returns an error on failure.
func assets_static_layui_font_iconfont_woff2() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\font\\iconfont.woff2",
		"assets/static/layui/font/iconfont.woff2",
	)
}

// assets_static_layui_formselects_v4_js reads file data from disk. It returns an error on failure.
func assets_static_layui_formselects_v4_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\formSelects-v4.js",
		"assets/static/layui/formSelects-v4.js",
	)
}

// assets_static_layui_images_face_0_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_0_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\0.gif",
		"assets/static/layui/images/face/0.gif",
	)
}

// assets_static_layui_images_face_1_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_1_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\1.gif",
		"assets/static/layui/images/face/1.gif",
	)
}

// assets_static_layui_images_face_10_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_10_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\10.gif",
		"assets/static/layui/images/face/10.gif",
	)
}

// assets_static_layui_images_face_11_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_11_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\11.gif",
		"assets/static/layui/images/face/11.gif",
	)
}

// assets_static_layui_images_face_12_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_12_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\12.gif",
		"assets/static/layui/images/face/12.gif",
	)
}

// assets_static_layui_images_face_13_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_13_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\13.gif",
		"assets/static/layui/images/face/13.gif",
	)
}

// assets_static_layui_images_face_14_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_14_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\14.gif",
		"assets/static/layui/images/face/14.gif",
	)
}

// assets_static_layui_images_face_15_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_15_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\15.gif",
		"assets/static/layui/images/face/15.gif",
	)
}

// assets_static_layui_images_face_16_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_16_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\16.gif",
		"assets/static/layui/images/face/16.gif",
	)
}

// assets_static_layui_images_face_17_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_17_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\17.gif",
		"assets/static/layui/images/face/17.gif",
	)
}

// assets_static_layui_images_face_18_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_18_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\18.gif",
		"assets/static/layui/images/face/18.gif",
	)
}

// assets_static_layui_images_face_19_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_19_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\19.gif",
		"assets/static/layui/images/face/19.gif",
	)
}

// assets_static_layui_images_face_2_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_2_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\2.gif",
		"assets/static/layui/images/face/2.gif",
	)
}

// assets_static_layui_images_face_20_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_20_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\20.gif",
		"assets/static/layui/images/face/20.gif",
	)
}

// assets_static_layui_images_face_21_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_21_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\21.gif",
		"assets/static/layui/images/face/21.gif",
	)
}

// assets_static_layui_images_face_22_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_22_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\22.gif",
		"assets/static/layui/images/face/22.gif",
	)
}

// assets_static_layui_images_face_23_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_23_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\23.gif",
		"assets/static/layui/images/face/23.gif",
	)
}

// assets_static_layui_images_face_24_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_24_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\24.gif",
		"assets/static/layui/images/face/24.gif",
	)
}

// assets_static_layui_images_face_25_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_25_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\25.gif",
		"assets/static/layui/images/face/25.gif",
	)
}

// assets_static_layui_images_face_26_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_26_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\26.gif",
		"assets/static/layui/images/face/26.gif",
	)
}

// assets_static_layui_images_face_27_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_27_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\27.gif",
		"assets/static/layui/images/face/27.gif",
	)
}

// assets_static_layui_images_face_28_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_28_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\28.gif",
		"assets/static/layui/images/face/28.gif",
	)
}

// assets_static_layui_images_face_29_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_29_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\29.gif",
		"assets/static/layui/images/face/29.gif",
	)
}

// assets_static_layui_images_face_3_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_3_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\3.gif",
		"assets/static/layui/images/face/3.gif",
	)
}

// assets_static_layui_images_face_30_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_30_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\30.gif",
		"assets/static/layui/images/face/30.gif",
	)
}

// assets_static_layui_images_face_31_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_31_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\31.gif",
		"assets/static/layui/images/face/31.gif",
	)
}

// assets_static_layui_images_face_32_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_32_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\32.gif",
		"assets/static/layui/images/face/32.gif",
	)
}

// assets_static_layui_images_face_33_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_33_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\33.gif",
		"assets/static/layui/images/face/33.gif",
	)
}

// assets_static_layui_images_face_34_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_34_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\34.gif",
		"assets/static/layui/images/face/34.gif",
	)
}

// assets_static_layui_images_face_35_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_35_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\35.gif",
		"assets/static/layui/images/face/35.gif",
	)
}

// assets_static_layui_images_face_36_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_36_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\36.gif",
		"assets/static/layui/images/face/36.gif",
	)
}

// assets_static_layui_images_face_37_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_37_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\37.gif",
		"assets/static/layui/images/face/37.gif",
	)
}

// assets_static_layui_images_face_38_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_38_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\38.gif",
		"assets/static/layui/images/face/38.gif",
	)
}

// assets_static_layui_images_face_39_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_39_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\39.gif",
		"assets/static/layui/images/face/39.gif",
	)
}

// assets_static_layui_images_face_4_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_4_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\4.gif",
		"assets/static/layui/images/face/4.gif",
	)
}

// assets_static_layui_images_face_40_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_40_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\40.gif",
		"assets/static/layui/images/face/40.gif",
	)
}

// assets_static_layui_images_face_41_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_41_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\41.gif",
		"assets/static/layui/images/face/41.gif",
	)
}

// assets_static_layui_images_face_42_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_42_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\42.gif",
		"assets/static/layui/images/face/42.gif",
	)
}

// assets_static_layui_images_face_43_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_43_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\43.gif",
		"assets/static/layui/images/face/43.gif",
	)
}

// assets_static_layui_images_face_44_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_44_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\44.gif",
		"assets/static/layui/images/face/44.gif",
	)
}

// assets_static_layui_images_face_45_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_45_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\45.gif",
		"assets/static/layui/images/face/45.gif",
	)
}

// assets_static_layui_images_face_46_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_46_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\46.gif",
		"assets/static/layui/images/face/46.gif",
	)
}

// assets_static_layui_images_face_47_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_47_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\47.gif",
		"assets/static/layui/images/face/47.gif",
	)
}

// assets_static_layui_images_face_48_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_48_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\48.gif",
		"assets/static/layui/images/face/48.gif",
	)
}

// assets_static_layui_images_face_49_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_49_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\49.gif",
		"assets/static/layui/images/face/49.gif",
	)
}

// assets_static_layui_images_face_5_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_5_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\5.gif",
		"assets/static/layui/images/face/5.gif",
	)
}

// assets_static_layui_images_face_50_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_50_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\50.gif",
		"assets/static/layui/images/face/50.gif",
	)
}

// assets_static_layui_images_face_51_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_51_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\51.gif",
		"assets/static/layui/images/face/51.gif",
	)
}

// assets_static_layui_images_face_52_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_52_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\52.gif",
		"assets/static/layui/images/face/52.gif",
	)
}

// assets_static_layui_images_face_53_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_53_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\53.gif",
		"assets/static/layui/images/face/53.gif",
	)
}

// assets_static_layui_images_face_54_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_54_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\54.gif",
		"assets/static/layui/images/face/54.gif",
	)
}

// assets_static_layui_images_face_55_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_55_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\55.gif",
		"assets/static/layui/images/face/55.gif",
	)
}

// assets_static_layui_images_face_56_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_56_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\56.gif",
		"assets/static/layui/images/face/56.gif",
	)
}

// assets_static_layui_images_face_57_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_57_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\57.gif",
		"assets/static/layui/images/face/57.gif",
	)
}

// assets_static_layui_images_face_58_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_58_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\58.gif",
		"assets/static/layui/images/face/58.gif",
	)
}

// assets_static_layui_images_face_59_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_59_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\59.gif",
		"assets/static/layui/images/face/59.gif",
	)
}

// assets_static_layui_images_face_6_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_6_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\6.gif",
		"assets/static/layui/images/face/6.gif",
	)
}

// assets_static_layui_images_face_60_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_60_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\60.gif",
		"assets/static/layui/images/face/60.gif",
	)
}

// assets_static_layui_images_face_61_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_61_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\61.gif",
		"assets/static/layui/images/face/61.gif",
	)
}

// assets_static_layui_images_face_62_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_62_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\62.gif",
		"assets/static/layui/images/face/62.gif",
	)
}

// assets_static_layui_images_face_63_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_63_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\63.gif",
		"assets/static/layui/images/face/63.gif",
	)
}

// assets_static_layui_images_face_64_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_64_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\64.gif",
		"assets/static/layui/images/face/64.gif",
	)
}

// assets_static_layui_images_face_65_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_65_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\65.gif",
		"assets/static/layui/images/face/65.gif",
	)
}

// assets_static_layui_images_face_66_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_66_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\66.gif",
		"assets/static/layui/images/face/66.gif",
	)
}

// assets_static_layui_images_face_67_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_67_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\67.gif",
		"assets/static/layui/images/face/67.gif",
	)
}

// assets_static_layui_images_face_68_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_68_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\68.gif",
		"assets/static/layui/images/face/68.gif",
	)
}

// assets_static_layui_images_face_69_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_69_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\69.gif",
		"assets/static/layui/images/face/69.gif",
	)
}

// assets_static_layui_images_face_7_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_7_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\7.gif",
		"assets/static/layui/images/face/7.gif",
	)
}

// assets_static_layui_images_face_70_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_70_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\70.gif",
		"assets/static/layui/images/face/70.gif",
	)
}

// assets_static_layui_images_face_71_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_71_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\71.gif",
		"assets/static/layui/images/face/71.gif",
	)
}

// assets_static_layui_images_face_8_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_8_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\8.gif",
		"assets/static/layui/images/face/8.gif",
	)
}

// assets_static_layui_images_face_9_gif reads file data from disk. It returns an error on failure.
func assets_static_layui_images_face_9_gif() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\images\\face\\9.gif",
		"assets/static/layui/images/face/9.gif",
	)
}

// assets_static_layui_lay_modules_carousel_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_carousel_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\carousel.js",
		"assets/static/layui/lay/modules/carousel.js",
	)
}

// assets_static_layui_lay_modules_code_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_code_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\code.js",
		"assets/static/layui/lay/modules/code.js",
	)
}

// assets_static_layui_lay_modules_colorpicker_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_colorpicker_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\colorpicker.js",
		"assets/static/layui/lay/modules/colorpicker.js",
	)
}

// assets_static_layui_lay_modules_element_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_element_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\element.js",
		"assets/static/layui/lay/modules/element.js",
	)
}

// assets_static_layui_lay_modules_flow_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_flow_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\flow.js",
		"assets/static/layui/lay/modules/flow.js",
	)
}

// assets_static_layui_lay_modules_form_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_form_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\form.js",
		"assets/static/layui/lay/modules/form.js",
	)
}

// assets_static_layui_lay_modules_jquery_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_jquery_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\jquery.js",
		"assets/static/layui/lay/modules/jquery.js",
	)
}

// assets_static_layui_lay_modules_laydate_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_laydate_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\laydate.js",
		"assets/static/layui/lay/modules/laydate.js",
	)
}

// assets_static_layui_lay_modules_layedit_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_layedit_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\layedit.js",
		"assets/static/layui/lay/modules/layedit.js",
	)
}

// assets_static_layui_lay_modules_layer_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_layer_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\layer.js",
		"assets/static/layui/lay/modules/layer.js",
	)
}

// assets_static_layui_lay_modules_laypage_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_laypage_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\laypage.js",
		"assets/static/layui/lay/modules/laypage.js",
	)
}

// assets_static_layui_lay_modules_laytpl_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_laytpl_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\laytpl.js",
		"assets/static/layui/lay/modules/laytpl.js",
	)
}

// assets_static_layui_lay_modules_mobile_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_mobile_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\mobile.js",
		"assets/static/layui/lay/modules/mobile.js",
	)
}

// assets_static_layui_lay_modules_rate_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_rate_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\rate.js",
		"assets/static/layui/lay/modules/rate.js",
	)
}

// assets_static_layui_lay_modules_slider_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_slider_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\slider.js",
		"assets/static/layui/lay/modules/slider.js",
	)
}

// assets_static_layui_lay_modules_table_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_table_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\table.js",
		"assets/static/layui/lay/modules/table.js",
	)
}

// assets_static_layui_lay_modules_transfer_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_transfer_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\transfer.js",
		"assets/static/layui/lay/modules/transfer.js",
	)
}

// assets_static_layui_lay_modules_tree_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_tree_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\tree.js",
		"assets/static/layui/lay/modules/tree.js",
	)
}

// assets_static_layui_lay_modules_upload_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_upload_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\upload.js",
		"assets/static/layui/lay/modules/upload.js",
	)
}

// assets_static_layui_lay_modules_util_js reads file data from disk. It returns an error on failure.
func assets_static_layui_lay_modules_util_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\lay\\modules\\util.js",
		"assets/static/layui/lay/modules/util.js",
	)
}

// assets_static_layui_layui_all_js reads file data from disk. It returns an error on failure.
func assets_static_layui_layui_all_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\layui.all.js",
		"assets/static/layui/layui.all.js",
	)
}

// assets_static_layui_layui_js reads file data from disk. It returns an error on failure.
func assets_static_layui_layui_js() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\layui\\layui.js",
		"assets/static/layui/layui.js",
	)
}

// assets_static_upload_imgs_20190905_1567651858020_647_png reads file data from disk. It returns an error on failure.
func assets_static_upload_imgs_20190905_1567651858020_647_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\static\\upload\\imgs\\20190905\\1567651858020_647.png",
		"assets/static/upload/imgs/20190905/1567651858020_647.png",
	)
}

// assets_upload_imgs_20190905_1567651858020_647_png reads file data from disk. It returns an error on failure.
func assets_upload_imgs_20190905_1567651858020_647_png() ([]byte, error) {
	return bindata_read(
		"D:\\GoPath\\src\\letinvr.com\\echo-web\\assets\\upload\\imgs\\20190905\\1567651858020_647.png",
		"assets/upload/imgs/20190905/1567651858020_647.png",
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
	"assets/bootstrap/css/bootstrap.min.css": assets_bootstrap_css_bootstrap_min_css,
	"assets/bootstrap/css/bootstrap.min.css.map": assets_bootstrap_css_bootstrap_min_css_map,
	"assets/bootstrap/fonts/glyphicons-halflings-regular.eot": assets_bootstrap_fonts_glyphicons_halflings_regular_eot,
	"assets/bootstrap/fonts/glyphicons-halflings-regular.svg": assets_bootstrap_fonts_glyphicons_halflings_regular_svg,
	"assets/bootstrap/fonts/glyphicons-halflings-regular.ttf": assets_bootstrap_fonts_glyphicons_halflings_regular_ttf,
	"assets/bootstrap/fonts/glyphicons-halflings-regular.woff": assets_bootstrap_fonts_glyphicons_halflings_regular_woff,
	"assets/bootstrap/fonts/glyphicons-halflings-regular.woff2": assets_bootstrap_fonts_glyphicons_halflings_regular_woff2,
	"assets/bootstrap/js/bootstrap.min.js": assets_bootstrap_js_bootstrap_min_js,
	"assets/css/admin-login.css": assets_css_admin_login_css,
	"assets/css/admin.css": assets_css_admin_css,
	"assets/css/animate.css": assets_css_animate_css,
	"assets/css/bootstrap-combined.min.css": assets_css_bootstrap_combined_min_css,
	"assets/css/bootstrap-theme.css": assets_css_bootstrap_theme_css,
	"assets/css/bootstrap-theme.css.map": assets_css_bootstrap_theme_css_map,
	"assets/css/bootstrap-theme.min.css": assets_css_bootstrap_theme_min_css,
	"assets/css/bootstrap.css": assets_css_bootstrap_css,
	"assets/css/bootstrap.css.map": assets_css_bootstrap_css_map,
	"assets/css/bootstrap.min.css": assets_css_bootstrap_min_css,
	"assets/css/jquery.treetable.css": assets_css_jquery_treetable_css,
	"assets/css/jquery.treetable.theme.default.css": assets_css_jquery_treetable_theme_default_css,
	"assets/font-awesome-4.7.0/HELP-US-OUT.txt": assets_font_awesome_4_7_0_help_us_out_txt,
	"assets/font-awesome-4.7.0/css/font-awesome.css": assets_font_awesome_4_7_0_css_font_awesome_css,
	"assets/font-awesome-4.7.0/css/font-awesome.min.css": assets_font_awesome_4_7_0_css_font_awesome_min_css,
	"assets/font-awesome-4.7.0/fonts/FontAwesome.otf": assets_font_awesome_4_7_0_fonts_fontawesome_otf,
	"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.eot": assets_font_awesome_4_7_0_fonts_fontawesome_webfont_eot,
	"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.svg": assets_font_awesome_4_7_0_fonts_fontawesome_webfont_svg,
	"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.ttf": assets_font_awesome_4_7_0_fonts_fontawesome_webfont_ttf,
	"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.woff": assets_font_awesome_4_7_0_fonts_fontawesome_webfont_woff,
	"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.woff2": assets_font_awesome_4_7_0_fonts_fontawesome_webfont_woff2,
	"assets/font-awesome-4.7.0/less/animated.less": assets_font_awesome_4_7_0_less_animated_less,
	"assets/font-awesome-4.7.0/less/bordered-pulled.less": assets_font_awesome_4_7_0_less_bordered_pulled_less,
	"assets/font-awesome-4.7.0/less/core.less": assets_font_awesome_4_7_0_less_core_less,
	"assets/font-awesome-4.7.0/less/fixed-width.less": assets_font_awesome_4_7_0_less_fixed_width_less,
	"assets/font-awesome-4.7.0/less/font-awesome.less": assets_font_awesome_4_7_0_less_font_awesome_less,
	"assets/font-awesome-4.7.0/less/icons.less": assets_font_awesome_4_7_0_less_icons_less,
	"assets/font-awesome-4.7.0/less/larger.less": assets_font_awesome_4_7_0_less_larger_less,
	"assets/font-awesome-4.7.0/less/list.less": assets_font_awesome_4_7_0_less_list_less,
	"assets/font-awesome-4.7.0/less/mixins.less": assets_font_awesome_4_7_0_less_mixins_less,
	"assets/font-awesome-4.7.0/less/path.less": assets_font_awesome_4_7_0_less_path_less,
	"assets/font-awesome-4.7.0/less/rotated-flipped.less": assets_font_awesome_4_7_0_less_rotated_flipped_less,
	"assets/font-awesome-4.7.0/less/screen-reader.less": assets_font_awesome_4_7_0_less_screen_reader_less,
	"assets/font-awesome-4.7.0/less/stacked.less": assets_font_awesome_4_7_0_less_stacked_less,
	"assets/font-awesome-4.7.0/less/variables.less": assets_font_awesome_4_7_0_less_variables_less,
	"assets/font-awesome-4.7.0/scss/_animated.scss": assets_font_awesome_4_7_0_scss_animated_scss,
	"assets/font-awesome-4.7.0/scss/_bordered-pulled.scss": assets_font_awesome_4_7_0_scss_bordered_pulled_scss,
	"assets/font-awesome-4.7.0/scss/_core.scss": assets_font_awesome_4_7_0_scss_core_scss,
	"assets/font-awesome-4.7.0/scss/_fixed-width.scss": assets_font_awesome_4_7_0_scss_fixed_width_scss,
	"assets/font-awesome-4.7.0/scss/_icons.scss": assets_font_awesome_4_7_0_scss_icons_scss,
	"assets/font-awesome-4.7.0/scss/_larger.scss": assets_font_awesome_4_7_0_scss_larger_scss,
	"assets/font-awesome-4.7.0/scss/_list.scss": assets_font_awesome_4_7_0_scss_list_scss,
	"assets/font-awesome-4.7.0/scss/_mixins.scss": assets_font_awesome_4_7_0_scss_mixins_scss,
	"assets/font-awesome-4.7.0/scss/_path.scss": assets_font_awesome_4_7_0_scss_path_scss,
	"assets/font-awesome-4.7.0/scss/_rotated-flipped.scss": assets_font_awesome_4_7_0_scss_rotated_flipped_scss,
	"assets/font-awesome-4.7.0/scss/_screen-reader.scss": assets_font_awesome_4_7_0_scss_screen_reader_scss,
	"assets/font-awesome-4.7.0/scss/_stacked.scss": assets_font_awesome_4_7_0_scss_stacked_scss,
	"assets/font-awesome-4.7.0/scss/_variables.scss": assets_font_awesome_4_7_0_scss_variables_scss,
	"assets/font-awesome-4.7.0/scss/font-awesome.scss": assets_font_awesome_4_7_0_scss_font_awesome_scss,
	"assets/fonts/glyphicons-halflings-regular.eot": assets_fonts_glyphicons_halflings_regular_eot,
	"assets/fonts/glyphicons-halflings-regular.svg": assets_fonts_glyphicons_halflings_regular_svg,
	"assets/fonts/glyphicons-halflings-regular.ttf": assets_fonts_glyphicons_halflings_regular_ttf,
	"assets/fonts/glyphicons-halflings-regular.woff": assets_fonts_glyphicons_halflings_regular_woff,
	"assets/fonts/glyphicons-halflings-regular.woff2": assets_fonts_glyphicons_halflings_regular_woff2,
	"assets/images/bg_2.jpg": assets_images_bg_2_jpg,
	"assets/images/chahao.png": assets_images_chahao_png,
	"assets/images/duihao.png": assets_images_duihao_png,
	"assets/images/geometry2.png": assets_images_geometry2_png,
	"assets/images/hevr.png": assets_images_hevr_png,
	"assets/images/logo.png": assets_images_logo_png,
	"assets/images/qrcode.jpg": assets_images_qrcode_jpg,
	"assets/images/success.png": assets_images_success_png,
	"assets/images/upload.jpg": assets_images_upload_jpg,
	"assets/img/favicon.ico": assets_img_favicon_ico,
	"assets/img/icon_brand.png": assets_img_icon_brand_png,
	"assets/js/admin-login.js": assets_js_admin_login_js,
	"assets/js/admin.js": assets_js_admin_js,
	"assets/js/app/jwt/jwt.js": assets_js_app_jwt_jwt_js,
	"assets/js/app/treetable/treetable-page.js": assets_js_app_treetable_treetable_page_js,
	"assets/js/common.js": assets_js_common_js,
	"assets/js/echarts.min.js": assets_js_echarts_min_js,
	"assets/js/jquery.min.js": assets_js_jquery_min_js,
	"assets/js/jquery.placeholder.min.js": assets_js_jquery_placeholder_min_js,
	"assets/js/jquery.waypoints.min.js": assets_js_jquery_waypoints_min_js,
	"assets/js/lib/bootstrap/bootstrap.js": assets_js_lib_bootstrap_bootstrap_js,
	"assets/js/lib/bootstrap/bootstrap.min.js": assets_js_lib_bootstrap_bootstrap_min_js,
	"assets/js/lib/bootstrap/npm.js": assets_js_lib_bootstrap_npm_js,
	"assets/js/lib/jquery-form/jquery.form.3.51.js": assets_js_lib_jquery_form_jquery_form_3_51_js,
	"assets/js/lib/jquery-treetable/jquery.treetable.js": assets_js_lib_jquery_treetable_jquery_treetable_js,
	"assets/js/lib/jquery/jquery-2.1.3.js": assets_js_lib_jquery_jquery_2_1_3_js,
	"assets/js/lib/jquery/jquery-2.1.3.min.js": assets_js_lib_jquery_jquery_2_1_3_min_js,
	"assets/js/lib/require-setup.js": assets_js_lib_require_setup_js,
	"assets/js/lib/require.js": assets_js_lib_require_js,
	"assets/js/modernizr-2.6.2.min.js": assets_js_modernizr_2_6_2_min_js,
	"assets/js/modernizr-custom-v2.7.1.min.js": assets_js_modernizr_custom_v2_7_1_min_js,
	"assets/js/tools/r.js": assets_js_tools_r_js,
	"assets/layui/css/layui.css": assets_layui_css_layui_css,
	"assets/layui/css/layui.mobile.css": assets_layui_css_layui_mobile_css,
	"assets/layui/css/modules/code.css": assets_layui_css_modules_code_css,
	"assets/layui/css/modules/laydate/default/laydate.css": assets_layui_css_modules_laydate_default_laydate_css,
	"assets/layui/css/modules/layer/default/icon-ext.png": assets_layui_css_modules_layer_default_icon_ext_png,
	"assets/layui/css/modules/layer/default/icon.png": assets_layui_css_modules_layer_default_icon_png,
	"assets/layui/css/modules/layer/default/layer.css": assets_layui_css_modules_layer_default_layer_css,
	"assets/layui/css/modules/layer/default/loading-0.gif": assets_layui_css_modules_layer_default_loading_0_gif,
	"assets/layui/css/modules/layer/default/loading-1.gif": assets_layui_css_modules_layer_default_loading_1_gif,
	"assets/layui/css/modules/layer/default/loading-2.gif": assets_layui_css_modules_layer_default_loading_2_gif,
	"assets/layui/font/iconfont.eot": assets_layui_font_iconfont_eot,
	"assets/layui/font/iconfont.svg": assets_layui_font_iconfont_svg,
	"assets/layui/font/iconfont.ttf": assets_layui_font_iconfont_ttf,
	"assets/layui/font/iconfont.woff": assets_layui_font_iconfont_woff,
	"assets/layui/font/iconfont.woff2": assets_layui_font_iconfont_woff2,
	"assets/layui/images/face/0.gif": assets_layui_images_face_0_gif,
	"assets/layui/images/face/1.gif": assets_layui_images_face_1_gif,
	"assets/layui/images/face/10.gif": assets_layui_images_face_10_gif,
	"assets/layui/images/face/11.gif": assets_layui_images_face_11_gif,
	"assets/layui/images/face/12.gif": assets_layui_images_face_12_gif,
	"assets/layui/images/face/13.gif": assets_layui_images_face_13_gif,
	"assets/layui/images/face/14.gif": assets_layui_images_face_14_gif,
	"assets/layui/images/face/15.gif": assets_layui_images_face_15_gif,
	"assets/layui/images/face/16.gif": assets_layui_images_face_16_gif,
	"assets/layui/images/face/17.gif": assets_layui_images_face_17_gif,
	"assets/layui/images/face/18.gif": assets_layui_images_face_18_gif,
	"assets/layui/images/face/19.gif": assets_layui_images_face_19_gif,
	"assets/layui/images/face/2.gif": assets_layui_images_face_2_gif,
	"assets/layui/images/face/20.gif": assets_layui_images_face_20_gif,
	"assets/layui/images/face/21.gif": assets_layui_images_face_21_gif,
	"assets/layui/images/face/22.gif": assets_layui_images_face_22_gif,
	"assets/layui/images/face/23.gif": assets_layui_images_face_23_gif,
	"assets/layui/images/face/24.gif": assets_layui_images_face_24_gif,
	"assets/layui/images/face/25.gif": assets_layui_images_face_25_gif,
	"assets/layui/images/face/26.gif": assets_layui_images_face_26_gif,
	"assets/layui/images/face/27.gif": assets_layui_images_face_27_gif,
	"assets/layui/images/face/28.gif": assets_layui_images_face_28_gif,
	"assets/layui/images/face/29.gif": assets_layui_images_face_29_gif,
	"assets/layui/images/face/3.gif": assets_layui_images_face_3_gif,
	"assets/layui/images/face/30.gif": assets_layui_images_face_30_gif,
	"assets/layui/images/face/31.gif": assets_layui_images_face_31_gif,
	"assets/layui/images/face/32.gif": assets_layui_images_face_32_gif,
	"assets/layui/images/face/33.gif": assets_layui_images_face_33_gif,
	"assets/layui/images/face/34.gif": assets_layui_images_face_34_gif,
	"assets/layui/images/face/35.gif": assets_layui_images_face_35_gif,
	"assets/layui/images/face/36.gif": assets_layui_images_face_36_gif,
	"assets/layui/images/face/37.gif": assets_layui_images_face_37_gif,
	"assets/layui/images/face/38.gif": assets_layui_images_face_38_gif,
	"assets/layui/images/face/39.gif": assets_layui_images_face_39_gif,
	"assets/layui/images/face/4.gif": assets_layui_images_face_4_gif,
	"assets/layui/images/face/40.gif": assets_layui_images_face_40_gif,
	"assets/layui/images/face/41.gif": assets_layui_images_face_41_gif,
	"assets/layui/images/face/42.gif": assets_layui_images_face_42_gif,
	"assets/layui/images/face/43.gif": assets_layui_images_face_43_gif,
	"assets/layui/images/face/44.gif": assets_layui_images_face_44_gif,
	"assets/layui/images/face/45.gif": assets_layui_images_face_45_gif,
	"assets/layui/images/face/46.gif": assets_layui_images_face_46_gif,
	"assets/layui/images/face/47.gif": assets_layui_images_face_47_gif,
	"assets/layui/images/face/48.gif": assets_layui_images_face_48_gif,
	"assets/layui/images/face/49.gif": assets_layui_images_face_49_gif,
	"assets/layui/images/face/5.gif": assets_layui_images_face_5_gif,
	"assets/layui/images/face/50.gif": assets_layui_images_face_50_gif,
	"assets/layui/images/face/51.gif": assets_layui_images_face_51_gif,
	"assets/layui/images/face/52.gif": assets_layui_images_face_52_gif,
	"assets/layui/images/face/53.gif": assets_layui_images_face_53_gif,
	"assets/layui/images/face/54.gif": assets_layui_images_face_54_gif,
	"assets/layui/images/face/55.gif": assets_layui_images_face_55_gif,
	"assets/layui/images/face/56.gif": assets_layui_images_face_56_gif,
	"assets/layui/images/face/57.gif": assets_layui_images_face_57_gif,
	"assets/layui/images/face/58.gif": assets_layui_images_face_58_gif,
	"assets/layui/images/face/59.gif": assets_layui_images_face_59_gif,
	"assets/layui/images/face/6.gif": assets_layui_images_face_6_gif,
	"assets/layui/images/face/60.gif": assets_layui_images_face_60_gif,
	"assets/layui/images/face/61.gif": assets_layui_images_face_61_gif,
	"assets/layui/images/face/62.gif": assets_layui_images_face_62_gif,
	"assets/layui/images/face/63.gif": assets_layui_images_face_63_gif,
	"assets/layui/images/face/64.gif": assets_layui_images_face_64_gif,
	"assets/layui/images/face/65.gif": assets_layui_images_face_65_gif,
	"assets/layui/images/face/66.gif": assets_layui_images_face_66_gif,
	"assets/layui/images/face/67.gif": assets_layui_images_face_67_gif,
	"assets/layui/images/face/68.gif": assets_layui_images_face_68_gif,
	"assets/layui/images/face/69.gif": assets_layui_images_face_69_gif,
	"assets/layui/images/face/7.gif": assets_layui_images_face_7_gif,
	"assets/layui/images/face/70.gif": assets_layui_images_face_70_gif,
	"assets/layui/images/face/71.gif": assets_layui_images_face_71_gif,
	"assets/layui/images/face/8.gif": assets_layui_images_face_8_gif,
	"assets/layui/images/face/9.gif": assets_layui_images_face_9_gif,
	"assets/layui/lay/modules/carousel.js": assets_layui_lay_modules_carousel_js,
	"assets/layui/lay/modules/code.js": assets_layui_lay_modules_code_js,
	"assets/layui/lay/modules/colorpicker.js": assets_layui_lay_modules_colorpicker_js,
	"assets/layui/lay/modules/element.js": assets_layui_lay_modules_element_js,
	"assets/layui/lay/modules/flow.js": assets_layui_lay_modules_flow_js,
	"assets/layui/lay/modules/form.js": assets_layui_lay_modules_form_js,
	"assets/layui/lay/modules/jquery.js": assets_layui_lay_modules_jquery_js,
	"assets/layui/lay/modules/laydate.js": assets_layui_lay_modules_laydate_js,
	"assets/layui/lay/modules/layedit.js": assets_layui_lay_modules_layedit_js,
	"assets/layui/lay/modules/layer.js": assets_layui_lay_modules_layer_js,
	"assets/layui/lay/modules/laypage.js": assets_layui_lay_modules_laypage_js,
	"assets/layui/lay/modules/laytpl.js": assets_layui_lay_modules_laytpl_js,
	"assets/layui/lay/modules/mobile.js": assets_layui_lay_modules_mobile_js,
	"assets/layui/lay/modules/rate.js": assets_layui_lay_modules_rate_js,
	"assets/layui/lay/modules/slider.js": assets_layui_lay_modules_slider_js,
	"assets/layui/lay/modules/table.js": assets_layui_lay_modules_table_js,
	"assets/layui/lay/modules/transfer.js": assets_layui_lay_modules_transfer_js,
	"assets/layui/lay/modules/tree.js": assets_layui_lay_modules_tree_js,
	"assets/layui/lay/modules/upload.js": assets_layui_lay_modules_upload_js,
	"assets/layui/lay/modules/util.js": assets_layui_lay_modules_util_js,
	"assets/layui/layui.all.js": assets_layui_layui_all_js,
	"assets/layui/layui.js": assets_layui_layui_js,
	"assets/login.html": assets_login_html,
	"assets/static/bootstrap/css/bootstrap.min.css": assets_static_bootstrap_css_bootstrap_min_css,
	"assets/static/bootstrap/css/bootstrap.min.css.map": assets_static_bootstrap_css_bootstrap_min_css_map,
	"assets/static/bootstrap/fonts/glyphicons-halflings-regular.eot": assets_static_bootstrap_fonts_glyphicons_halflings_regular_eot,
	"assets/static/bootstrap/fonts/glyphicons-halflings-regular.svg": assets_static_bootstrap_fonts_glyphicons_halflings_regular_svg,
	"assets/static/bootstrap/fonts/glyphicons-halflings-regular.ttf": assets_static_bootstrap_fonts_glyphicons_halflings_regular_ttf,
	"assets/static/bootstrap/fonts/glyphicons-halflings-regular.woff": assets_static_bootstrap_fonts_glyphicons_halflings_regular_woff,
	"assets/static/bootstrap/fonts/glyphicons-halflings-regular.woff2": assets_static_bootstrap_fonts_glyphicons_halflings_regular_woff2,
	"assets/static/bootstrap/js/bootstrap.min.js": assets_static_bootstrap_js_bootstrap_min_js,
	"assets/static/css/admin-login.css": assets_static_css_admin_login_css,
	"assets/static/css/admin.css": assets_static_css_admin_css,
	"assets/static/css/animate.css": assets_static_css_animate_css,
	"assets/static/font-awesome-4.7.0/HELP-US-OUT.txt": assets_static_font_awesome_4_7_0_help_us_out_txt,
	"assets/static/font-awesome-4.7.0/css/font-awesome.css": assets_static_font_awesome_4_7_0_css_font_awesome_css,
	"assets/static/font-awesome-4.7.0/css/font-awesome.min.css": assets_static_font_awesome_4_7_0_css_font_awesome_min_css,
	"assets/static/font-awesome-4.7.0/fonts/FontAwesome.otf": assets_static_font_awesome_4_7_0_fonts_fontawesome_otf,
	"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.eot": assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_eot,
	"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.svg": assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_svg,
	"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.ttf": assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_ttf,
	"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.woff": assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_woff,
	"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.woff2": assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_woff2,
	"assets/static/font-awesome-4.7.0/less/animated.less": assets_static_font_awesome_4_7_0_less_animated_less,
	"assets/static/font-awesome-4.7.0/less/bordered-pulled.less": assets_static_font_awesome_4_7_0_less_bordered_pulled_less,
	"assets/static/font-awesome-4.7.0/less/core.less": assets_static_font_awesome_4_7_0_less_core_less,
	"assets/static/font-awesome-4.7.0/less/fixed-width.less": assets_static_font_awesome_4_7_0_less_fixed_width_less,
	"assets/static/font-awesome-4.7.0/less/font-awesome.less": assets_static_font_awesome_4_7_0_less_font_awesome_less,
	"assets/static/font-awesome-4.7.0/less/icons.less": assets_static_font_awesome_4_7_0_less_icons_less,
	"assets/static/font-awesome-4.7.0/less/larger.less": assets_static_font_awesome_4_7_0_less_larger_less,
	"assets/static/font-awesome-4.7.0/less/list.less": assets_static_font_awesome_4_7_0_less_list_less,
	"assets/static/font-awesome-4.7.0/less/mixins.less": assets_static_font_awesome_4_7_0_less_mixins_less,
	"assets/static/font-awesome-4.7.0/less/path.less": assets_static_font_awesome_4_7_0_less_path_less,
	"assets/static/font-awesome-4.7.0/less/rotated-flipped.less": assets_static_font_awesome_4_7_0_less_rotated_flipped_less,
	"assets/static/font-awesome-4.7.0/less/screen-reader.less": assets_static_font_awesome_4_7_0_less_screen_reader_less,
	"assets/static/font-awesome-4.7.0/less/stacked.less": assets_static_font_awesome_4_7_0_less_stacked_less,
	"assets/static/font-awesome-4.7.0/less/variables.less": assets_static_font_awesome_4_7_0_less_variables_less,
	"assets/static/font-awesome-4.7.0/scss/_animated.scss": assets_static_font_awesome_4_7_0_scss_animated_scss,
	"assets/static/font-awesome-4.7.0/scss/_bordered-pulled.scss": assets_static_font_awesome_4_7_0_scss_bordered_pulled_scss,
	"assets/static/font-awesome-4.7.0/scss/_core.scss": assets_static_font_awesome_4_7_0_scss_core_scss,
	"assets/static/font-awesome-4.7.0/scss/_fixed-width.scss": assets_static_font_awesome_4_7_0_scss_fixed_width_scss,
	"assets/static/font-awesome-4.7.0/scss/_icons.scss": assets_static_font_awesome_4_7_0_scss_icons_scss,
	"assets/static/font-awesome-4.7.0/scss/_larger.scss": assets_static_font_awesome_4_7_0_scss_larger_scss,
	"assets/static/font-awesome-4.7.0/scss/_list.scss": assets_static_font_awesome_4_7_0_scss_list_scss,
	"assets/static/font-awesome-4.7.0/scss/_mixins.scss": assets_static_font_awesome_4_7_0_scss_mixins_scss,
	"assets/static/font-awesome-4.7.0/scss/_path.scss": assets_static_font_awesome_4_7_0_scss_path_scss,
	"assets/static/font-awesome-4.7.0/scss/_rotated-flipped.scss": assets_static_font_awesome_4_7_0_scss_rotated_flipped_scss,
	"assets/static/font-awesome-4.7.0/scss/_screen-reader.scss": assets_static_font_awesome_4_7_0_scss_screen_reader_scss,
	"assets/static/font-awesome-4.7.0/scss/_stacked.scss": assets_static_font_awesome_4_7_0_scss_stacked_scss,
	"assets/static/font-awesome-4.7.0/scss/_variables.scss": assets_static_font_awesome_4_7_0_scss_variables_scss,
	"assets/static/font-awesome-4.7.0/scss/font-awesome.scss": assets_static_font_awesome_4_7_0_scss_font_awesome_scss,
	"assets/static/images/bg_2.jpg": assets_static_images_bg_2_jpg,
	"assets/static/images/chahao.png": assets_static_images_chahao_png,
	"assets/static/images/duihao.png": assets_static_images_duihao_png,
	"assets/static/images/geometry2.png": assets_static_images_geometry2_png,
	"assets/static/images/hevr.png": assets_static_images_hevr_png,
	"assets/static/images/logo.png": assets_static_images_logo_png,
	"assets/static/images/qrcode.jpg": assets_static_images_qrcode_jpg,
	"assets/static/images/success.png": assets_static_images_success_png,
	"assets/static/images/upload.jpg": assets_static_images_upload_jpg,
	"assets/static/js/admin-login.js": assets_static_js_admin_login_js,
	"assets/static/js/admin.js": assets_static_js_admin_js,
	"assets/static/js/common.js": assets_static_js_common_js,
	"assets/static/js/echarts.min.js": assets_static_js_echarts_min_js,
	"assets/static/js/jquery.min.js": assets_static_js_jquery_min_js,
	"assets/static/js/jquery.placeholder.min.js": assets_static_js_jquery_placeholder_min_js,
	"assets/static/js/jquery.waypoints.min.js": assets_static_js_jquery_waypoints_min_js,
	"assets/static/js/modernizr-2.6.2.min.js": assets_static_js_modernizr_2_6_2_min_js,
	"assets/static/js/modernizr-custom-v2.7.1.min.js": assets_static_js_modernizr_custom_v2_7_1_min_js,
	"assets/static/layui/css/formSelects-v4.css": assets_static_layui_css_formselects_v4_css,
	"assets/static/layui/css/layui.css": assets_static_layui_css_layui_css,
	"assets/static/layui/css/layui.mobile.css": assets_static_layui_css_layui_mobile_css,
	"assets/static/layui/css/modules/code.css": assets_static_layui_css_modules_code_css,
	"assets/static/layui/css/modules/laydate/default/laydate.css": assets_static_layui_css_modules_laydate_default_laydate_css,
	"assets/static/layui/css/modules/layer/default/icon-ext.png": assets_static_layui_css_modules_layer_default_icon_ext_png,
	"assets/static/layui/css/modules/layer/default/icon.png": assets_static_layui_css_modules_layer_default_icon_png,
	"assets/static/layui/css/modules/layer/default/layer.css": assets_static_layui_css_modules_layer_default_layer_css,
	"assets/static/layui/css/modules/layer/default/loading-0.gif": assets_static_layui_css_modules_layer_default_loading_0_gif,
	"assets/static/layui/css/modules/layer/default/loading-1.gif": assets_static_layui_css_modules_layer_default_loading_1_gif,
	"assets/static/layui/css/modules/layer/default/loading-2.gif": assets_static_layui_css_modules_layer_default_loading_2_gif,
	"assets/static/layui/font/iconfont.eot": assets_static_layui_font_iconfont_eot,
	"assets/static/layui/font/iconfont.svg": assets_static_layui_font_iconfont_svg,
	"assets/static/layui/font/iconfont.ttf": assets_static_layui_font_iconfont_ttf,
	"assets/static/layui/font/iconfont.woff": assets_static_layui_font_iconfont_woff,
	"assets/static/layui/font/iconfont.woff2": assets_static_layui_font_iconfont_woff2,
	"assets/static/layui/formSelects-v4.js": assets_static_layui_formselects_v4_js,
	"assets/static/layui/images/face/0.gif": assets_static_layui_images_face_0_gif,
	"assets/static/layui/images/face/1.gif": assets_static_layui_images_face_1_gif,
	"assets/static/layui/images/face/10.gif": assets_static_layui_images_face_10_gif,
	"assets/static/layui/images/face/11.gif": assets_static_layui_images_face_11_gif,
	"assets/static/layui/images/face/12.gif": assets_static_layui_images_face_12_gif,
	"assets/static/layui/images/face/13.gif": assets_static_layui_images_face_13_gif,
	"assets/static/layui/images/face/14.gif": assets_static_layui_images_face_14_gif,
	"assets/static/layui/images/face/15.gif": assets_static_layui_images_face_15_gif,
	"assets/static/layui/images/face/16.gif": assets_static_layui_images_face_16_gif,
	"assets/static/layui/images/face/17.gif": assets_static_layui_images_face_17_gif,
	"assets/static/layui/images/face/18.gif": assets_static_layui_images_face_18_gif,
	"assets/static/layui/images/face/19.gif": assets_static_layui_images_face_19_gif,
	"assets/static/layui/images/face/2.gif": assets_static_layui_images_face_2_gif,
	"assets/static/layui/images/face/20.gif": assets_static_layui_images_face_20_gif,
	"assets/static/layui/images/face/21.gif": assets_static_layui_images_face_21_gif,
	"assets/static/layui/images/face/22.gif": assets_static_layui_images_face_22_gif,
	"assets/static/layui/images/face/23.gif": assets_static_layui_images_face_23_gif,
	"assets/static/layui/images/face/24.gif": assets_static_layui_images_face_24_gif,
	"assets/static/layui/images/face/25.gif": assets_static_layui_images_face_25_gif,
	"assets/static/layui/images/face/26.gif": assets_static_layui_images_face_26_gif,
	"assets/static/layui/images/face/27.gif": assets_static_layui_images_face_27_gif,
	"assets/static/layui/images/face/28.gif": assets_static_layui_images_face_28_gif,
	"assets/static/layui/images/face/29.gif": assets_static_layui_images_face_29_gif,
	"assets/static/layui/images/face/3.gif": assets_static_layui_images_face_3_gif,
	"assets/static/layui/images/face/30.gif": assets_static_layui_images_face_30_gif,
	"assets/static/layui/images/face/31.gif": assets_static_layui_images_face_31_gif,
	"assets/static/layui/images/face/32.gif": assets_static_layui_images_face_32_gif,
	"assets/static/layui/images/face/33.gif": assets_static_layui_images_face_33_gif,
	"assets/static/layui/images/face/34.gif": assets_static_layui_images_face_34_gif,
	"assets/static/layui/images/face/35.gif": assets_static_layui_images_face_35_gif,
	"assets/static/layui/images/face/36.gif": assets_static_layui_images_face_36_gif,
	"assets/static/layui/images/face/37.gif": assets_static_layui_images_face_37_gif,
	"assets/static/layui/images/face/38.gif": assets_static_layui_images_face_38_gif,
	"assets/static/layui/images/face/39.gif": assets_static_layui_images_face_39_gif,
	"assets/static/layui/images/face/4.gif": assets_static_layui_images_face_4_gif,
	"assets/static/layui/images/face/40.gif": assets_static_layui_images_face_40_gif,
	"assets/static/layui/images/face/41.gif": assets_static_layui_images_face_41_gif,
	"assets/static/layui/images/face/42.gif": assets_static_layui_images_face_42_gif,
	"assets/static/layui/images/face/43.gif": assets_static_layui_images_face_43_gif,
	"assets/static/layui/images/face/44.gif": assets_static_layui_images_face_44_gif,
	"assets/static/layui/images/face/45.gif": assets_static_layui_images_face_45_gif,
	"assets/static/layui/images/face/46.gif": assets_static_layui_images_face_46_gif,
	"assets/static/layui/images/face/47.gif": assets_static_layui_images_face_47_gif,
	"assets/static/layui/images/face/48.gif": assets_static_layui_images_face_48_gif,
	"assets/static/layui/images/face/49.gif": assets_static_layui_images_face_49_gif,
	"assets/static/layui/images/face/5.gif": assets_static_layui_images_face_5_gif,
	"assets/static/layui/images/face/50.gif": assets_static_layui_images_face_50_gif,
	"assets/static/layui/images/face/51.gif": assets_static_layui_images_face_51_gif,
	"assets/static/layui/images/face/52.gif": assets_static_layui_images_face_52_gif,
	"assets/static/layui/images/face/53.gif": assets_static_layui_images_face_53_gif,
	"assets/static/layui/images/face/54.gif": assets_static_layui_images_face_54_gif,
	"assets/static/layui/images/face/55.gif": assets_static_layui_images_face_55_gif,
	"assets/static/layui/images/face/56.gif": assets_static_layui_images_face_56_gif,
	"assets/static/layui/images/face/57.gif": assets_static_layui_images_face_57_gif,
	"assets/static/layui/images/face/58.gif": assets_static_layui_images_face_58_gif,
	"assets/static/layui/images/face/59.gif": assets_static_layui_images_face_59_gif,
	"assets/static/layui/images/face/6.gif": assets_static_layui_images_face_6_gif,
	"assets/static/layui/images/face/60.gif": assets_static_layui_images_face_60_gif,
	"assets/static/layui/images/face/61.gif": assets_static_layui_images_face_61_gif,
	"assets/static/layui/images/face/62.gif": assets_static_layui_images_face_62_gif,
	"assets/static/layui/images/face/63.gif": assets_static_layui_images_face_63_gif,
	"assets/static/layui/images/face/64.gif": assets_static_layui_images_face_64_gif,
	"assets/static/layui/images/face/65.gif": assets_static_layui_images_face_65_gif,
	"assets/static/layui/images/face/66.gif": assets_static_layui_images_face_66_gif,
	"assets/static/layui/images/face/67.gif": assets_static_layui_images_face_67_gif,
	"assets/static/layui/images/face/68.gif": assets_static_layui_images_face_68_gif,
	"assets/static/layui/images/face/69.gif": assets_static_layui_images_face_69_gif,
	"assets/static/layui/images/face/7.gif": assets_static_layui_images_face_7_gif,
	"assets/static/layui/images/face/70.gif": assets_static_layui_images_face_70_gif,
	"assets/static/layui/images/face/71.gif": assets_static_layui_images_face_71_gif,
	"assets/static/layui/images/face/8.gif": assets_static_layui_images_face_8_gif,
	"assets/static/layui/images/face/9.gif": assets_static_layui_images_face_9_gif,
	"assets/static/layui/lay/modules/carousel.js": assets_static_layui_lay_modules_carousel_js,
	"assets/static/layui/lay/modules/code.js": assets_static_layui_lay_modules_code_js,
	"assets/static/layui/lay/modules/colorpicker.js": assets_static_layui_lay_modules_colorpicker_js,
	"assets/static/layui/lay/modules/element.js": assets_static_layui_lay_modules_element_js,
	"assets/static/layui/lay/modules/flow.js": assets_static_layui_lay_modules_flow_js,
	"assets/static/layui/lay/modules/form.js": assets_static_layui_lay_modules_form_js,
	"assets/static/layui/lay/modules/jquery.js": assets_static_layui_lay_modules_jquery_js,
	"assets/static/layui/lay/modules/laydate.js": assets_static_layui_lay_modules_laydate_js,
	"assets/static/layui/lay/modules/layedit.js": assets_static_layui_lay_modules_layedit_js,
	"assets/static/layui/lay/modules/layer.js": assets_static_layui_lay_modules_layer_js,
	"assets/static/layui/lay/modules/laypage.js": assets_static_layui_lay_modules_laypage_js,
	"assets/static/layui/lay/modules/laytpl.js": assets_static_layui_lay_modules_laytpl_js,
	"assets/static/layui/lay/modules/mobile.js": assets_static_layui_lay_modules_mobile_js,
	"assets/static/layui/lay/modules/rate.js": assets_static_layui_lay_modules_rate_js,
	"assets/static/layui/lay/modules/slider.js": assets_static_layui_lay_modules_slider_js,
	"assets/static/layui/lay/modules/table.js": assets_static_layui_lay_modules_table_js,
	"assets/static/layui/lay/modules/transfer.js": assets_static_layui_lay_modules_transfer_js,
	"assets/static/layui/lay/modules/tree.js": assets_static_layui_lay_modules_tree_js,
	"assets/static/layui/lay/modules/upload.js": assets_static_layui_lay_modules_upload_js,
	"assets/static/layui/lay/modules/util.js": assets_static_layui_lay_modules_util_js,
	"assets/static/layui/layui.all.js": assets_static_layui_layui_all_js,
	"assets/static/layui/layui.js": assets_static_layui_layui_js,
	"assets/static/upload/imgs/20190905/1567651858020_647.png": assets_static_upload_imgs_20190905_1567651858020_647_png,
	"assets/upload/imgs/20190905/1567651858020_647.png": assets_upload_imgs_20190905_1567651858020_647_png,
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
	"assets/bootstrap/css/bootstrap.min.css": &_bintree_t{assets_bootstrap_css_bootstrap_min_css, map[string]*_bintree_t{
	}},
	"assets/bootstrap/css/bootstrap.min.css.map": &_bintree_t{assets_bootstrap_css_bootstrap_min_css_map, map[string]*_bintree_t{
	}},
	"assets/bootstrap/fonts/glyphicons-halflings-regular.eot": &_bintree_t{assets_bootstrap_fonts_glyphicons_halflings_regular_eot, map[string]*_bintree_t{
	}},
	"assets/bootstrap/fonts/glyphicons-halflings-regular.svg": &_bintree_t{assets_bootstrap_fonts_glyphicons_halflings_regular_svg, map[string]*_bintree_t{
	}},
	"assets/bootstrap/fonts/glyphicons-halflings-regular.ttf": &_bintree_t{assets_bootstrap_fonts_glyphicons_halflings_regular_ttf, map[string]*_bintree_t{
	}},
	"assets/bootstrap/fonts/glyphicons-halflings-regular.woff": &_bintree_t{assets_bootstrap_fonts_glyphicons_halflings_regular_woff, map[string]*_bintree_t{
	}},
	"assets/bootstrap/fonts/glyphicons-halflings-regular.woff2": &_bintree_t{assets_bootstrap_fonts_glyphicons_halflings_regular_woff2, map[string]*_bintree_t{
	}},
	"assets/bootstrap/js/bootstrap.min.js": &_bintree_t{assets_bootstrap_js_bootstrap_min_js, map[string]*_bintree_t{
	}},
	"assets/css/admin-login.css": &_bintree_t{assets_css_admin_login_css, map[string]*_bintree_t{
	}},
	"assets/css/admin.css": &_bintree_t{assets_css_admin_css, map[string]*_bintree_t{
	}},
	"assets/css/animate.css": &_bintree_t{assets_css_animate_css, map[string]*_bintree_t{
	}},
	"assets/css/bootstrap-combined.min.css": &_bintree_t{assets_css_bootstrap_combined_min_css, map[string]*_bintree_t{
	}},
	"assets/css/bootstrap-theme.css": &_bintree_t{assets_css_bootstrap_theme_css, map[string]*_bintree_t{
	}},
	"assets/css/bootstrap-theme.css.map": &_bintree_t{assets_css_bootstrap_theme_css_map, map[string]*_bintree_t{
	}},
	"assets/css/bootstrap-theme.min.css": &_bintree_t{assets_css_bootstrap_theme_min_css, map[string]*_bintree_t{
	}},
	"assets/css/bootstrap.css": &_bintree_t{assets_css_bootstrap_css, map[string]*_bintree_t{
	}},
	"assets/css/bootstrap.css.map": &_bintree_t{assets_css_bootstrap_css_map, map[string]*_bintree_t{
	}},
	"assets/css/bootstrap.min.css": &_bintree_t{assets_css_bootstrap_min_css, map[string]*_bintree_t{
	}},
	"assets/css/jquery.treetable.css": &_bintree_t{assets_css_jquery_treetable_css, map[string]*_bintree_t{
	}},
	"assets/css/jquery.treetable.theme.default.css": &_bintree_t{assets_css_jquery_treetable_theme_default_css, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/HELP-US-OUT.txt": &_bintree_t{assets_font_awesome_4_7_0_help_us_out_txt, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/css/font-awesome.css": &_bintree_t{assets_font_awesome_4_7_0_css_font_awesome_css, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/css/font-awesome.min.css": &_bintree_t{assets_font_awesome_4_7_0_css_font_awesome_min_css, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/fonts/FontAwesome.otf": &_bintree_t{assets_font_awesome_4_7_0_fonts_fontawesome_otf, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.eot": &_bintree_t{assets_font_awesome_4_7_0_fonts_fontawesome_webfont_eot, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.svg": &_bintree_t{assets_font_awesome_4_7_0_fonts_fontawesome_webfont_svg, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.ttf": &_bintree_t{assets_font_awesome_4_7_0_fonts_fontawesome_webfont_ttf, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.woff": &_bintree_t{assets_font_awesome_4_7_0_fonts_fontawesome_webfont_woff, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/fonts/fontawesome-webfont.woff2": &_bintree_t{assets_font_awesome_4_7_0_fonts_fontawesome_webfont_woff2, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/animated.less": &_bintree_t{assets_font_awesome_4_7_0_less_animated_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/bordered-pulled.less": &_bintree_t{assets_font_awesome_4_7_0_less_bordered_pulled_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/core.less": &_bintree_t{assets_font_awesome_4_7_0_less_core_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/fixed-width.less": &_bintree_t{assets_font_awesome_4_7_0_less_fixed_width_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/font-awesome.less": &_bintree_t{assets_font_awesome_4_7_0_less_font_awesome_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/icons.less": &_bintree_t{assets_font_awesome_4_7_0_less_icons_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/larger.less": &_bintree_t{assets_font_awesome_4_7_0_less_larger_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/list.less": &_bintree_t{assets_font_awesome_4_7_0_less_list_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/mixins.less": &_bintree_t{assets_font_awesome_4_7_0_less_mixins_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/path.less": &_bintree_t{assets_font_awesome_4_7_0_less_path_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/rotated-flipped.less": &_bintree_t{assets_font_awesome_4_7_0_less_rotated_flipped_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/screen-reader.less": &_bintree_t{assets_font_awesome_4_7_0_less_screen_reader_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/stacked.less": &_bintree_t{assets_font_awesome_4_7_0_less_stacked_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/less/variables.less": &_bintree_t{assets_font_awesome_4_7_0_less_variables_less, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/_animated.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_animated_scss, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/_bordered-pulled.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_bordered_pulled_scss, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/_core.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_core_scss, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/_fixed-width.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_fixed_width_scss, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/_icons.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_icons_scss, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/_larger.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_larger_scss, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/_list.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_list_scss, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/_mixins.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_mixins_scss, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/_path.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_path_scss, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/_rotated-flipped.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_rotated_flipped_scss, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/_screen-reader.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_screen_reader_scss, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/_stacked.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_stacked_scss, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/_variables.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_variables_scss, map[string]*_bintree_t{
	}},
	"assets/font-awesome-4.7.0/scss/font-awesome.scss": &_bintree_t{assets_font_awesome_4_7_0_scss_font_awesome_scss, map[string]*_bintree_t{
	}},
	"assets/fonts/glyphicons-halflings-regular.eot": &_bintree_t{assets_fonts_glyphicons_halflings_regular_eot, map[string]*_bintree_t{
	}},
	"assets/fonts/glyphicons-halflings-regular.svg": &_bintree_t{assets_fonts_glyphicons_halflings_regular_svg, map[string]*_bintree_t{
	}},
	"assets/fonts/glyphicons-halflings-regular.ttf": &_bintree_t{assets_fonts_glyphicons_halflings_regular_ttf, map[string]*_bintree_t{
	}},
	"assets/fonts/glyphicons-halflings-regular.woff": &_bintree_t{assets_fonts_glyphicons_halflings_regular_woff, map[string]*_bintree_t{
	}},
	"assets/fonts/glyphicons-halflings-regular.woff2": &_bintree_t{assets_fonts_glyphicons_halflings_regular_woff2, map[string]*_bintree_t{
	}},
	"assets/images/bg_2.jpg": &_bintree_t{assets_images_bg_2_jpg, map[string]*_bintree_t{
	}},
	"assets/images/chahao.png": &_bintree_t{assets_images_chahao_png, map[string]*_bintree_t{
	}},
	"assets/images/duihao.png": &_bintree_t{assets_images_duihao_png, map[string]*_bintree_t{
	}},
	"assets/images/geometry2.png": &_bintree_t{assets_images_geometry2_png, map[string]*_bintree_t{
	}},
	"assets/images/hevr.png": &_bintree_t{assets_images_hevr_png, map[string]*_bintree_t{
	}},
	"assets/images/logo.png": &_bintree_t{assets_images_logo_png, map[string]*_bintree_t{
	}},
	"assets/images/qrcode.jpg": &_bintree_t{assets_images_qrcode_jpg, map[string]*_bintree_t{
	}},
	"assets/images/success.png": &_bintree_t{assets_images_success_png, map[string]*_bintree_t{
	}},
	"assets/images/upload.jpg": &_bintree_t{assets_images_upload_jpg, map[string]*_bintree_t{
	}},
	"assets/img/favicon.ico": &_bintree_t{assets_img_favicon_ico, map[string]*_bintree_t{
	}},
	"assets/img/icon_brand.png": &_bintree_t{assets_img_icon_brand_png, map[string]*_bintree_t{
	}},
	"assets/js/admin-login.js": &_bintree_t{assets_js_admin_login_js, map[string]*_bintree_t{
	}},
	"assets/js/admin.js": &_bintree_t{assets_js_admin_js, map[string]*_bintree_t{
	}},
	"assets/js/app/jwt/jwt.js": &_bintree_t{assets_js_app_jwt_jwt_js, map[string]*_bintree_t{
	}},
	"assets/js/app/treetable/treetable-page.js": &_bintree_t{assets_js_app_treetable_treetable_page_js, map[string]*_bintree_t{
	}},
	"assets/js/common.js": &_bintree_t{assets_js_common_js, map[string]*_bintree_t{
	}},
	"assets/js/echarts.min.js": &_bintree_t{assets_js_echarts_min_js, map[string]*_bintree_t{
	}},
	"assets/js/jquery.min.js": &_bintree_t{assets_js_jquery_min_js, map[string]*_bintree_t{
	}},
	"assets/js/jquery.placeholder.min.js": &_bintree_t{assets_js_jquery_placeholder_min_js, map[string]*_bintree_t{
	}},
	"assets/js/jquery.waypoints.min.js": &_bintree_t{assets_js_jquery_waypoints_min_js, map[string]*_bintree_t{
	}},
	"assets/js/lib/bootstrap/bootstrap.js": &_bintree_t{assets_js_lib_bootstrap_bootstrap_js, map[string]*_bintree_t{
	}},
	"assets/js/lib/bootstrap/bootstrap.min.js": &_bintree_t{assets_js_lib_bootstrap_bootstrap_min_js, map[string]*_bintree_t{
	}},
	"assets/js/lib/bootstrap/npm.js": &_bintree_t{assets_js_lib_bootstrap_npm_js, map[string]*_bintree_t{
	}},
	"assets/js/lib/jquery-form/jquery.form.3.51.js": &_bintree_t{assets_js_lib_jquery_form_jquery_form_3_51_js, map[string]*_bintree_t{
	}},
	"assets/js/lib/jquery-treetable/jquery.treetable.js": &_bintree_t{assets_js_lib_jquery_treetable_jquery_treetable_js, map[string]*_bintree_t{
	}},
	"assets/js/lib/jquery/jquery-2.1.3.js": &_bintree_t{assets_js_lib_jquery_jquery_2_1_3_js, map[string]*_bintree_t{
	}},
	"assets/js/lib/jquery/jquery-2.1.3.min.js": &_bintree_t{assets_js_lib_jquery_jquery_2_1_3_min_js, map[string]*_bintree_t{
	}},
	"assets/js/lib/require-setup.js": &_bintree_t{assets_js_lib_require_setup_js, map[string]*_bintree_t{
	}},
	"assets/js/lib/require.js": &_bintree_t{assets_js_lib_require_js, map[string]*_bintree_t{
	}},
	"assets/js/modernizr-2.6.2.min.js": &_bintree_t{assets_js_modernizr_2_6_2_min_js, map[string]*_bintree_t{
	}},
	"assets/js/modernizr-custom-v2.7.1.min.js": &_bintree_t{assets_js_modernizr_custom_v2_7_1_min_js, map[string]*_bintree_t{
	}},
	"assets/js/tools/r.js": &_bintree_t{assets_js_tools_r_js, map[string]*_bintree_t{
	}},
	"assets/layui/css/layui.css": &_bintree_t{assets_layui_css_layui_css, map[string]*_bintree_t{
	}},
	"assets/layui/css/layui.mobile.css": &_bintree_t{assets_layui_css_layui_mobile_css, map[string]*_bintree_t{
	}},
	"assets/layui/css/modules/code.css": &_bintree_t{assets_layui_css_modules_code_css, map[string]*_bintree_t{
	}},
	"assets/layui/css/modules/laydate/default/laydate.css": &_bintree_t{assets_layui_css_modules_laydate_default_laydate_css, map[string]*_bintree_t{
	}},
	"assets/layui/css/modules/layer/default/icon-ext.png": &_bintree_t{assets_layui_css_modules_layer_default_icon_ext_png, map[string]*_bintree_t{
	}},
	"assets/layui/css/modules/layer/default/icon.png": &_bintree_t{assets_layui_css_modules_layer_default_icon_png, map[string]*_bintree_t{
	}},
	"assets/layui/css/modules/layer/default/layer.css": &_bintree_t{assets_layui_css_modules_layer_default_layer_css, map[string]*_bintree_t{
	}},
	"assets/layui/css/modules/layer/default/loading-0.gif": &_bintree_t{assets_layui_css_modules_layer_default_loading_0_gif, map[string]*_bintree_t{
	}},
	"assets/layui/css/modules/layer/default/loading-1.gif": &_bintree_t{assets_layui_css_modules_layer_default_loading_1_gif, map[string]*_bintree_t{
	}},
	"assets/layui/css/modules/layer/default/loading-2.gif": &_bintree_t{assets_layui_css_modules_layer_default_loading_2_gif, map[string]*_bintree_t{
	}},
	"assets/layui/font/iconfont.eot": &_bintree_t{assets_layui_font_iconfont_eot, map[string]*_bintree_t{
	}},
	"assets/layui/font/iconfont.svg": &_bintree_t{assets_layui_font_iconfont_svg, map[string]*_bintree_t{
	}},
	"assets/layui/font/iconfont.ttf": &_bintree_t{assets_layui_font_iconfont_ttf, map[string]*_bintree_t{
	}},
	"assets/layui/font/iconfont.woff": &_bintree_t{assets_layui_font_iconfont_woff, map[string]*_bintree_t{
	}},
	"assets/layui/font/iconfont.woff2": &_bintree_t{assets_layui_font_iconfont_woff2, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/0.gif": &_bintree_t{assets_layui_images_face_0_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/1.gif": &_bintree_t{assets_layui_images_face_1_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/10.gif": &_bintree_t{assets_layui_images_face_10_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/11.gif": &_bintree_t{assets_layui_images_face_11_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/12.gif": &_bintree_t{assets_layui_images_face_12_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/13.gif": &_bintree_t{assets_layui_images_face_13_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/14.gif": &_bintree_t{assets_layui_images_face_14_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/15.gif": &_bintree_t{assets_layui_images_face_15_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/16.gif": &_bintree_t{assets_layui_images_face_16_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/17.gif": &_bintree_t{assets_layui_images_face_17_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/18.gif": &_bintree_t{assets_layui_images_face_18_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/19.gif": &_bintree_t{assets_layui_images_face_19_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/2.gif": &_bintree_t{assets_layui_images_face_2_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/20.gif": &_bintree_t{assets_layui_images_face_20_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/21.gif": &_bintree_t{assets_layui_images_face_21_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/22.gif": &_bintree_t{assets_layui_images_face_22_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/23.gif": &_bintree_t{assets_layui_images_face_23_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/24.gif": &_bintree_t{assets_layui_images_face_24_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/25.gif": &_bintree_t{assets_layui_images_face_25_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/26.gif": &_bintree_t{assets_layui_images_face_26_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/27.gif": &_bintree_t{assets_layui_images_face_27_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/28.gif": &_bintree_t{assets_layui_images_face_28_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/29.gif": &_bintree_t{assets_layui_images_face_29_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/3.gif": &_bintree_t{assets_layui_images_face_3_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/30.gif": &_bintree_t{assets_layui_images_face_30_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/31.gif": &_bintree_t{assets_layui_images_face_31_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/32.gif": &_bintree_t{assets_layui_images_face_32_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/33.gif": &_bintree_t{assets_layui_images_face_33_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/34.gif": &_bintree_t{assets_layui_images_face_34_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/35.gif": &_bintree_t{assets_layui_images_face_35_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/36.gif": &_bintree_t{assets_layui_images_face_36_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/37.gif": &_bintree_t{assets_layui_images_face_37_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/38.gif": &_bintree_t{assets_layui_images_face_38_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/39.gif": &_bintree_t{assets_layui_images_face_39_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/4.gif": &_bintree_t{assets_layui_images_face_4_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/40.gif": &_bintree_t{assets_layui_images_face_40_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/41.gif": &_bintree_t{assets_layui_images_face_41_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/42.gif": &_bintree_t{assets_layui_images_face_42_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/43.gif": &_bintree_t{assets_layui_images_face_43_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/44.gif": &_bintree_t{assets_layui_images_face_44_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/45.gif": &_bintree_t{assets_layui_images_face_45_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/46.gif": &_bintree_t{assets_layui_images_face_46_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/47.gif": &_bintree_t{assets_layui_images_face_47_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/48.gif": &_bintree_t{assets_layui_images_face_48_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/49.gif": &_bintree_t{assets_layui_images_face_49_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/5.gif": &_bintree_t{assets_layui_images_face_5_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/50.gif": &_bintree_t{assets_layui_images_face_50_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/51.gif": &_bintree_t{assets_layui_images_face_51_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/52.gif": &_bintree_t{assets_layui_images_face_52_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/53.gif": &_bintree_t{assets_layui_images_face_53_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/54.gif": &_bintree_t{assets_layui_images_face_54_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/55.gif": &_bintree_t{assets_layui_images_face_55_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/56.gif": &_bintree_t{assets_layui_images_face_56_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/57.gif": &_bintree_t{assets_layui_images_face_57_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/58.gif": &_bintree_t{assets_layui_images_face_58_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/59.gif": &_bintree_t{assets_layui_images_face_59_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/6.gif": &_bintree_t{assets_layui_images_face_6_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/60.gif": &_bintree_t{assets_layui_images_face_60_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/61.gif": &_bintree_t{assets_layui_images_face_61_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/62.gif": &_bintree_t{assets_layui_images_face_62_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/63.gif": &_bintree_t{assets_layui_images_face_63_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/64.gif": &_bintree_t{assets_layui_images_face_64_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/65.gif": &_bintree_t{assets_layui_images_face_65_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/66.gif": &_bintree_t{assets_layui_images_face_66_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/67.gif": &_bintree_t{assets_layui_images_face_67_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/68.gif": &_bintree_t{assets_layui_images_face_68_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/69.gif": &_bintree_t{assets_layui_images_face_69_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/7.gif": &_bintree_t{assets_layui_images_face_7_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/70.gif": &_bintree_t{assets_layui_images_face_70_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/71.gif": &_bintree_t{assets_layui_images_face_71_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/8.gif": &_bintree_t{assets_layui_images_face_8_gif, map[string]*_bintree_t{
	}},
	"assets/layui/images/face/9.gif": &_bintree_t{assets_layui_images_face_9_gif, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/carousel.js": &_bintree_t{assets_layui_lay_modules_carousel_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/code.js": &_bintree_t{assets_layui_lay_modules_code_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/colorpicker.js": &_bintree_t{assets_layui_lay_modules_colorpicker_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/element.js": &_bintree_t{assets_layui_lay_modules_element_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/flow.js": &_bintree_t{assets_layui_lay_modules_flow_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/form.js": &_bintree_t{assets_layui_lay_modules_form_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/jquery.js": &_bintree_t{assets_layui_lay_modules_jquery_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/laydate.js": &_bintree_t{assets_layui_lay_modules_laydate_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/layedit.js": &_bintree_t{assets_layui_lay_modules_layedit_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/layer.js": &_bintree_t{assets_layui_lay_modules_layer_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/laypage.js": &_bintree_t{assets_layui_lay_modules_laypage_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/laytpl.js": &_bintree_t{assets_layui_lay_modules_laytpl_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/mobile.js": &_bintree_t{assets_layui_lay_modules_mobile_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/rate.js": &_bintree_t{assets_layui_lay_modules_rate_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/slider.js": &_bintree_t{assets_layui_lay_modules_slider_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/table.js": &_bintree_t{assets_layui_lay_modules_table_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/transfer.js": &_bintree_t{assets_layui_lay_modules_transfer_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/tree.js": &_bintree_t{assets_layui_lay_modules_tree_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/upload.js": &_bintree_t{assets_layui_lay_modules_upload_js, map[string]*_bintree_t{
	}},
	"assets/layui/lay/modules/util.js": &_bintree_t{assets_layui_lay_modules_util_js, map[string]*_bintree_t{
	}},
	"assets/layui/layui.all.js": &_bintree_t{assets_layui_layui_all_js, map[string]*_bintree_t{
	}},
	"assets/layui/layui.js": &_bintree_t{assets_layui_layui_js, map[string]*_bintree_t{
	}},
	"assets/login.html": &_bintree_t{assets_login_html, map[string]*_bintree_t{
	}},
	"assets/static/bootstrap/css/bootstrap.min.css": &_bintree_t{assets_static_bootstrap_css_bootstrap_min_css, map[string]*_bintree_t{
	}},
	"assets/static/bootstrap/css/bootstrap.min.css.map": &_bintree_t{assets_static_bootstrap_css_bootstrap_min_css_map, map[string]*_bintree_t{
	}},
	"assets/static/bootstrap/fonts/glyphicons-halflings-regular.eot": &_bintree_t{assets_static_bootstrap_fonts_glyphicons_halflings_regular_eot, map[string]*_bintree_t{
	}},
	"assets/static/bootstrap/fonts/glyphicons-halflings-regular.svg": &_bintree_t{assets_static_bootstrap_fonts_glyphicons_halflings_regular_svg, map[string]*_bintree_t{
	}},
	"assets/static/bootstrap/fonts/glyphicons-halflings-regular.ttf": &_bintree_t{assets_static_bootstrap_fonts_glyphicons_halflings_regular_ttf, map[string]*_bintree_t{
	}},
	"assets/static/bootstrap/fonts/glyphicons-halflings-regular.woff": &_bintree_t{assets_static_bootstrap_fonts_glyphicons_halflings_regular_woff, map[string]*_bintree_t{
	}},
	"assets/static/bootstrap/fonts/glyphicons-halflings-regular.woff2": &_bintree_t{assets_static_bootstrap_fonts_glyphicons_halflings_regular_woff2, map[string]*_bintree_t{
	}},
	"assets/static/bootstrap/js/bootstrap.min.js": &_bintree_t{assets_static_bootstrap_js_bootstrap_min_js, map[string]*_bintree_t{
	}},
	"assets/static/css/admin-login.css": &_bintree_t{assets_static_css_admin_login_css, map[string]*_bintree_t{
	}},
	"assets/static/css/admin.css": &_bintree_t{assets_static_css_admin_css, map[string]*_bintree_t{
	}},
	"assets/static/css/animate.css": &_bintree_t{assets_static_css_animate_css, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/HELP-US-OUT.txt": &_bintree_t{assets_static_font_awesome_4_7_0_help_us_out_txt, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/css/font-awesome.css": &_bintree_t{assets_static_font_awesome_4_7_0_css_font_awesome_css, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/css/font-awesome.min.css": &_bintree_t{assets_static_font_awesome_4_7_0_css_font_awesome_min_css, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/fonts/FontAwesome.otf": &_bintree_t{assets_static_font_awesome_4_7_0_fonts_fontawesome_otf, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.eot": &_bintree_t{assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_eot, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.svg": &_bintree_t{assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_svg, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.ttf": &_bintree_t{assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_ttf, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.woff": &_bintree_t{assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_woff, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/fonts/fontawesome-webfont.woff2": &_bintree_t{assets_static_font_awesome_4_7_0_fonts_fontawesome_webfont_woff2, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/animated.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_animated_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/bordered-pulled.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_bordered_pulled_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/core.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_core_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/fixed-width.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_fixed_width_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/font-awesome.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_font_awesome_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/icons.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_icons_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/larger.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_larger_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/list.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_list_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/mixins.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_mixins_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/path.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_path_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/rotated-flipped.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_rotated_flipped_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/screen-reader.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_screen_reader_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/stacked.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_stacked_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/less/variables.less": &_bintree_t{assets_static_font_awesome_4_7_0_less_variables_less, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/_animated.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_animated_scss, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/_bordered-pulled.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_bordered_pulled_scss, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/_core.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_core_scss, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/_fixed-width.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_fixed_width_scss, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/_icons.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_icons_scss, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/_larger.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_larger_scss, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/_list.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_list_scss, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/_mixins.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_mixins_scss, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/_path.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_path_scss, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/_rotated-flipped.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_rotated_flipped_scss, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/_screen-reader.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_screen_reader_scss, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/_stacked.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_stacked_scss, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/_variables.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_variables_scss, map[string]*_bintree_t{
	}},
	"assets/static/font-awesome-4.7.0/scss/font-awesome.scss": &_bintree_t{assets_static_font_awesome_4_7_0_scss_font_awesome_scss, map[string]*_bintree_t{
	}},
	"assets/static/images/bg_2.jpg": &_bintree_t{assets_static_images_bg_2_jpg, map[string]*_bintree_t{
	}},
	"assets/static/images/chahao.png": &_bintree_t{assets_static_images_chahao_png, map[string]*_bintree_t{
	}},
	"assets/static/images/duihao.png": &_bintree_t{assets_static_images_duihao_png, map[string]*_bintree_t{
	}},
	"assets/static/images/geometry2.png": &_bintree_t{assets_static_images_geometry2_png, map[string]*_bintree_t{
	}},
	"assets/static/images/hevr.png": &_bintree_t{assets_static_images_hevr_png, map[string]*_bintree_t{
	}},
	"assets/static/images/logo.png": &_bintree_t{assets_static_images_logo_png, map[string]*_bintree_t{
	}},
	"assets/static/images/qrcode.jpg": &_bintree_t{assets_static_images_qrcode_jpg, map[string]*_bintree_t{
	}},
	"assets/static/images/success.png": &_bintree_t{assets_static_images_success_png, map[string]*_bintree_t{
	}},
	"assets/static/images/upload.jpg": &_bintree_t{assets_static_images_upload_jpg, map[string]*_bintree_t{
	}},
	"assets/static/js/admin-login.js": &_bintree_t{assets_static_js_admin_login_js, map[string]*_bintree_t{
	}},
	"assets/static/js/admin.js": &_bintree_t{assets_static_js_admin_js, map[string]*_bintree_t{
	}},
	"assets/static/js/common.js": &_bintree_t{assets_static_js_common_js, map[string]*_bintree_t{
	}},
	"assets/static/js/echarts.min.js": &_bintree_t{assets_static_js_echarts_min_js, map[string]*_bintree_t{
	}},
	"assets/static/js/jquery.min.js": &_bintree_t{assets_static_js_jquery_min_js, map[string]*_bintree_t{
	}},
	"assets/static/js/jquery.placeholder.min.js": &_bintree_t{assets_static_js_jquery_placeholder_min_js, map[string]*_bintree_t{
	}},
	"assets/static/js/jquery.waypoints.min.js": &_bintree_t{assets_static_js_jquery_waypoints_min_js, map[string]*_bintree_t{
	}},
	"assets/static/js/modernizr-2.6.2.min.js": &_bintree_t{assets_static_js_modernizr_2_6_2_min_js, map[string]*_bintree_t{
	}},
	"assets/static/js/modernizr-custom-v2.7.1.min.js": &_bintree_t{assets_static_js_modernizr_custom_v2_7_1_min_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/css/formSelects-v4.css": &_bintree_t{assets_static_layui_css_formselects_v4_css, map[string]*_bintree_t{
	}},
	"assets/static/layui/css/layui.css": &_bintree_t{assets_static_layui_css_layui_css, map[string]*_bintree_t{
	}},
	"assets/static/layui/css/layui.mobile.css": &_bintree_t{assets_static_layui_css_layui_mobile_css, map[string]*_bintree_t{
	}},
	"assets/static/layui/css/modules/code.css": &_bintree_t{assets_static_layui_css_modules_code_css, map[string]*_bintree_t{
	}},
	"assets/static/layui/css/modules/laydate/default/laydate.css": &_bintree_t{assets_static_layui_css_modules_laydate_default_laydate_css, map[string]*_bintree_t{
	}},
	"assets/static/layui/css/modules/layer/default/icon-ext.png": &_bintree_t{assets_static_layui_css_modules_layer_default_icon_ext_png, map[string]*_bintree_t{
	}},
	"assets/static/layui/css/modules/layer/default/icon.png": &_bintree_t{assets_static_layui_css_modules_layer_default_icon_png, map[string]*_bintree_t{
	}},
	"assets/static/layui/css/modules/layer/default/layer.css": &_bintree_t{assets_static_layui_css_modules_layer_default_layer_css, map[string]*_bintree_t{
	}},
	"assets/static/layui/css/modules/layer/default/loading-0.gif": &_bintree_t{assets_static_layui_css_modules_layer_default_loading_0_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/css/modules/layer/default/loading-1.gif": &_bintree_t{assets_static_layui_css_modules_layer_default_loading_1_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/css/modules/layer/default/loading-2.gif": &_bintree_t{assets_static_layui_css_modules_layer_default_loading_2_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/font/iconfont.eot": &_bintree_t{assets_static_layui_font_iconfont_eot, map[string]*_bintree_t{
	}},
	"assets/static/layui/font/iconfont.svg": &_bintree_t{assets_static_layui_font_iconfont_svg, map[string]*_bintree_t{
	}},
	"assets/static/layui/font/iconfont.ttf": &_bintree_t{assets_static_layui_font_iconfont_ttf, map[string]*_bintree_t{
	}},
	"assets/static/layui/font/iconfont.woff": &_bintree_t{assets_static_layui_font_iconfont_woff, map[string]*_bintree_t{
	}},
	"assets/static/layui/font/iconfont.woff2": &_bintree_t{assets_static_layui_font_iconfont_woff2, map[string]*_bintree_t{
	}},
	"assets/static/layui/formSelects-v4.js": &_bintree_t{assets_static_layui_formselects_v4_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/0.gif": &_bintree_t{assets_static_layui_images_face_0_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/1.gif": &_bintree_t{assets_static_layui_images_face_1_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/10.gif": &_bintree_t{assets_static_layui_images_face_10_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/11.gif": &_bintree_t{assets_static_layui_images_face_11_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/12.gif": &_bintree_t{assets_static_layui_images_face_12_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/13.gif": &_bintree_t{assets_static_layui_images_face_13_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/14.gif": &_bintree_t{assets_static_layui_images_face_14_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/15.gif": &_bintree_t{assets_static_layui_images_face_15_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/16.gif": &_bintree_t{assets_static_layui_images_face_16_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/17.gif": &_bintree_t{assets_static_layui_images_face_17_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/18.gif": &_bintree_t{assets_static_layui_images_face_18_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/19.gif": &_bintree_t{assets_static_layui_images_face_19_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/2.gif": &_bintree_t{assets_static_layui_images_face_2_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/20.gif": &_bintree_t{assets_static_layui_images_face_20_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/21.gif": &_bintree_t{assets_static_layui_images_face_21_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/22.gif": &_bintree_t{assets_static_layui_images_face_22_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/23.gif": &_bintree_t{assets_static_layui_images_face_23_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/24.gif": &_bintree_t{assets_static_layui_images_face_24_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/25.gif": &_bintree_t{assets_static_layui_images_face_25_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/26.gif": &_bintree_t{assets_static_layui_images_face_26_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/27.gif": &_bintree_t{assets_static_layui_images_face_27_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/28.gif": &_bintree_t{assets_static_layui_images_face_28_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/29.gif": &_bintree_t{assets_static_layui_images_face_29_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/3.gif": &_bintree_t{assets_static_layui_images_face_3_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/30.gif": &_bintree_t{assets_static_layui_images_face_30_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/31.gif": &_bintree_t{assets_static_layui_images_face_31_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/32.gif": &_bintree_t{assets_static_layui_images_face_32_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/33.gif": &_bintree_t{assets_static_layui_images_face_33_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/34.gif": &_bintree_t{assets_static_layui_images_face_34_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/35.gif": &_bintree_t{assets_static_layui_images_face_35_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/36.gif": &_bintree_t{assets_static_layui_images_face_36_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/37.gif": &_bintree_t{assets_static_layui_images_face_37_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/38.gif": &_bintree_t{assets_static_layui_images_face_38_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/39.gif": &_bintree_t{assets_static_layui_images_face_39_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/4.gif": &_bintree_t{assets_static_layui_images_face_4_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/40.gif": &_bintree_t{assets_static_layui_images_face_40_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/41.gif": &_bintree_t{assets_static_layui_images_face_41_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/42.gif": &_bintree_t{assets_static_layui_images_face_42_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/43.gif": &_bintree_t{assets_static_layui_images_face_43_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/44.gif": &_bintree_t{assets_static_layui_images_face_44_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/45.gif": &_bintree_t{assets_static_layui_images_face_45_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/46.gif": &_bintree_t{assets_static_layui_images_face_46_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/47.gif": &_bintree_t{assets_static_layui_images_face_47_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/48.gif": &_bintree_t{assets_static_layui_images_face_48_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/49.gif": &_bintree_t{assets_static_layui_images_face_49_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/5.gif": &_bintree_t{assets_static_layui_images_face_5_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/50.gif": &_bintree_t{assets_static_layui_images_face_50_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/51.gif": &_bintree_t{assets_static_layui_images_face_51_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/52.gif": &_bintree_t{assets_static_layui_images_face_52_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/53.gif": &_bintree_t{assets_static_layui_images_face_53_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/54.gif": &_bintree_t{assets_static_layui_images_face_54_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/55.gif": &_bintree_t{assets_static_layui_images_face_55_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/56.gif": &_bintree_t{assets_static_layui_images_face_56_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/57.gif": &_bintree_t{assets_static_layui_images_face_57_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/58.gif": &_bintree_t{assets_static_layui_images_face_58_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/59.gif": &_bintree_t{assets_static_layui_images_face_59_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/6.gif": &_bintree_t{assets_static_layui_images_face_6_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/60.gif": &_bintree_t{assets_static_layui_images_face_60_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/61.gif": &_bintree_t{assets_static_layui_images_face_61_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/62.gif": &_bintree_t{assets_static_layui_images_face_62_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/63.gif": &_bintree_t{assets_static_layui_images_face_63_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/64.gif": &_bintree_t{assets_static_layui_images_face_64_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/65.gif": &_bintree_t{assets_static_layui_images_face_65_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/66.gif": &_bintree_t{assets_static_layui_images_face_66_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/67.gif": &_bintree_t{assets_static_layui_images_face_67_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/68.gif": &_bintree_t{assets_static_layui_images_face_68_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/69.gif": &_bintree_t{assets_static_layui_images_face_69_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/7.gif": &_bintree_t{assets_static_layui_images_face_7_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/70.gif": &_bintree_t{assets_static_layui_images_face_70_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/71.gif": &_bintree_t{assets_static_layui_images_face_71_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/8.gif": &_bintree_t{assets_static_layui_images_face_8_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/images/face/9.gif": &_bintree_t{assets_static_layui_images_face_9_gif, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/carousel.js": &_bintree_t{assets_static_layui_lay_modules_carousel_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/code.js": &_bintree_t{assets_static_layui_lay_modules_code_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/colorpicker.js": &_bintree_t{assets_static_layui_lay_modules_colorpicker_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/element.js": &_bintree_t{assets_static_layui_lay_modules_element_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/flow.js": &_bintree_t{assets_static_layui_lay_modules_flow_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/form.js": &_bintree_t{assets_static_layui_lay_modules_form_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/jquery.js": &_bintree_t{assets_static_layui_lay_modules_jquery_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/laydate.js": &_bintree_t{assets_static_layui_lay_modules_laydate_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/layedit.js": &_bintree_t{assets_static_layui_lay_modules_layedit_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/layer.js": &_bintree_t{assets_static_layui_lay_modules_layer_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/laypage.js": &_bintree_t{assets_static_layui_lay_modules_laypage_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/laytpl.js": &_bintree_t{assets_static_layui_lay_modules_laytpl_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/mobile.js": &_bintree_t{assets_static_layui_lay_modules_mobile_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/rate.js": &_bintree_t{assets_static_layui_lay_modules_rate_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/slider.js": &_bintree_t{assets_static_layui_lay_modules_slider_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/table.js": &_bintree_t{assets_static_layui_lay_modules_table_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/transfer.js": &_bintree_t{assets_static_layui_lay_modules_transfer_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/tree.js": &_bintree_t{assets_static_layui_lay_modules_tree_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/upload.js": &_bintree_t{assets_static_layui_lay_modules_upload_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/lay/modules/util.js": &_bintree_t{assets_static_layui_lay_modules_util_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/layui.all.js": &_bintree_t{assets_static_layui_layui_all_js, map[string]*_bintree_t{
	}},
	"assets/static/layui/layui.js": &_bintree_t{assets_static_layui_layui_js, map[string]*_bintree_t{
	}},
	"assets/static/upload/imgs/20190905/1567651858020_647.png": &_bintree_t{assets_static_upload_imgs_20190905_1567651858020_647_png, map[string]*_bintree_t{
	}},
	"assets/upload/imgs/20190905/1567651858020_647.png": &_bintree_t{assets_upload_imgs_20190905_1567651858020_647_png, map[string]*_bintree_t{
	}},
}}
