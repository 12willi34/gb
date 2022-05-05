package gb

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
)

var reader = bufio.NewReader(os.Stdin)
var global_s int = 0

type Debugger struct {
  Mu *memoryunit
  Cpu *cpu
  Timer *timer
  Interrupter *interrupter
  Gpu *Gpu
  first_rom_part []byte
  global_i int
  do_dump bool
}

func NewDebugger(boot [256]byte, game string, dump bool) Debugger {
  mu, first_rom_part := NewMemoryUnit(boot, game)
  mu_pointer := &mu
  cpu := NewCPU(mu_pointer)
  cpu_pointer := &cpu
  mu_pointer.Processor = cpu_pointer
  interrupter := NewInterrupter(mu_pointer, cpu_pointer)
  interrupter_pointer := &interrupter
  timer := NewTimer(mu_pointer, interrupter_pointer)
  gpu := NewGpu(mu_pointer, interrupter_pointer)
  return Debugger {
    Mu: mu_pointer,
    Cpu: cpu_pointer,
    Timer: &timer,
    Interrupter: interrupter_pointer,
    Gpu: &gpu,
    first_rom_part: first_rom_part,
    global_i: 0,
    do_dump: dump,
  }
}

func (this Debugger) Init() {
  fmt.Println("starting debugger")
  this.debug_boot_loop(this.do_dump)
  println("boot done")
  this.global_i = 2323814 //wo wird i hier auf 0 gesetzt??
  this.debug_loop(this.do_dump)
}

func (this Debugger) debug_boot_loop(dump bool) {
  for true {
    this.global_i++
    if(!dump) {
      this.showStatus(true)
    }
    this.Interrupter.handle()
    steps := this.Cpu.Step()
    if(steps == -1) { return }
    this.Gpu.Step(steps)
    this.Timer.Timing(steps)
    if(this.Cpu.pc.value >= 0x100) {
      for i := 0; i < 0x100; i++ {
        this.Mu.addr[i] = this.first_rom_part[i]
      }
      return
    }
  }
}

func (this Debugger) debug_loop(dump bool) {
  for true {
    this.global_i++
    if(dump) {
      this.dump()
    } else {
      this.showStatus(false)
    }
    this.Interrupter.handle()
    steps := this.Cpu.Step()
    if(steps == -1) { return }
    this.Gpu.Step(steps)
    this.Timer.Timing(steps)
  }
}

func (this Debugger) dump() {
  op := this.Mu.Read_8(this.Cpu.pc.value)
  var dump string
  if(op == 0xcb) {
    dump = fmt.Sprintf("op: cb %02x\n", this.Mu.Read_8(this.Cpu.pc.value + 1))
  } else {
    dump = fmt.Sprintf("op: %02x\n", this.Mu.Read_8(this.Cpu.pc.value))
  }
  dump += fmt.Sprintf("af: %04x\n", this.Cpu.af.value)
  dump += fmt.Sprintf("bc: %04x\n", this.Cpu.bc.value)
  dump += fmt.Sprintf("de: %04x\n", this.Cpu.de.value)
  dump += fmt.Sprintf("hl: %04x\n", this.Cpu.hl.value)
  dump += fmt.Sprintf("sp: %04x\n", this.Cpu.sp.value)
  dump += fmt.Sprintf("pc: %04x\n\n", this.Cpu.pc.value)
  f, err := os.OpenFile("./gb.dump.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
  if(err != nil) {
    fmt.Println("error opening dumpfile")
    panic(err)
  }
  defer f.Close()
  _, err = f.WriteString(dump)
  if(err != nil) {
    fmt.Println("error writing to dumpfile")
    panic(err)
  }
}

func (this Debugger) showStatus(boot bool) {
  if global_s > 0 {
    global_s--
    return
  }

  fmt.Println(">>>")
  fmt.Printf("op: %02x\n", this.Mu.Read_8(this.Cpu.pc.value))
  fmt.Printf("next op: %02x\n", this.Mu.Read_8(this.Cpu.pc.value + 1))
  fmt.Println("i:", this.global_i)
  fmt.Println()

  fmt.Printf("requested interrupts: %8b\n", this.Mu.addr[0xff0f])
  fmt.Printf("enabled interrupts:   %8b\n", this.Mu.addr[0xffff])
  fmt.Println("IE:", this.Cpu.Interrupt)
  fmt.Println("halted:", this.Cpu.Halt)
  fmt.Println()

  fmt.Printf("stack values\n")
  fmt.Printf("%04x: %02x\n", this.Cpu.sp.value, this.Mu.addr[this.Cpu.sp.value])
  fmt.Printf("%04x: %02x\n\n", (this.Cpu.sp.value + 1), this.Mu.addr[this.Cpu.sp.value + 1])

  fmt.Printf("gpu clock: %d\n", this.Gpu.PubClock)
  fmt.Printf("line: %d\n", this.Mu.addr[0xff44])
  fmt.Printf("gpu mode: %d\n", this.Gpu.PubMode)
  fmt.Println()

  fmt.Printf("af: %04x\n", this.Cpu.af.value)
  fmt.Printf("bc: %04x\n", this.Cpu.bc.value)
  fmt.Printf("de: %04x\n", this.Cpu.de.value)
  fmt.Printf("hl: %04x\n", this.Cpu.hl.value)
  fmt.Printf("sp: %04x\n", this.Cpu.sp.value)
  fmt.Printf("pc: %04x\n", this.Cpu.pc.value)
  fmt.Println()

  x, _ := reader.ReadString('\n')
  num, err := strconv.ParseInt(x[:len(x) - 1], 10, 64)
  if err == nil {
    global_s = int(num)
  }
}
