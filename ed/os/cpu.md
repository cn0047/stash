CPU
-

CPU Instruction Types:
* Data handling and memory operations
* Arithmetic and logic operations
* Control flow operations
* Coprocessor instructions

Almost all processes alternate between two states in a continuing cycle:
* A CPU burst - performing calculations.
* An I/O burst - waiting for data transfer in or out of the system.

CPU bursts vary from process to process, and from program to program.

Whenever the CPU becomes idle, it is the job of the CPU Scheduler (short-term scheduler)
to select another process from the ready queue (FIFO) to run next.
A scheduling system allows one process to use the CPU while another is waiting for I/O.

OS thread is just a sequence of instructions that can be executed independently by a processor.
OS threads are lighter than the process so you can spawn a lot of them.
Linux doesn’t distinguish between threads and processes and both are called tasks.

CPU utilization - % of CPU busyness.
Throughput - Number of processes completed per unit time.
Turnaround time - Time required for a particular process to complete.
Waiting time - How much time processes spend in the ready queue waiting their turn to get on the CPU.
Load average - The average number of processes sitting in the ready queue waiting their turn to get into the CPU.

CPU specific metrics:
* utilization  - the % how CPU is loaded (sum of work handled by CPU).
* idle         - the % of time that CPU was idle.
* i/o wait     - the % of time that CPU was idle during disk I/O request.
* interrupts   - the % of time that CPU spent to service hardware interrupts.
* system level - the % of CPU utilization that occurred while executing at the system level.

Thread Scheduling:
* Contention Scope.
* Pthread Scheduling.

#### Scheduling Algorithms:

* First-Come First-Serve Scheduling, FCFS
  (just a FIFO queue).
* Shortest-Job-First Scheduling, SJF
  (to pick the quickest fastest little job that needs to be done).
* Priority Scheduling
  (is a more general case of SJF, in which each job is assigned a priority).
* Round Robin Scheduling.
* Multilevel Queue Scheduling.
* Multilevel Feedback-Queue Scheduling.

#### Multiple-Processor Scheduling:

* Asymmetric multiprocessing.
* Symmetric multiprocessing, SMP.

#### Arch

386 (aka i386, 80386) - 32-bit microprocessor introduced in 1985.
x64 (aka x86_64, amd64, intel64) - 64-bit version of the x86 instruction set in 1999.

* x86: Intel, AMD; CISC - Complex Instruction Set Computer;
* ARM: ; RISC - Reduced Instruction Set Computing;
