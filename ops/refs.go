package ops

import (
	"gogit/fs"
)

func SetHead(ref string) {
	// Set the HEAD file to the specified ref
	// This is used to set the HEAD to a branch
	// or to a commit
	fs.WriteBytesToFile(".gogit/HEAD", []byte(ref))
}
