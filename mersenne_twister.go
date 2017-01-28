package mt

const (
	n          = 624
	m          = 397
	matrixA   = 0x9908b0df
	upperMask = 0x80000000
	lowerMask = 0x7fffffff
)

var mt [n]uint32
var mti uint32 = n + 1

/* InitGenrand initializes random sequence with a seed value */
func InitGenrand(s uint32) {
	mt[0] = s & 0xffffffff
	for mti = 1; mti < n; mti++ {
		mt[mti] = (1812433253*(mt[mti-1]^(mt[mti-1]>>30)) + mti)
		mt[mti] &= 0xffffffff
	}
}

/* InitByArray initializes random sequence with a slice */
func InitByArray(init_key []uint32) {
	var i, j, k uint32
	InitGenrand(19650218)
	i = 1
	j = 0
	var key_length uint32 = uint32(len(init_key))
	k = key_length
	if n > key_length {
		k = n
	}
	for ; k != 0; k-- {
		mt[i] = (mt[i] ^ ((mt[i-1] ^ (mt[i-1] >> 30)) * 1664525)) + init_key[j] + j
		mt[i] &= 0xffffffff
		i++
		j++
		if i >= n {
			mt[0] = mt[n-1]
			i = 1
		}
		if j >= key_length {
			j = 0
		}
	}
	for k = n - 1; k != 0; k-- {
		mt[i] = (mt[i] ^ ((mt[i-1] ^ (mt[i-1] >> 30)) * 1566083941)) - i
		mt[i] &= 0xffffffff
		i++
		if i >= n {
			mt[0] = mt[n-1]
			i = 1
		}
	}

	mt[0] = 0x80000000
}

/* GenrandInt32 generates a random 32bit unsigned int number */
func GenrandInt32() uint32 {
	var y uint32
	var mag01 [2]uint32 = [2]uint32{0x0, matrixA}

	if mti >= n {
		var kk int

		if mti == n+1 {
			InitGenrand(5489)
		}
		for kk = 0; kk < n-m; kk++ {
			y = (mt[kk] & upperMask) | (mt[kk+1] & lowerMask)
			mt[kk] = mt[kk+m] ^ (y >> 1) ^ mag01[y&0x1]
		}
		for ; kk < n-1; kk++ {
			y = (mt[kk] & upperMask) | (mt[kk+1] & lowerMask)
			mt[kk] = mt[kk+(m-n)] ^ (y >> 1) ^ mag01[y&0x1]
		}
		y = (mt[n-1] & upperMask) | (mt[0] & lowerMask)
		mt[n-1] = mt[m-1] ^ (y >> 1) ^ mag01[y&0x1]

		mti = 0
	}

	y = mt[mti]
	mti++

	y ^= (y >> 11)
	y ^= (y << 7) & 0x9d2c5680
	y ^= (y << 15) & 0xefc60000
	y ^= (y >> 18)

	return y
}

/* GenrandInt31 generates a 31bit unsigned int random number */
/* note: return type is uint32 */
func GenrandInt31() uint32 {
	return uint32(GenrandInt32() >> 1)
}

/* GenrandReal1 generates a 32bit [0, 1] real random number */
/* note: return type is float64, not float32 */
func GenrandReal1() float64 {
	return float64(GenrandInt32()) * (1.0 / 4294967295.0)
}

/* GenrandReal2 generates a 32bit [0, 1) real random number */
/* note: return type is float64, not float32 */
func GenrandReal2() float64 {
	return float64(GenrandInt32()) * (1.0 / 4294967296.0)
}

/* GenrandReal3 generates a 32bit (0, 1) real random  number */
/* note: return type is float64, not float32 */
func genrandReal3() float64 {
	return ((float64(GenrandInt32())) + 0.5) * (1.0 / 4294967296.0)
}

/* GenrandRes53 generates a [0, 1) random number with 53-bit resolution*/
func genrandRes53() float64 {
	var a uint32 = GenrandInt32() >> 5
	var b uint32 = GenrandInt32() >> 6
	return (float64(a)*67108864.0 + float64(b)) * (1.0 / 9007199254740992.0)
}
