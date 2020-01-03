require 'test/unit'

require_relative 'class.Roman.rb'

class TestRoman < Test::Unit::TestCase
    def test_simple
        assert_equal("i", Roman.new(1).to_s)
        assert_equal("ix", Roman.new(9).to_s)
        #
        assert_equal("i", Roman.new(1).to_s) 
        assert_equal("ii", Roman.new(2).to_s) 
        assert_equal("iii", Roman.new(3).to_s) 
        assert_equal("iv", Roman.new(4).to_s) 
        assert_equal("ix", Roman.new(9).to_s)
    end
end
