class Roman
    MAX_ROMAN = 4999
    
    def initialize(value)
        if value <= 0 || value > MAX_ROMAN
            fail "Roman values must be > 0 and <= #{MAX_ROMAN}"
        end
        @value = value
    end

    FACTORS = [["m", 1000], ["cm", 900], ["d", 500], ["cd", 400],
               ["c", 100], ["xc", 90], ["l", 50], ["xl", 40],
               ["x", 10], ["ix", 9], ["v", 5], ["iv",   4],
               ["i",    1]
            ]     
    def to_s
        value = @value
        roman = ""
        for code, factor in FACTORS
           count, value = value.divmod(factor)
           roman << code unless count.zero?
        end
        roman
    end
end
