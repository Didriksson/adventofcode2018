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

def taBortAllaForekomster(bokstav, line):
    lower = bokstav.lower()
    upper = bokstav.upper()
    return line.replace(lower, "").replace(upper, "")

if __name__ == '__main__':
    f = open("day5.txt", "r")
    line = f.readline().rstrip()
    print("Warning! If you are using this on real input - go for a coffee!")
    print("Part 1:", len(taBortBokstavsskillnad(line)))
    print("Part 2:", min([len(taBortBokstavsskillnad(taBortAllaForekomster(it, line))) for it in "ABCDEFGHIJKLMNOPQRSTUVXYZ"]))   