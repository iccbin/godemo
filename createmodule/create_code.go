package createmodule

import "os"

func CreatePath(moduleName string, rootPath string) error {
	return os.Mkdir(rootPath + moduleName, 0777)
}