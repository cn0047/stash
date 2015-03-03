vim
-

version 7.4

````
:version
:redo or CTRL+R
:e!                 - reload file
:%s/text/newText/

ctrl+p - show competition

:split         - 2 windows
ctrl+shift+w   - switch window
:split file    - split with new file
:11 split file - split & show 11 lines of file

%   - bracket
~   - change case
J   - join lines
.   - do last edit action again
{   - paragraph begin
}   - paragraph end
]]  - next function
[[  - prev function
2<< - move twice to left
>>  - to right

v  - visual mode
d  - delete selected
y  - cope
j  - join
gJ - join without spaces

K  - LOOK TO SCREEN
````

####.vimrc
````
:set nonumber
:syntax on
:set paste
:set autoindent
:set autowrite
:set shiftwidth=4
:set hlsearch
:set incsearch
:set filetype=php
:set tabstop=4
:set undolevels=5000
:set formatoptions=cr
:set ignorecase
:set list listchars=tab:→\ ,trail:·

:set showmatch
:set nobackup
:set nowritebackup
:set noswapfile
:set smarttab
:set et
:set ai
:set cin
:set ruler

map <S-left>  :tabp<cr>
map <S-right> :tabn<cr>
map <S-down>  :tabc<cr>
map <S-up>    :qa!<cr>

set pastetoggle=<F4>

````
