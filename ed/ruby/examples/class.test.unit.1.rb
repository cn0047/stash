require "test/unit"

require_relative "class.1.rb"

class TestSongList < Test::Unit::TestCase
    def test_delete
        list = SongList.new
        s1 = Song.new("title1", "artist1", 1)
        s2 = Song.new("title2", "artist2", 2)
        s3 = Song.new("title3", "artist3", 3)
        s4 = Song.new("title4", "artist4", 4)
        list.append(s1).append(s2).append(s3).append(s4)
        assert_equal(s1, list[0])
        assert_equal(s3, list[2])
        assert_nil(list[9])
        #
        assert_equal(list.with_title("title2").name, "title2")
        assert_equal(list.with_title_2("title2").name, "title2")
        #
        assert_equal(s1, list.delete_first)
        assert_equal(s2, list.delete_first)
        assert_equal(s4, list.delete_last)
        assert_equal(s3, list.delete_last)
        assert_nil(list.delete_last)
    end
end
