def taBortBokstavsskillnad(line):
    for index, item in enumerate(line):
        if index < (len(line) - 1):
            if item != line[index+1] and item.lower() == line[index + 1].lower():
                return taBortBokstavsskillnad(line[:index] + line[index+2:])
    return line

if __name__ == '__main__':
    f = open("day5.txt", "r")
    for l in f.readlines():
        print(len(taBortBokstavsskillnad(l.rstrip())))
