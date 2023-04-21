package validator

import (
	"github.com/stretchr/testify/assert"
	"github.com/to-com/wp/internal/dto"
	"testing"
)

func TestValidateWavesNotEmpty(t *testing.T) {
	t.Parallel()

	var waves []rawWave
	err := validateWavesNotEmpty(waves)

	assert.NotNil(t, err)
	assert.Equal(t, &WaveError{isCritical: true, Err: ErrEmptyWaves}, err)
}

func TestValidateRequiredFields(t *testing.T) {
	t.Parallel()

	t.Run("when wave plan contains empty required keys then critical error returned",
		func(t *testing.T) {
			t.Parallel()
			waves := []rawWave{
				{
					cutoffTime: "14:59",
					fromTime:   "16:00",
				},
			}
			err := validateRequiredFields(waves)

			assert.NotNil(t, err)
			assert.Equal(t, &WaveError{isCritical: true, Err: ErrEmptyRequiredField}, err)
		})

	t.Run("when wave plan contains empty value for required field then critical error returned",
		func(t *testing.T) {
			t.Parallel()
			waves := []rawWave{
				{
					cutoffTime: "14:59",
					fromTime:   "16:00",
					toTime:     "",
				},
			}
			err := validateRequiredFields(waves)

			assert.NotNil(t, err)
			assert.Equal(t, &WaveError{isCritical: true, Err: ErrEmptyRequiredField}, err)
		})
}

func TestTimeFormatValidation(t *testing.T) {
	t.Parallel()

	t.Run("when wave plan contains time in the wrong format then critical error returned",
		func(t *testing.T) {
			t.Parallel()
			waves := []rawWave{
				{
					cutoffTime: "14:61",
					fromTime:   "16:00",
					toTime:     "18:00",
				},
			}
			err := validateTimeFormat(defaultFieldsToValidate)(waves)

			assert.NotNil(t, err)
			assert.Equal(t, &WaveError{
				isCritical: true,
				Err:        ErrWrongTimeFormat,
				Cutoff:     "14:61",
				ErrFields:  []string{cutoffTime},
			}, err)
		})

	t.Run("when wave plan contains isps schedule in the wrong format then critical error returned",
		func(t *testing.T) {
			t.Parallel()
			waves := []rawWave{
				{
					cutoffTime:     "14:59",
					fromTime:       "16:00",
					toTime:         "18:00",
					prelimSchedule: "12:61",
				},
			}
			fieldsToValidate := append(defaultFieldsToValidate, ispsFieldsToValidate...)
			err := validateTimeFormat(fieldsToValidate)(waves)

			assert.NotNil(t, err)
			assert.Equal(t, &WaveError{
				isCritical: true,
				Err:        ErrWrongTimeFormat,
				Cutoff:     "14:59",
				ErrFields:  []string{prelimSchedule},
			}, err)
		})
}

func TestValidateAllCutoffUnique(t *testing.T) {
	t.Parallel()

	waves := []rawWave{
		{
			cutoffTime: "10:00",
			fromTime:   "14:00",
			toTime:     "16:00",
		},
		{
			cutoffTime: "10:00",
			fromTime:   "18:00",
			toTime:     "20:00",
		},
	}
	err := validateAllCutoffUnique(waves)

	assert.NotNil(t, err)
	assert.Equal(t, &WaveError{
		isCritical: true,
		Err:        ErrSameCutoffTime,
		Cutoff:     "10:00",
		ErrFields:  []string{cutoffTime},
	}, err)
}

func TestValidateCutoffBeforeStageTime(t *testing.T) {
	t.Parallel()

	waves := []rawWave{
		{
			cutoffTime: "18:30",
			fromTime:   "18:00",
			toTime:     "19:00",
		},
	}
	err := validateCutoffBeforeStageTime(waves)

	assert.NotNil(t, err)
	assert.Equal(t, &WaveError{
		isCritical: true,
		Err:        ErrIncorrectCutoff,
		Cutoff:     "18:30",
		ErrFields:  []string{cutoffTime},
	}, err)
}

func TestValidateDifferentStagePeriod(t *testing.T) {
	t.Parallel()

	waves := []rawWave{
		{
			cutoffTime: "14:00",
			fromTime:   "18:00",
			toTime:     "18:00",
		},
	}
	err := validateDifferentStagePeriod(waves)

	if assert.NotNil(t, err) {
		assert.Equal(t, &WaveError{
			isCritical: true,
			Err:        ErrSameFromToTime,
			Cutoff:     "14:00",
			ErrFields:  []string{fromTime, toTime},
		}, err)
	}
}

func TestValidateDifferentStagePeriodAndCutoffBeforeStageTime(t *testing.T) {
	t.Parallel()

	waves := []rawWave{
		{
			cutoffTime: "14:00",
			fromTime:   "18:00",
			toTime:     "18:00",
		},
		{
			cutoffTime: "15:00",
			fromTime:   "18:01",
			toTime:     "19:00",
		},
	}
	err := validateDifferentStagePeriod(waves)

	if assert.NotNil(t, err) {
		assert.Equal(t, &WaveError{
			isCritical: true,
			Err:        ErrSameFromToTime,
			Cutoff:     "14:00",
			ErrFields:  []string{fromTime, toTime},
		}, err)
	}
}

func TestValidateConfigCorrectness(t *testing.T) {
	t.Parallel()

	waves := []rawWave{
		{
			cutoffTime:     "14:59",
			fromTime:       "16:00",
			toTime:         "18:00",
			prelimSchedule: "12:00",
			deltaSchedule:  "13:00",
		},
	}
	errs := validateConfigCorrectness(false)(waves)

	if assert.NotNil(t, errs) {
		assert.Len(t, errs, 1)
		assert.Equal(t, &WaveError{
			Err:       ErrISPSTurnedOff,
			Cutoff:    "14:59",
			ErrFields: []string{prelimSchedule, deltaSchedule},
		}, errs[0])
	}
}

func TestValidateDifferentISPSTriggersTime(t *testing.T) {
	t.Parallel()

	waves := []rawWave{
		{
			cutoffTime:     "14:00",
			fromTime:       "16:00",
			toTime:         "18:00",
			prelimSchedule: "12:00",
			deltaSchedule:  "12:00",
		},
	}
	errs := validateDifferentISPSTriggersTime(waves)

	if assert.NotNil(t, errs) {
		assert.Len(t, errs, 1)
		assert.Equal(t, &WaveError{
			Err:       ErrSameISPSTriggers,
			Cutoff:    "14:00",
			ErrFields: []string{prelimSchedule, deltaSchedule},
		}, errs[0])
	}
}

func TestValidateProperlyOrderedWaves(t *testing.T) {
	t.Parallel()

	t.Run("when waves are not ordered properly then error returned", func(t *testing.T) {
		t.Parallel()
		waves := []rawWave{
			{
				cutoffTime: "09:00",
				fromTime:   "10:00",
				toTime:     "12:00",
			},
			{
				cutoffTime: "08:00",
				fromTime:   "14:00",
				toTime:     "16:00",
			},
		}
		errs := validateProperlyOrderedWaves(waves)

		if assert.NotNil(t, errs) {
			assert.Len(t, errs, 2)
			assert.Equal(t, &WaveError{
				Err:       ErrWrongOrderedWaves,
				Cutoff:    "09:00",
				ErrFields: []string{cutoffTime, fromTime, toTime},
			}, errs[0])
			assert.Equal(t, &WaveError{
				Err:       ErrWrongOrderedWaves,
				Cutoff:    "08:00",
				ErrFields: []string{cutoffTime, fromTime, toTime},
			}, errs[1])
		}
	})

	t.Run("when waves are properly ordered", func(t *testing.T) {
		t.Parallel()
		waves := []rawWave{
			{
				cutoffTime: "23:00",
				fromTime:   "02:00",
				toTime:     "04:00",
			},
			{
				cutoffTime: "00:00",
				fromTime:   "04:00",
				toTime:     "06:00",
			},
		}
		errs := validateProperlyOrderedWaves(waves)
		assert.Nil(t, errs)
	})

	t.Run("when waves are properly ordered and cutoffs in different days", func(t *testing.T) {
		t.Parallel()
		waves := []rawWave{
			{
				cutoffTime: "22:00",
				fromTime:   "02:00",
				toTime:     "10:00",
			},
			{
				cutoffTime: "05:00",
				fromTime:   "10:00",
				toTime:     "02:00",
			},
		}
		errs := validateProperlyOrderedWaves(waves)
		assert.Nil(t, errs)
	})
}

func TestValidateNoOverlapping(t *testing.T) {
	t.Parallel()

	t.Run("when two waves are overlap", func(t *testing.T) {
		t.Parallel()
		waves := []rawWave{
			{
				cutoffTime: "09:00",
				fromTime:   "10:00",
				toTime:     "16:00",
			},
			{
				cutoffTime: "10:00",
				fromTime:   "11:00",
				toTime:     "18:00",
			},
		}
		errs := validateNoOverlapping(waves)
		if assert.NotNil(t, errs) {
			assert.Len(t, errs, 2)
			assert.Equal(t, &WaveError{
				Err:       ErrWavesOverlapping,
				Cutoff:    "09:00",
				ErrFields: []string{cutoffTime, fromTime, toTime},
			}, errs[0])
			assert.Equal(t, &WaveError{
				Err:       ErrWavesOverlapping,
				Cutoff:    "10:00",
				ErrFields: []string{cutoffTime, fromTime, toTime},
			}, errs[1])
		}
	})

	t.Run("when two waves don't overlap, case #1", func(t *testing.T) {
		t.Parallel()
		waves := []rawWave{
			{
				cutoffTime: "09:00",
				fromTime:   "12:00",
				toTime:     "13:00",
			},
			{
				cutoffTime: "10:00",
				fromTime:   "13:00",
				toTime:     "14:00",
			},
		}
		errs := validateNoOverlapping(waves)
		assert.Nil(t, errs)
	})

	t.Run("when two waves don't overlap, case #2", func(t *testing.T) {
		t.Parallel()
		waves := []rawWave{
			{
				cutoffTime: "09:00",
				fromTime:   "19:00",
				toTime:     "19:59",
			},
			{
				cutoffTime: "11:00",
				fromTime:   "20:00",
				toTime:     "06:59",
			},
		}
		errs := validateNoOverlapping(waves)
		assert.Nil(t, errs)
	})
}

func TestValidateWavesCover24h(t *testing.T) {
	t.Parallel()

	t.Run("when waves cover 24h", func(t *testing.T) {
		t.Parallel()
		waves := []rawWave{
			{
				cutoffTime: "09:00",
				fromTime:   "00:00",
				toTime:     "15:59",
			},
			{
				cutoffTime: "15:00",
				fromTime:   "16:00",
				toTime:     "23:59",
			},
		}
		errs := validateWavesCover24h(waves)
		assert.Nil(t, errs)
	})

	t.Run("when waves don't cover 24h then error returned", func(t *testing.T) {
		t.Parallel()
		waves := []rawWave{
			{
				cutoffTime: "09:00",
				fromTime:   "00:00",
				toTime:     "15:59",
			},
			{
				cutoffTime: "15:00",
				fromTime:   "16:00",
				toTime:     "23:57",
			},
		}
		errs := validateWavesCover24h(waves)

		if assert.NotNil(t, errs) {
			assert.Len(t, errs, 1)
			assert.Equal(t, &WaveError{Err: ErrFullDayNotCovered}, errs[0])
		}
	})
}

func TestValidationErrorContainsOneCriticalError(t *testing.T) {
	t.Run("when contains two critical errors then only the first one will be return within ValidationError",
		func(t *testing.T) {
			wp := dto.wpRequest{Waves: []dto.WaveRequest{
				// first critical error - time in wrong format
				{Cutoff: "06:61", FromTime: "08:00", ToTime: "09:59"},
				// second critical error - empty required field
				{Cutoff: "08:00", FromTime: "", ToTime: "11:59"},
				{Cutoff: "10:00", FromTime: "12:00", ToTime: "13:59"}},
			}

			v := New()
			err := v.Validate(wp)

			if assert.NotNil(t, err) {
				assert.IsType(t, &ValidationError{}, err)
				assert.Len(t, err.WavesErrors, 1)
			}
		})
}

func TestValidatewps(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		isISPSEnabled bool
		input         dto.wpRequest
	}{
		"setup 1": {
			isISPSEnabled: false,
			input: dto.wpRequest{Waves: []dto.WaveRequest{
				{Cutoff: "06:00", FromTime: "08:00", ToTime: "09:59"},
				{Cutoff: "08:00", FromTime: "10:00", ToTime: "11:59"},
				{Cutoff: "10:00", FromTime: "12:00", ToTime: "13:59"},
				{Cutoff: "12:00", FromTime: "14:00", ToTime: "15:59"},
				{Cutoff: "14:00", FromTime: "16:00", ToTime: "07:59"},
			},
			},
		},
		"setup 2": {
			isISPSEnabled: false,
			input: dto.wpRequest{Waves: []dto.WaveRequest{
				{Cutoff: "05:00", FromTime: "09:00", ToTime: "10:59"},
				{Cutoff: "08:00", FromTime: "11:00", ToTime: "12:59"},
				{Cutoff: "11:00", FromTime: "13:00", ToTime: "14:59"},
				{Cutoff: "13:00", FromTime: "15:00", ToTime: "16:59"},
				{Cutoff: "15:00", FromTime: "17:00", ToTime: "18:59"},
				{Cutoff: "17:00", FromTime: "19:00", ToTime: "08:59"},
			},
			},
		},
		"setup 3": {
			isISPSEnabled: false,
			input: dto.wpRequest{Waves: []dto.WaveRequest{
				{Cutoff: "03:45", FromTime: "06:00", ToTime: "06:04"},
				{Cutoff: "05:45", FromTime: "06:05", ToTime: "10:00"},
				{Cutoff: "07:45", FromTime: "10:01", ToTime: "10:04"},
				{Cutoff: "09:45", FromTime: "10:05", ToTime: "11:29"},
				{Cutoff: "10:30", FromTime: "11:30", ToTime: "12:49"},
				{Cutoff: "12:30", FromTime: "12:50", ToTime: "05:59"},
			},
			},
		},
		"setup 4": {
			isISPSEnabled: true,
			input: dto.wpRequest{Waves: []dto.WaveRequest{
				{Cutoff: "06:30", FromTime: "09:00", ToTime: "10:59", PrelimTime: "04:00", DeltaTime: "05:00"},
				{Cutoff: "08:30", FromTime: "11:00", ToTime: "12:59"},
				{Cutoff: "11:00", FromTime: "13:00", ToTime: "14:59"},
				{Cutoff: "13:00", FromTime: "15:00", ToTime: "16:59"},
				{Cutoff: "15:00", FromTime: "17:00", ToTime: "18:59"},
				{Cutoff: "17:00", FromTime: "19:00", ToTime: "08:59"},
			},
			},
		},
	}

	for name, tt := range tests {
		// pin to be able running tests in parallel
		tt := tt
		var opts []Options
		if tt.isISPSEnabled {
			opts = append(opts, WithISPS())
		}
		v := New(opts...)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := v.Validate(tt.input)
			assert.Nil(t, err)
		})
	}
}
