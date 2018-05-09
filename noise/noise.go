package noise

import (
	"math"
	"runtime"
	"sync"

	"github.com/PrawnSkunk/genetic-al-go-rithm/lib/snoise2"
)

// Type indicates which noise type to generate
type Type int

const (
	// FBM is noise type 1
	FBM Type = iota
	// TURBULENCE is noise type 2
	TURBULENCE
)

// MakeNoise generates a 2D block of noise
func MakeNoise(Type Type, frequency, lacunarity, gain float32, octaves, w, h int) (noise []float32, min, max float32) {
	// Because we declare these in our return, we don't need to declare, only initialize
	noise = make([]float32, w*h)
	innerMin := float32(math.MaxFloat32)
	innerMax := float32(-math.MaxFloat32)

	// Set num goroutines
	numRoutines := runtime.NumCPU()
	batchSize := len(noise) / numRoutines

	// Channels, list goroutines can put things and pass to others
	minMaxChan := make(chan float32, numRoutines*2) // min and max in 1 channel (so *2)

	// Make waitgroup
	var wg sync.WaitGroup
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go func(i int) {
			defer wg.Done()
			start := i * batchSize
			end := start + batchSize - 1
			for j := start; j < end; j++ {
				x := j % w
				y := (j - x) / w

				// Determine noise type, using branch prediction
				if Type == TURBULENCE {
					noise[j] = Turbulence(float32(x), float32(y), frequency, lacunarity, gain, octaves)
				} else if Type == FBM {
					noise[j] = Fbm2(float32(x), float32(y), frequency, lacunarity, gain, octaves)
				}

				// Check if we have noise below/above the min/max
				if noise[j] < innerMin {
					innerMin = noise[j]
				} else if noise[j] > innerMax {
					innerMax = noise[j]
				}
			}

			// Put values in channel
			minMaxChan <- innerMin
			minMaxChan <- innerMax
		}(i)
	}
	wg.Wait()

	// Close our channel
	close(minMaxChan)

	// Iterate over channel
	for v := range minMaxChan {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}

	return noise, min, max
}

// Fbm2 genereates Fractional Brownian motion noise
func Fbm2(x, y, frequency, lacunarity, gain float32, octaves int) float32 {
	var sum float32
	amplitude := float32(1.0)
	for i := 0; i < octaves; i++ {
		sum += snoise2.Snoise2(x*frequency, y*frequency) * amplitude
		frequency = frequency * lacunarity
		amplitude = amplitude * gain
	}
	return sum
}

// Turbulence generates turbulent fractal noise
func Turbulence(x, y, frequency, lacuarity, gain float32, octaves int) float32 {
	var sum float32
	amplitude := float32(1)
	for i := 0; i < octaves; i++ {
		f := snoise2.Snoise2(x*frequency, y*frequency) * amplitude
		if f < 0 {
			f = -1.0 * f
		}
		sum += f
		frequency *= lacuarity
		amplitude *= gain
	}
	return sum
}
