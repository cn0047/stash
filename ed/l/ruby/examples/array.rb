a = [ 1, "cat", 3.14 ]

print "===\n"
animals = %w( ant bee cat dog elk )
animals.each {|animal| puts animal }

print "===> collect\n"
["H", "A", "L"].collect {|x| x.succ } # → ["I", "B", "M"]

print "===> inject\n"
puts [1,3,5,7].inject(100) {|sum, element| sum+element}

print "===\n"
a = [ 1, 3, 5, 7, 9 ]
puts a[-1] # → 9
puts a[-2] # → 7
puts a[-99] # → nil

print "===\n"
puts a[1, 3].inspect # → [3, 5, 7]
puts a[3,1].inspect # → [7]
puts a[-3, 2].inspect # → [5, 7]

print "===> ranges\n"
puts a[1..3].inspect # → [3, 5, 7]
puts a[1...3].inspect # → [3, 5]
puts a[3..3].inspect # →[7]
puts a[-3..-1].inspect # → [5, 7, 9]

print "===> sequences\n"
(1..10).to_a # → [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
('bar'..'bat').to_a # → ["bar", "bas", "bat"]
