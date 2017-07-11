ruby
-
2.0.0p648

## Data Types:

* Number (integers and floating-point)
* String
* Symbol
* True & False (Boolean)
* Array
* Hash
* Nil

+ Range (Range, Sequence)

Symbols are used to represent other objects. A Ruby symbol cannot be changed at runtime.

There are two types of conversion: implicit and explicit.
Implicit type conversion, also known as coercion, is an automatic type conversion by the compiler.
Ruby has only explicit conversion.

````rb
i1, i2 = 1, 1 # parallel assignment (i1 = 1 and i2 = 1)

File.dirname(__FILE__)

# call function
@action.call(self)

#number
Integer(v1) # cast to int

# string
printf("Number: %5.2f,\nString: %s\n", 1.23, "hello")

'escape using "\\"' → escape using "\"
'That\'s right' → That's right

"Seconds/day: #{24*60*60}" # → Seconds/day: 86400
"#{'Ho! '*3}Merry Christmas!" # → Ho! Ho! Ho! Merry Christmas!
"This is line #$." # → This is line 3

string = <<END_OF_STRING
  The body of the string
END_OF_STRING

file, length, name, title = line.chomp.split(/\s*\|\s*/)

# range
1..10
'a'..'z'

digits = 0..9
digits.include?(5) # → true
digits.min # → 0
digits.max # → 9
digits.reject {|i| i < 5 } # → [5, 6, 7, 8, 9]
digits.each {|digit| dial(digit) } # → 0..9

(1..10)    === 5 # → true 
(1..10)    === 15 # → false 
(1..10)    === 3.14159 # → true 
('a'..'j') === 'c' # → true 
('a'..'j') === 'z' # → false

# regex
if line =~ /Perl|Python/
    puts "Scripting language mentioned: #{line}"
end

# unit test
ruby test_roman.rb --name test_range

# command-line arguments
ruby -w test.rb "Hello World" a1 1.6180
ARGV.each {|arg| p arg }
````

````rb
if count > 10
  #
elsif tries == 3
  #
else
  #
end

begin
  raise "Missing name" if name.nil?
  raise InterfaceException
rescue SystemCallError
  #
else
  #
ensure
  #
end

while weight < 100 and num_pallets <= 30 pallet = next_pallet()
    #
end

6.times do
  #
end

3.times { print "X " }
1.upto(5) {|i| print i, " " }
99.downto(95) {|i| print i, " " }
50.step(80, 5) {|i| print i, " " }
````
````
Exception
    fatal (used internally by Ruby) NoMemoryError
    ScriptError
        LoadError
        NotImplementedError
        SyntaxError
    SignalException
        Interrupt
    StandardError
        ArgumentError
        IOError
            EOFError
        IndexError
        LocalJumpError
        NameError
            NoMethodError
        RangeError
            FloatDomainError
        RegexpError
        RuntimeError
        SecurityError
        SystemCallError
            system-dependent exceptions (Errno::xxx)
        ThreadError
        TypeError
        ZeroDivisionError
    SystemExit
    SystemStackError
````

## OOP

Protected method, may be called by any instance of the defining class or its subclasses.
If a method is private, it may be called only within the context of the calling object.

## Hash

Compared with arrays, hashes have one significant advantage: they can use any object as an index.
However, they also have a significant disadvantage: their elements are not ordered,
so you cannot easily use a hash as a stack or a queue.

A second, bigger problem is that a hash does not support multiple keys with the same value.

## Yield

Whenever a yield is executed, it invokes the code in the block.

````rb
class Array
      def find
        for i in 0...size
          value = self[i]
          return value if yield(value)
        end
        return nil
      end
end
[1,3,5,7,9].find{|v|v*v>30} → 7
````

## IRB (Interactive Ruby)

````
irb
  -r load-module

# subsessions
irb "wombat"

exit, quit, irb_exit, irb_quit

conf, context, irb_context
conf conf.auto_indent_mode=true

cb, irb_change_binding

# lists irb subsessions
jobs, irb_jobs

kill n, irb_kill n

eval "var = 1"
````

## Debug

````
ruby -r debug ed/ruby/examples/class.1.rb
````

From book "Programming ruby, 2nd ed" Finished chapters: 1-5, 10, 12, 15.
