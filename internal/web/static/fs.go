package static

import "embed"

//go:embed *.js *.css
var Theme embed.FS
