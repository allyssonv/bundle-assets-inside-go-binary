// Code generated by rice embed-go; DO NOT EDIT.
package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "app.js",
		FileModTime: time.Unix(1568553002, 0),

		Content: string("document.getElementById(\"text\").innerHTML = \"Hello World </br> <p>Created by Allysson Vieira</p>\""),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "estilo.css",
		FileModTime: time.Unix(1568553431, 0),

		Content: string(".container {\n\tbackground-color: #348A9D;\n\twidth: 50%;\n\theight: 50%;\n\n}\n\n.text {\n\tfont-family: Monaco;\n\tcolor: white;\n\tfont-size: 3em;\n\ttext-align: center;\n}"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1568553228, 0),

		Content: string("<!DOCTYPE html>\n<html>\n<head>\n\t<title>Bundle assets inside go binary</title>\n\t<link rel=\"stylesheet\" type=\"text/css\" href=\"estilo.css\">\n</head>\n<body>\n\n\t<div class=\"container\">\n\t\t<p id=\"text\" class=\"text\"></p>\n\t</div>\n\n\n<script src=\"app.js\"></script>\n</body>\n</html>"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1568553196, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "app.js"
			file3, // "estilo.css"
			file4, // "index.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`website`, &embedded.EmbeddedBox{
		Name: `website`,
		Time: time.Unix(1568553196, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"app.js":     file2,
			"estilo.css": file3,
			"index.html": file4,
		},
	})
}
