def getRectForField(coords):
    maxX = max(coords, key=lambda x: x[0])
    maxY = max(coords, key=lambda x: x[1])
    return (maxX[0], maxY[1])


def parseCoords(input):
    coords = [list(map(int, line.rstrip().split(", "))) for line in input]
    coords.sort()
    return [(x[0], x[1]) for x in coords]

def closestPoint(coords, dx, dy):    
    distanceList = [(c, manhattanDistance(c[0], c[1], dx, dy)) for c in coords]
    minVal = min(distanceList, key=lambda x: x[1])
    distvalues = [x[1] for x in distanceList]
    if distvalues.count(minVal[1]) == 1:
        return minVal
    return None


def manhattanDistance(x, y, dx, dy):
    return abs(dx - x) + abs(dy - y)

def getInfinitePoints(grid, maxX, maxY):
    infinitePoints = set()
    for x in range (maxX):
        if grid[(x, 0)]:
            infinitePoints.add(grid[(x, 0)][0])
        if grid[(x, maxY)]:
            infinitePoints.add(grid[(x, maxY)][0])
    for y in range (maxY):
        if grid[(0, y)]:
            infinitePoints.add(grid[(0, y)][0])
        if grid[(maxX, y)]:
            infinitePoints.add(grid[(maxX, y)][0])

    return infinitePoints

def getAreaForPoint(p, grid):
    area = 0
    for points in grid.values():
        if points:
            for point in points:
                if p == point:
                    area = area + 1
    return area

def doPartA(coords):
    rectMax = getRectForField(coords)
    grid = {}
    for y in range(rectMax[1]+1):
        for x in range (rectMax[0]+1):
            grid[(x, y)] = closestPoint(coords, x, y)
    infinitePoints = getInfinitePoints(grid, rectMax[0], rectMax[1])
    finitePoints = [v for v in coords if v not in infinitePoints]
    print("Part 1:", max([(p, getAreaForPoint(p, grid)) for p in finitePoints], key= lambda x: x[1]))

def getSafe(limit, coords,  x, y):
    distanceList = [(c, manhattanDistance(c[0], c[1], x, y)) for c in coords]
    summan = sum([x[1] for x in distanceList])
    if summan < limit:
        return True
    return False

def doPartB(limit, coords):
    rectMax = getRectForField(coords)
    grid = {}
    for y in range(rectMax[1]+1):
        for x in range (rectMax[0]+1):
            grid[(x, y)] = getSafe(limit, coords, x, y)
    print("Part 2:", len([safe for safe in grid.values() if safe == True]))
if __name__ == '__main__':
    f = open("day6.txt", "r") 
    coords = parseCoords(f.readlines())  
    doPartA(coords)
    doPartB(10000, coords)