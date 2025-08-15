vim
-
<br>version 7.4

````sh
# open file & goto line 7
vim file +7
````

````sh
ctrl+p  - show competition
shift+v - visual line mode
ctrl+b  - visual block mode

:%!xxd    # transform to hex editor
:%!xxd -r # reverse transform

:q  # quit
:q! #
:w  #
:x  #
:e! # reload file

:version
u               # undo last change
:redo or CTRL+R

:%s/text/newText/ # replace

:split         # 2 windows
ctrl+shift+w   # switch window
:split file    # split with new file
:11 split file # split & show 11 lines of file
:11            # goto line 11

%   # bracket
~   # change case
J   # join lines
.   # do last edit action again
{   # paragraph begin
}   # paragraph end
]]  # next function
[[  # prev function
2<< # move twice to left
>>  # to right

v   # visual mode
d   # delete selected
y   # copy
j   # join
gJ  # join without spaces

gg      # go to first line
shift+g # go to last line

K   # LOOK TO SCREEN

ma  # bookmark a
d'a # delete to bookmark a
````

#### ~/.vimrc

````
:set nowrap
````
````
:syntax on
:set nonumber
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
