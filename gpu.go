package gb

import ()

const width = 160
const height = 144

type Gpu struct {
  mu *memoryunit
}

func NewGpu(mu *memoryunit) *Gpu {
  return &(Gpu {
    mu: mu,
  })
}
