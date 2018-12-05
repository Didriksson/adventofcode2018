def taBortBokstavsskillnad(line):
    rensat = True
    while rensat:
        rensat = False
        for index, item in enumerate(line):
            if index < (len(line) - 1):
                if item != line[index+1] and item.lower() == line[index + 1].lower():
                    line = line[:index] + line[index+2:]
                    rensat = True
                    break
    return line

if __name__ == '__main__':
    f = open("day5.txt", "r")
    for l in f.readlines():
        print(len(taBortBokstavsskillnad(l.rstrip())))