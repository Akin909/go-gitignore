# go-gitignore
A Small Go program which creates a gitignore in the current directory

Go-gitignore takes two arguments at the moment `--rm` which is a boolean which indicates whether or not the existing gitignore should
be replaced. The second argument is `--ft={node|go|elm}` the options are node, go or elm, it prepopulates the gitignore with one of 
git ignores for these file types. Finally it takes an unlimited amount of space delimited arguments which are files
which will be added to the gitignore e.g. `go-gitignore node_modules build .DS_Store`
