JetBrains
-

#### Live templates:

````php
d1 - file_put_contents('/tmp/debug.log', var_export($END$, 1)."\n", FILE_APPEND); /// tail -f /tmp/debug.log
````

````js
l - console.log($END$);
````

#### .editorconfig

````
indent_style = space
indent_size = 4

end_of_line = lf
charset = utf-8
trim_trailing_whitespace = true
insert_final_newline = true

[*.md]
trim_trailing_whitespace = true

[package.json]
indent_size = 2
````

#### PhpStorm

Plugins:
* PhpInspectionsEA
* Symfony2 Plugin
* PHP Annotation
* Lines Sorter

Keys:
* `Shift+Enter` Show intention actions (quick fixes).
* `Ctrl+Alt+T` Surround with.
* `Ctrl+Alt+M` Extract method refactoring.
* `Ctrl+Alt+O` Optimize imports (delete non-used imported classes).

#### GoLand

[File Watchers](https://monosnap.com/file/HHVwIfCi5rwezBvUx3tEftZdU5icCh):

````sh
# go fmt
$GoExecPath$
fmt $FilePath$
````

#### IntelliJ IDEA

Plugins: Lombok, Checkstyle.

````
Settings > Code Style > Java > Imports > General > Class count to use import with '*'
````
