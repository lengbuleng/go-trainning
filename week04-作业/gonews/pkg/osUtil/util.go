package osUtil

import "os"

func GetStaticPath() string {

	dir, _ := os.Getwd()
	static := dir + "/static"

	return static

}

func MakeDirAll(dir string) {
	os.MkdirAll(dir, 0755)
}
