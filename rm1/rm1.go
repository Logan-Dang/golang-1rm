// Package golang1rm provides functions to calculate one-rep max (1RM) estimates
// using various formulas and to predict repetitions based on a given 1RM.
package rm1

const (
	Version = "1.0.1"
)

// Rm1Formula represents the different formulas used for 1RM calculations.
type Rm1Formula string

const (
	// The Epley formula tends to be most accurate between 1 and 10 reps.
	Epley Rm1Formula = "Epley"

	// The Brzycki formula is often used between 1 and 5 reps.
	Brzycki Rm1Formula = "Brzycki"

	// The Lombardi formula is typically used for higher rep ranges of 10+.
	Lombardi Rm1Formula = "Lombardi"

	// The Mayhew formula was derived from a bench press study, and so is most popular for bench press calculations.
	Mayhew Rm1Formula = "Mayhew"

	// The Wathan formula is a more conservative version of the Mayhew formula.
	Wathan Rm1Formula = "Wathan"

	// Specify this to have the formula selection default to the most appropriate one based on reps.
	Default Rm1Formula = "Default"
)

// Rm1Epley calculates the one-rep max using the Epley formula.
func Rm1Epley(weight float64, reps float64) float64 {
	return weight * (1 + reps/30)
}

// Rm1Brzycki calculates the one-rep max using the Brzycki formula.
func Rm1Brzycki(weight float64, reps float64) float64 {
	return weight / (1.0278 - 0.0278*reps)
}

// Rm1Lombardi calculates the one-rep max using the Lombardi formula.
func Rm1Lombardi(weight float64, reps float64) float64 {
	return weight * (1 + reps/40)
}

// Rm1Mayhew calculates the one-rep max using the Mayhew formula.
func Rm1Mayhew(weight float64, reps float64) float64 {
	return weight * (100 / (52.2 + 41.9*reps/100))
}

// Rm1Wathan calculates the one-rep max using the Wathan formula.
func Rm1Wathan(weight float64, reps float64) float64 {
	return weight * (100 / (48.8 + 53.8*reps/100))
}

// Rm1Default selects the most appropriate formula based on the number of reps.
func Rm1Default(weight float64, reps float64) float64 {
	if reps <= 1 {
		return weight
	}
	if reps <= 5 {
		return Rm1Brzycki(weight, reps)
	}
	if reps <= 10 {
		return Rm1Epley(weight, reps)
	}
	return Rm1Wathan(weight, reps)
}

// Rm1All calculates the one-rep max using all formulas.
func Rm1All(weight float64, reps float64) map[Rm1Formula]float64 {
	return map[Rm1Formula]float64{
		Epley:    Rm1Epley(weight, reps),
		Brzycki:  Rm1Brzycki(weight, reps),
		Lombardi: Rm1Lombardi(weight, reps),
		Mayhew:   Rm1Mayhew(weight, reps),
		Wathan:   Rm1Wathan(weight, reps),
	}
}

// Rm1 calculates the one-rep max using the specified formula.
func Rm1(weight float64, reps float64, formula Rm1Formula) float64 {
	switch formula {
	case Epley:
		return Rm1Epley(weight, reps)
	case Brzycki:
		return Rm1Brzycki(weight, reps)
	case Lombardi:
		return Rm1Lombardi(weight, reps)
	case Mayhew:
		return Rm1Mayhew(weight, reps)
	case Wathan:
		return Rm1Wathan(weight, reps)
	default:
		return Rm1Default(weight, reps)
	}
}

// RepPredictEpley predicts repetitions based on a given 1RM using an inverse of the Epley formula.
func RepPredictEpley(rm1 float64, weight float64) float64 {
	if rm1 <= 0 {
		return 0
	}
	return 30 * (rm1/weight - 1)
}

// RepPredictBrzycki predicts repetitions based on a given 1RM using an inverse of the Brzycki formula.
func RepPredictBrzycki(rm1 float64, weight float64) float64 {
	if rm1 <= 0 {
		return 0
	}
	return (1.0278 - rm1/weight) / 0.0278
}

// RepPredictLombardi predicts repetitions based on a given 1RM using an inverse of the Lombardi formula.
func RepPredictLombardi(rm1 float64, weight float64) float64 {
	if rm1 <= 0 {
		return 0
	}
	return 40 * (rm1/weight - 1)
}

// RepPredictMayhew predicts repetitions based on a given 1RM using an inverse of the Mayhew formula.
func RepPredictMayhew(rm1 float64, weight float64) float64 {
	if rm1 <= 0 {
		return 0
	}
	return 100 * (52.2 - 100*weight/rm1) / 41.9
}

// RepPredictWathan predicts repetitions based on a given 1RM using an inverse of the Wathan formula.
func RepPredictWathan(rm1 float64, weight float64) float64 {
	if rm1 <= 0 {
		return 0
	}
	return 100 * (48.8 - 100*weight/rm1) / 53.8
}

// RepPredictAll predicts repetitions using all formulas based on a given 1RM.
func RepPredictAll(rm1 float64, weight float64) map[Rm1Formula]float64 {
	return map[Rm1Formula]float64{
		Epley:    RepPredictEpley(rm1, weight),
		Brzycki:  RepPredictBrzycki(rm1, weight),
		Lombardi: RepPredictLombardi(rm1, weight),
		Mayhew:   RepPredictMayhew(rm1, weight),
		Wathan:   RepPredictWathan(rm1, weight),
	}
}

// RepPredict calculates repetitions based on a given 1RM using the specified formula.
func RepPredict(rm1 float64, weight float64, formula Rm1Formula) float64 {
	switch formula {
	case Brzycki:
		return RepPredictBrzycki(rm1, weight)
	case Lombardi:
		return RepPredictLombardi(rm1, weight)
	case Mayhew:
		return RepPredictMayhew(rm1, weight)
	case Wathan:
		return RepPredictWathan(rm1, weight)
	default:
		return RepPredictEpley(rm1, weight)
	}
}
