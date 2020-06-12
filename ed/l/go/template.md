template
-

[pkg](https://golang.org/pkg/text/template/)

````
{{ (print "\n---\n") -}}

{{range $prj := .projects}}
  <option value="{{ $prj.ID }}">{{ $prj.ID }}</option>
{{end}}

{{if .public -}}
  public
{{else if .x -}}
  unknown
{{else -}}
  private
{{end -}}
````
