#!/usr/local/bin/ruby -w
puts "Hello, world!"

def say_goodnight(name)
    result = "Good night, " + name
    result = "Good night, #{name}"
    return result
end
# Time for bed...
puts say_goodnight("John-Boy")
puts say_goodnight("Mary-Ellen")
