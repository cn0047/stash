template
-

[pkg](https://golang.org/pkg/text/template/)

````
{{range $prj := .projects}}
  <option value="{{ $prj.ID }}">{{ $prj.ID }}</option>
{{end}}
````
