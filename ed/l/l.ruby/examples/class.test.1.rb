require_relative "class.1.rb"

song = Song.new("Bicylops", "Fleck", 260)
puts song.inspect
puts song.artist
puts song.duration
song.duration = 275
puts song.duration
puts song.to_s
puts song.to_ss
puts song.play
puts song.play
puts song.play
puts song.play
song2 = KaraokeSong.new("My Way", "Sinatra", 225, "And now, the...") 
puts song2.inspect
puts song2.play
