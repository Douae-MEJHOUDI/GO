module package_example/hello

go 1.23.4

replace simplemath => ../simplemath

require (
	github.com/pkg/math v0.0.0-20141027224758-f2ed9e40e245 // indirect
	simplemath v0.0.0-00010101000000-000000000000 // indirect
)
