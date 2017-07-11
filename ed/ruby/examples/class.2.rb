class File
    def File.open_and_process(*args)
        f = File.open(*args)
        yield f
        f.close()
    end
end

File.open_and_process("hw.rb", "r") do |file|
    while line = file.gets
        puts line
    end
end
