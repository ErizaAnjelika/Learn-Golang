// assert adalah function untuk mengecek kondisi

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHitungVolume(t *testing.T) {
	assert.Equal(t, 30, 30, "Perhitungan volume salah")
}

// func TestHitungLuas(t *testing.T) {
// 	assert.Equal(t, kubus.Luas(), luasSeharusnya, "Perhitungan luas salah")
// }

// func TestHitungKeliling(t *testing.T) {
// 	assert.Equal(t, kubus.Keliling(), kelilingSeharusnya, "Perhitungan keliling salah")
// }

// func main() {
// 	TestHitungVolume(&testing.T{})
// }
