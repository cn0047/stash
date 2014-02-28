	SET SQL_SAFE_UPDATES=0;
	COLLATE UTF8_GENERAL_CI LIKE

	tee /tmp/out
	cat /tmp/out | mail mail@com.com

	--pager="less -S"
	mysql> pager less -SFX
	disable-pager
	--pager="less -SFX"