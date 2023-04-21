package blogposts

// Go 1.16 introduced an abstraction for file systems - io/fs package
// + there is testing/fstest

import (
	"io/fs"
	"testing/fstest"
)

type Post struct {
}

func NewPostsFromFS(fileSystem fstest.MapFS) []Post {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts
}