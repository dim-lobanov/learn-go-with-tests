package blogpost_test

import (
	"testing"
	"testing/fstest"
	"learn-go-with-tests/blogposts" 
	// <modulename>/<relative-filepath-from-module-root>
	// directory name has to correspond to the package name for us to import it
)

func TestNewBlogPosts(t *testing.T) {
	// A MapFS is a simple in-memory file system for use in tests, represented as a map from 
	// path names (arguments to Open) to information about the files or directories they represent.
	fs := fstest.MapFS{
		"hello_world.md": {Data: []byte("hi")},
		"hello_world2.md": {Data: []byte("hola")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, want %d", len(posts), len(fs))
	}

}