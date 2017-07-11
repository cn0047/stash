#!/usr/bin/ruby

# brew install mysql
# sudo gem install mysql2
# sudo gem install ruby-mysql
require "mysql"

con = Mysql.new("mysql-master", "dbu", "dbp", "test")
# rs = con.query("SELECT NOW() as d")
# rs.each_hash { |h| puts h["d"]}
# con.close
