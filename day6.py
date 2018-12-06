def getRectForField(coords):
    maxX = max(coords, key=lambda x: x[0])
    maxY = max(coords, key=lambda x: x[1])
    return (maxX[0], maxY[1])
def parseCoords(input):
    coords = [list(map(int, line.rstrip().split(", "))) for line in input]
    coords.sort()
    return coords

if __name__ == '__main__':
    f = open("day6.txt", "r")    
    print(getRectForField(parseCoords(f.readlines())))