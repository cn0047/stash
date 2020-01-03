class Song

    @@plays = 0
      
    def initialize(name, artist, duration)
        @name     = name
        @artist   = artist
        @duration = duration
        @plays    = 0
    end
    
    def method1
    end

    def method2
    end

    def name
        @name
    end
    
    def artist
      @artist
    end
    
    def duration
      @duration
    end

    public
    def duration=(new_duration)
      @duration = new_duration
    end
    
    def to_ss
        "Song: #@name - #@artist (#@duration) [#@lyrics]"
    end

    def play
        @plays += 1
        @@plays += 1
        "This song: #@plays plays. Total #@@plays plays."
    end

    private   :method1, :method2

end

class KaraokeSong < Song
    def initialize(name, artist, duration, lyrics)
        super(name, artist, duration)
        @lyrics = lyrics
    end
end

class SongList
    def initialize
        @songs = Array.new
    end
    def [](index)
        @songs[index]
      end
    def append(song)
        @songs.push(song)
        self
    end
    def delete_first
        @songs.shift
    end
    def delete_last
        @songs.pop
    end
    def with_title(title)
        for i in 0...@songs.length
            return @songs[i] if title == @songs[i].name
        end
        return nil
    end
    def with_title_2(title)
        @songs.find {|sng| title == sng.name }
    end
end

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
