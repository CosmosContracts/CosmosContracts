package types

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestPhaseInflation(t *testing.T) {
	minter := DefaultInitialMinter()

	// Governing Mechanism:
	//    Juno tokenomics

	tests := []struct {
		phase        uint64
		expInflation sdk.Dec
	}{
		// phase 1, inflation: 40%
		{1, sdk.NewDecWithPrec(40, 2)},
		// phase 2, inflation: 20%
		{2, sdk.NewDecWithPrec(20, 2)},
		// phase 3, inflation: 10%
		{3, sdk.NewDecWithPrec(10, 2)},
		// phase 4, inflation: 9%
		{4, sdk.NewDecWithPrec(9, 2)},
		// phase 5, inflation: 8%
		{5, sdk.NewDecWithPrec(8, 2)},
		// phase 6, inflation: 7%
		{6, sdk.NewDecWithPrec(7, 2)},
		// phase 7, inflation: 6%
		{7, sdk.NewDecWithPrec(6, 2)},
		// phase 8, inflation: 5%
		{8, sdk.NewDecWithPrec(5, 2)},
		// phase 9, inflation: 4%
		{9, sdk.NewDecWithPrec(4, 2)},
		// phase 10, inflation: 3%
		{10, sdk.NewDecWithPrec(3, 2)},
		// phase 11, inflation: 2%
		{11, sdk.NewDecWithPrec(2, 2)},
		// phase 12, inflation: 1%
		{12, sdk.NewDecWithPrec(1, 2)},
		// end phase, inflation: 0%
		{13, sdk.NewDecWithPrec(0, 2)},
		// end phase, inflation: 0%
		{23, sdk.NewDecWithPrec(0, 2)},
	}
	for i, tc := range tests {
		inflation := minter.PhaseInflationRate(tc.phase)

		require.True(t, inflation.Equal(tc.expInflation),
			"Test Index: %v\nInflation:  %v\nExpected: %v\n", i, inflation, tc.expInflation)
	}
}

func TestBlockProvision(t *testing.T) {
	minter := InitialMinter(sdk.NewDecWithPrec(1, 1))
	params := DefaultParams()

	secondsPerYear := int64(60 * 60 * 8766)

	tests := []struct {
		annualProvisions int64
		expProvisions    int64
	}{
		{secondsPerYear / 5, 1},
		{secondsPerYear/5 + 1, 1},
		{(secondsPerYear / 5) * 2, 2},
		{(secondsPerYear / 5) / 2, 0},
	}
	for i, tc := range tests {
		minter.AnnualProvisions = sdk.NewDec(tc.annualProvisions)
		provisions := minter.BlockProvision(params)

		expProvisions := sdk.NewCoin(params.MintDenom,
			sdk.NewInt(tc.expProvisions))

		require.True(t, expProvisions.IsEqual(provisions),
			"test: %v\n\tExp: %v\n\tGot: %v\n",
			i, tc.expProvisions, provisions)
	}
}

// Benchmarking :)
// previously using sdk.Int operations:
// BenchmarkBlockProvision-4 5000000 220 ns/op
//
// using sdk.Dec operations: (current implementation)
// BenchmarkBlockProvision-4 3000000 429 ns/op
func BenchmarkBlockProvision(b *testing.B) {
	minter := InitialMinter(sdk.NewDecWithPrec(1, 1))
	params := DefaultParams()

	s1 := rand.NewSource(100)
	r1 := rand.New(s1)
	minter.AnnualProvisions = sdk.NewDec(r1.Int63n(1000000))

	// run the BlockProvision function b.N times
	for n := 0; n < b.N; n++ {
		minter.BlockProvision(params)
	}
}

// Next inflation benchmarking
// BenchmarkPhaseInflation-4 1000000 1828 ns/op
func BenchmarkPhaseInflation(b *testing.B) {
	minter := InitialMinter(sdk.NewDecWithPrec(1, 1))
	phase := uint64(4)

	// run the NextInflationRate function b.N times
	for n := 0; n < b.N; n++ {
		minter.PhaseInflationRate(phase)
	}

}

// Next annual provisions benchmarking
// BenchmarkNextAnnualProvisions-4 5000000 251 ns/op
func BenchmarkNextAnnualProvisions(b *testing.B) {
	minter := InitialMinter(sdk.NewDecWithPrec(1, 1))
	params := DefaultParams()
	totalSupply := sdk.NewInt(100000000000000)

	// run the NextAnnualProvisions function b.N times
	for n := 0; n < b.N; n++ {
		minter.NextAnnualProvisions(params, totalSupply)
	}

}
