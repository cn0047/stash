java -jar /usr/local/bin/selenium-server-standalone-2.33.0.jar -interactive

cmd=getNewBrowserSession&1=*firefox&2=http://www.google.com.ua
cmd=open&1=http://www.google.com.ua&sessionId=7fd2aedf629644c393c8d8e87d8cb8e7
cmd=type&1=q&2=selenium&sessionId=7fd2aedf629644c393c8d8e87d8cb8e7
cmd=click&1=btnG&sessionId=7fd2aedf629644c393c8d8e87d8cb8e7
cmd=isElementPresent&1=//a[contains(text(),"Selenium")]&sessionId=7fd2aedf629644c393c8d8e87d8cb8e7
cmd=testComplete&sessionId=7fd2aedf629644c393c8d8e87d8cb8e7
