package lrand

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-04-17

const (
	n         = 312
	m         = 156
	notSeeded = n + 1

	hiMask uint64 = 0xffffffff80000000
	loMask uint64 = 0x000000007fffffff

	matrixA uint64 = 0xB5026F5AA96619E9

	defaultSeed int64 = 5489
)

type Mersenne struct {
	state []uint64
	index int
}

func New() *Mersenne {
	res := &Mersenne{
		state: make([]uint64, n),
		index: notSeeded,
	}
	return res
}

func (mt *Mersenne) Seed(seed int64) {
	x := mt.state
	x[0] = uint64(seed)
	for i := uint64(1); i < n; i++ {
		x[i] = 6364136223846793005*(x[i-1]^(x[i-1]>>62)) + i
	}
	mt.index = n
}

func (mt *Mersenne) SeedFromSlice(key []uint64) {
	mt.Seed(19650218)

	x := mt.state
	i := uint64(1)
	j := 0
	k := len(key)
	if n > k {
		k = n
	}
	for k > 0 {
		x[i] = (x[i] ^ ((x[i-1] ^ (x[i-1] >> 62)) * 3935559000370003845) +
			key[j] + uint64(j))
		i++
		if i >= n {
			x[0] = x[n-1]
			i = 1
		}
		j++
		if j >= len(key) {
			j = 0
		}
		k--
	}
	for j := uint64(0); j < n-1; j++ {
		x[i] = x[i] ^ ((x[i-1] ^ (x[i-1] >> 62)) * 2862933555777941757) - i
		i++
		if i >= n {
			x[0] = x[n-1]
			i = 1
		}
	}
	x[0] = 1 << 63
}

func (mt *Mersenne) Uint64() uint64 {
	x := mt.state
	if mt.index >= n {
		if mt.index == notSeeded {
			mt.Seed(defaultSeed)
		}
		for i := 0; i < n-m; i++ {
			y := (x[i] & hiMask) | (x[i+1] & loMask)
			x[i] = x[i+m] ^ (y >> 1) ^ ((y & 1) * matrixA)
		}
		for i := n - m; i < n-1; i++ {
			y := (x[i] & hiMask) | (x[i+1] & loMask)
			x[i] = x[i+(m-n)] ^ (y >> 1) ^ ((y & 1) * matrixA)
		}
		y := (x[n-1] & hiMask) | (x[0] & loMask)
		x[n-1] = x[m-1] ^ (y >> 1) ^ ((y & 1) * matrixA)
		mt.index = 0
	}
	y := x[mt.index]
	y ^= (y >> 29) & 0x5555555555555555
	y ^= (y << 17) & 0x71D67FFFEDA60000
	y ^= (y << 37) & 0xFFF7EEE000000000
	y ^= (y >> 43)
	mt.index++
	return y
}

func (mt *Mersenne) Int63() int64 {
	x := mt.state
	if mt.index >= n {
		if mt.index == notSeeded {
			mt.Seed(defaultSeed)
		}
		for i := 0; i < n-m; i++ {
			y := (x[i] & hiMask) | (x[i+1] & loMask)
			x[i] = x[i+m] ^ (y >> 1) ^ ((y & 1) * matrixA)
		}
		for i := n - m; i < n-1; i++ {
			y := (x[i] & hiMask) | (x[i+1] & loMask)
			x[i] = x[i+(m-n)] ^ (y >> 1) ^ ((y & 1) * matrixA)
		}
		y := (x[n-1] & hiMask) | (x[0] & loMask)
		x[n-1] = x[m-1] ^ (y >> 1) ^ ((y & 1) * matrixA)
		mt.index = 0
	}
	y := x[mt.index]
	y ^= (y >> 29) & 0x5555555555555555
	y ^= (y << 17) & 0x71D67FFFEDA60000
	y ^= (y << 37) & 0xFFF7EEE000000000
	y ^= (y >> 43)
	mt.index++
	return int64(y & 0x7fffffffffffffff)
}

func (mt *Mersenne) Read(p []byte) (n int, err error) {
	for n+8 <= len(p) {
		val := mt.Uint64()
		p[n] = byte(val)
		p[n+1] = byte(val >> 8)
		p[n+2] = byte(val >> 16)
		p[n+3] = byte(val >> 24)
		p[n+4] = byte(val >> 32)
		p[n+5] = byte(val >> 40)
		p[n+6] = byte(val >> 48)
		p[n+7] = byte(val >> 56)
		n += 8
	}
	if n < len(p) {
		val := mt.Uint64()
		for n < len(p) {
			p[n] = byte(val)
			val >>= 8
			n++
		}
	}
	return n, nil
}
