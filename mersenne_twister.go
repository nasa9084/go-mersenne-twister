package mt

const (
	n         = 624
	m         = 397
	matrixA   = 0x9908b0df
	upperMask = 0x80000000
	lowerMask = 0x7fffffff
)

type mtRandom struct {
	mt  [n]uint32
	mti uint32
}

// create and initializes random sequence with a seed value
func New() *mtRandom {
	res := &mtRandom{
		mti: n + 1,
	}

	return res
}

func (r *mtRandom) InitGenrand(s uint32) {
	r.mt[0] = s & 0xffffffff
	for r.mti = 1; r.mti < n; r.mti++ {
		r.mt[r.mti] = 1812433253*(r.mt[r.mti-1]^(r.mt[r.mti-1]>>30)) + r.mti
		r.mt[r.mti] &= 0xffffffff
	}
}

// InitByArray initializes random sequence with a slice
func (r *mtRandom) InitByArray(initKey []uint32) {
	var i, j, k uint32
	r.InitGenrand(19650218)
	i = 1
	j = 0
	keyLength := uint32(len(initKey))
	k = keyLength
	if n > keyLength {
		k = n
	}
	for ; k != 0; k-- {
		r.mt[i] = (r.mt[i] ^ ((r.mt[i-1] ^ (r.mt[i-1] >> 30)) * 1664525)) + initKey[j] + j
		r.mt[i] &= 0xffffffff
		i++
		j++
		if i >= n {
			r.mt[0] = r.mt[n-1]
			i = 1
		}
		if j >= keyLength {
			j = 0
		}
	}
	for k = n - 1; k != 0; k-- {
		r.mt[i] = (r.mt[i] ^ ((r.mt[i-1] ^ (r.mt[i-1] >> 30)) * 1566083941)) - i
		r.mt[i] &= 0xffffffff
		i++
		if i >= n {
			r.mt[0] = r.mt[n-1]
			i = 1
		}
	}

	r.mt[0] = 0x80000000
}

// GenrandInt32 generates a random 32bit unsigned int number
func (r *mtRandom) GenrandInt32() uint32 {
	var y uint32
	mag01 := [2]uint32{0x0, matrixA}

	if r.mti >= n {
		var kk int

		if r.mti == n+1 {
			r.InitGenrand(5489)
		}
		for kk = 0; kk < n-m; kk++ {
			y = (r.mt[kk] & upperMask) | (r.mt[kk+1] & lowerMask)
			r.mt[kk] = r.mt[kk+m] ^ (y >> 1) ^ mag01[y&0x1]
		}
		for ; kk < n-1; kk++ {
			y = (r.mt[kk] & upperMask) | (r.mt[kk+1] & lowerMask)
			r.mt[kk] = r.mt[kk+(m-n)] ^ (y >> 1) ^ mag01[y&0x1]
		}
		y = (r.mt[n-1] & upperMask) | (r.mt[0] & lowerMask)
		r.mt[n-1] = r.mt[m-1] ^ (y >> 1) ^ mag01[y&0x1]

		r.mti = 0
	}

	y = r.mt[r.mti]
	r.mti++

	y ^= (y >> 11)
	y ^= (y << 7) & 0x9d2c5680
	y ^= (y << 15) & 0xefc60000
	y ^= (y >> 18)

	return y
}

// GenrandInt31 generates a 31bit unsigned int random number
// note: return type is uint32
func (r *mtRandom) GenrandInt31() uint32 {
	return uint32(r.GenrandInt32() >> 1)
}

// GenrandReal1 generates a 32bit [0, 1] real random number
// note: return type is float64, not float32
func (r *mtRandom) GenrandReal1() float64 {
	return float64(r.GenrandInt32()) * (1.0 / 4294967295.0)
}

// GenrandReal2 generates a 32bit [0, 1) real random number
// note: return type is float64, not float32
func (r *mtRandom) GenrandReal2() float64 {
	return float64(r.GenrandInt32()) * (1.0 / 4294967296.0)
}

// GenrandReal3 generates a 32bit (0, 1) real random  number
// note: return type is float64, not float32
func (r *mtRandom) GenrandReal3() float64 {
	return ((float64(r.GenrandInt32())) + 0.5) * (1.0 / 4294967296.0)
}

// GenrandRes53 generates a [0, 1) random number with 53-bit resolution
func (r *mtRandom) GenrandRes53() float64 {
	a := r.GenrandInt32() >> 5
	b := r.GenrandInt32() >> 6
	return (float64(a)*67108864.0 + float64(b)) * (1.0 / 9007199254740992.0)
}

var mt = New()

func InitGenrand(s uint32) {
	mt.InitGenrand(s)
}

func InitByArray(initKey []uint32) {
	mt.InitByArray(initKey)
}

func GenrandInt32() uint32 {
	return mt.GenrandInt32()
}

func GenrandInt31() uint32 {
	return mt.GenrandInt31()
}

func GenrandReal1() float64 {
	return mt.GenrandReal1()
}

func GenrandReal2() float64 {
	return mt.GenrandReal2()
}

func GenrandReal3() float64 {
	return mt.GenrandReal3()
}

func GenrandRes53() float64 {
	return mt.GenrandRes53()
}
