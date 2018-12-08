import java.io.File
import kotlin.system.exitProcess

fun main(args: Array<String>) {
    val input = File("day7.txt").inputStream().bufferedReader().use { it.readText() }
    partA(input)
    partB(input)
}


private fun partA(input: String) {
    val steps = parseInput(input)
    val orderToRemove = mutableListOf<Step>()

    while (steps.isNotEmpty()) {
        orderToRemove.add(doOneStep(steps))
    }

    println("Part A: ${orderToRemove.joinToString("") { it.id }}")
}


fun partB(input: String) {
    val workers = listOf(Worker(0), Worker(1), Worker(2), Worker(3), Worker(4))
    //val workers = listOf(Worker(0), Worker(1))
    val steps = parseInput(input)
    val orderToRemove = mutableListOf<Step>()
    var seconds = 0
    while (steps.isNotEmpty()) {
        val steptime = 60
        assignWorkToWorkers(steps, workers, steptime)
        orderToRemove.addAll(performWork(workers, steps))
        seconds++
    }

    println("Part B: ${seconds} with result ${orderToRemove.joinToString("") { it.id }}")
}

fun performWork(workers: List<Worker>, steps: MutableSet<Step>): MutableList<Step> {
    val workDone = mutableListOf<Step>()
    workers.forEach {
        val itemWorkingOn = it.workingOn
        if(itemWorkingOn != null){
            it.timeleft--
            if( it.timeleft == 0){
                workDone.add(itemWorkingOn)
                steps.remove(itemWorkingOn)
                it.workingOn = null
            }
        }
    }
    return workDone
}

fun assignWorkToWorkers(
    steps: MutableSet<Step>,
    workers: List<Worker>,
    steptime: Int
){
    val freeWorkers = workers.filter { it.workingOn == null }
    val rootNodes = findRootNodes(steps).toMutableList()
    val nodesBeingWorkedOn = workers.map { it.workingOn }
    rootNodes.removeAll(nodesBeingWorkedOn)
    freeWorkers.forEach {
        if(rootNodes.isNotEmpty()){
            val rootNode = rootNodes.first()
            if(workers.any {  it.workingOn == rootNode }.not()){
                it.workingOn = rootNode
                it.timeleft = steptime + rootNode.id.codePointAt(0) - 64
                rootNodes.remove(rootNode)
            }
        }
    }
}

fun doOneStep(steps: MutableSet<Step>): Step {
    val rootNode = findRootNodes(steps).first()
    steps.remove(rootNode)
    return rootNode
}

fun findRootNodes(steps: MutableSet<Step>): List<Step> {
    val notRootNodes = steps.flatMap { it.nextSteps }.map { it.id }.distinct()
    val rootnodes = steps.filter { notRootNodes.contains(it.id).not() }.distinct()
    return rootnodes.sortedBy { it.id }
}

private fun parseInput(input: String): MutableSet<Step> {
    val regex = Regex("Step (.?) must be finished before step (.?) can begin")
    val steps = mutableSetOf<Step>()
    regex.findAll(input).forEach {
        var beforeStep = steps.find { step -> it.groupValues[1] == step.id }
        var afterStep = steps.find { step -> it.groupValues[2] == step.id }

        if (beforeStep == null) {
            beforeStep = Step(it.groupValues[1])
            steps.add(beforeStep)
        }
        if (afterStep == null) {
            afterStep = Step(it.groupValues[2])
            steps.add(afterStep)
        }

        beforeStep.nextSteps.add(afterStep)
    }

    steps.forEach {
            s ->
                val sortedBy = s.nextSteps.sortedBy { it.id }
                s.nextSteps = sortedBy.toMutableList()
    }
    val sortedList = steps.sortedBy { it.id }
    return sortedList.toMutableSet()
}

data class Worker(val id: Int, var workingOn : Step? = null, var timeleft : Int = 0)


data class Step(val id : String, var nextSteps : MutableList<Step> = mutableListOf<Step>()){
    override fun toString(): String {
        val nextString = nextSteps.joinToString(", ") { it.id }
        return "Step(id='$id', nextSteps= [$nextString])"
    }

    override fun equals(other: Any?): Boolean {
        if (other is Step){
            return id == other.id
        }
        return false
    }

    override fun hashCode(): Int {
        return super.hashCode()
    }
}