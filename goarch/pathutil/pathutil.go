package pathutil

import "strings"

/* ExtendPath extends the path by given names
Ex: Given path: "./internal" and names are ["core", "logger"]
then the result will be "./internal/core/logger"
*/
func ExtendPath(path string, names ...string) string {
	return path + "/" + strings.Join(names, "/")
}
