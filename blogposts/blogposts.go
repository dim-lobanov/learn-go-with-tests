package blogposts

// Go 1.16 introduced an abstraction for f systems - io/fs package
// + there is testing/fstest

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".") // TODO error
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPostFromFile(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPostFromFile(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

// we can change to interface Reader, because we only use io.ReadAll and don't have to couple to FS
// Scanner provides an interface for reading data such as a file of newline-delimited lines of text.
func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile) // note that NewScanner accepts io.Reader

	// separating what from how by extracting this func
	readLine := func() string {
		scanner.Scan()        // returns next token or false
		return scanner.Text() // returns a new string from the latest token from scanner
		// string is newly allocated and holds it's bytes
	}

	title := readLine()[len(titleSeparator):]
	description := readLine()[len(descriptionSeparator):]
	tagsLine := readLine()[len(tagsSeparator):]
	var tags []string = strings.Split(tagsLine, ", ")
	body := readBody(scanner)

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()        // ignore a body separator line
	buf := bytes.Buffer{} // struct with bytes array, offset int and last read operation for Unread!
	for scanner.Scan() {  // Scan() returns a bool, so we can iterate until EOF (or error)
		fmt.Fprintln(&buf, scanner.Text()) // Buffer{} implements io.Writer
	}
	return strings.TrimSuffix(buf.String(), "\n") // remove last new line symbol
}
