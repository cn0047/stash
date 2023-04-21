package validator

import (
	"errors"
	"github.com/to-com/wp/internal/dto"
	"time"
)

type (
	rawWave map[string]string

	Options            func(*Validator)
	validateCriticalFn func([]rawWave) *WaveError
	validateLogicalFn  func([]rawWave) []*WaveError
)

// Validator struct is to validate provided wave plan against critical and logical errors
// isISPSEnabled - determines if InStorePicking validations must be run
// fieldsToValidate - runs validation against specified fields, can be enhanced if isISPSEnabled set to truth
// criticalValidations - list of functions to validate critical errors
// logicalValidations - list of functions to validate logical errors
// the difference between critical and logical validation implies returning an error immediately or keep iterating
// over validation functions
type Validator struct {
	isISPSEnabled       bool
	fieldsToValidate    []string
	criticalValidations []validateCriticalFn
	logicalValidations  []validateLogicalFn
}

type WaveError struct {
	isCritical bool
	Err        error
	Cutoff     string
	ErrFields  []string
}

func (we *WaveError) Error() string {
	return we.Err.Error()
}

func (we *WaveError) Unwrap() error {
	return we.Err
}

type ValidationError struct {
	Err         error
	WavesErrors []*WaveError
}

func (ve *ValidationError) Error() string {
	return "validation error occurred"
}

func (ve *ValidationError) Unwrap() error {
	return ve.Err
}

const (
	// timeLayout for all fields in the format HH:mm
	timeLayout = "15:04"

	required = "required"
	optional = "optional"

	cutoffTime     = "cutoff_time"
	fromTime       = "from_time"
	toTime         = "to_time"
	prelimSchedule = "prelim_picklist_schedule_time"
	deltaSchedule  = "delta_picklist_schedule_time"
)

var fieldsMeta = map[string]string{
	cutoffTime:     required,
	fromTime:       required,
	toTime:         required,
	prelimSchedule: optional,
	deltaSchedule:  optional,
}

var (
	defaultFieldsToValidate = []string{cutoffTime, fromTime, toTime}
	ispsFieldsToValidate    = []string{prelimSchedule, deltaSchedule}
)

var (
	// ErrEmptyWaves returned if no waves provided
	ErrEmptyWaves = errors.New("configuration must have at least one wave")
	// ErrEmptyRequiredField returned if one of the required fields is empty
	ErrEmptyRequiredField = errors.New("cutoff time, stage-by-datetime from, stage-by-datetime to fields cannot be empty")
	// ErrWrongTimeFormat returned in one of the time fields contain wrong format
	ErrWrongTimeFormat = errors.New("times must be in 24-hour time format (00:00 - 24:00). For example, 1:00 PM should be written as 13:00")
	// ErrIncorrectCutoff returned when cutoff field is not earlier than from_time
	ErrIncorrectCutoff = errors.New("cutoff times must be earlier than stage-by-datetime from and stage-by-datetime to period")
	// ErrSameCutoffTime returned when duplicated cutoffs found
	ErrSameCutoffTime = errors.New("cutoff times cannot be repeated")

	// ErrFullDayNotCovered returned if there is some uncovered time in a day
	ErrFullDayNotCovered = errors.New("wave configuration must cover a full 24-hour time period")
	// ErrWrongOrderedWaves returned in wave configured in incorrect order
	ErrWrongOrderedWaves = errors.New("the later stage-by-datetime must also have the later cutoff time")
	// ErrWavesOverlapping returned if waves overlap
	ErrWavesOverlapping = errors.New("waves cannot overlap")
	// ErrISPSTurnedOff returned if user provided InStorePicking schedules
	// but ISPS_ENABLED config turned off in the TSC
	ErrISPSTurnedOff = errors.New("cannot save schedules since InStorePicking configuration turned off")
	// ErrSameFromToTime returned when from and to period are the same
	ErrSameFromToTime = errors.New("stage-by-datetime from and stage-by-datetime to cannot be the same")
	// ErrSameISPSTriggers returned when time for triggers are the same
	ErrSameISPSTriggers = errors.New("in store picking triggers cannot be the same for cutoff")
)

// parseTime coerce string to time in HH:MM format
func parseTime(v string) (time.Time, error) {
	t, err := time.Parse(timeLayout, v)

	return t, err
}

// --- Critical Validations ---

// validateWavesNotEmpty validation doesn't allow client to provide empty waves
func validateWavesNotEmpty(waves []rawWave) *WaveError {
	if len(waves) > 0 {
		return nil
	}

	return &WaveError{
		isCritical: true,
		Err:        ErrEmptyWaves,
	}
}

// validateRequiredFields validation doesn't allow client to provide empty value for the required field
func validateRequiredFields(waves []rawWave) *WaveError {
	for _, wave := range waves {
		for fieldName, fieldMeta := range fieldsMeta {
			if fieldMeta == required {
				val, ok := wave[fieldName]
				if !ok || len(val) == 0 {
					return &WaveError{
						isCritical: true,
						Err:        ErrEmptyRequiredField,
					}
				}
			}
		}
	}

	return nil
}

// validateTimeFormat validation doesn't allow client to provide time fields in any formats except HH:mm
// additionally, it skips validation for optional fields if they are empty.
func validateTimeFormat(fieldsToValidate []string) validateCriticalFn {
	return func(waves []rawWave) *WaveError {
		for _, wave := range waves {
			for _, field := range fieldsToValidate {
				fieldMeta := fieldsMeta[field]
				if fieldMeta == optional && wave[field] == "" {
					continue
				}

				if _, err := parseTime(wave[field]); err != nil {
					return &WaveError{
						isCritical: true,
						Err:        ErrWrongTimeFormat,
						Cutoff:     wave[cutoffTime],
						ErrFields:  []string{field},
					}
				}
			}
		}

		return nil
	}
}

// validateAllCutoffUnique validation doesn't allow client to provide duplicated cutoffs.
func validateAllCutoffUnique(waves []rawWave) *WaveError {
	var m = make(map[string]bool)
	for _, wave := range waves {
		cutoff := wave[cutoffTime]
		if _, ok := m[cutoff]; ok {
			return &WaveError{
				isCritical: true,
				Err:        ErrSameCutoffTime,
				Cutoff:     cutoff,
				ErrFields:  []string{cutoffTime},
			}
		}
		m[cutoff] = true
	}

	return nil
}

// validateCutoffBeforeStageTime validation doesn't allow client to provide
// incorrect time for cutoff field (past from and to time, or between from and to time)
func validateCutoffBeforeStageTime(waves []rawWave) *WaveError {
	for _, wave := range waves {
		from, _ := parseTime(wave[fromTime])
		to, _ := parseTime(wave[toTime])
		cutoff, _ := parseTime(wave[cutoffTime])

		fromTimeFull, toTimeFull, cutoffFull := genTime(from), genTime(to), genTime(cutoff)

		// when from_time cross midnight
		if cutoffFull.After(fromTimeFull) {
			fromTimeFull = fromTimeFull.Add(24 * time.Hour)
		}

		// when to_time cross midnight
		if cutoffFull.After(toTimeFull) {
			toTimeFull = toTimeFull.Add(24 * time.Hour)
		}

		// cutoff must be always earlier from_time and to_time
		if !(cutoffFull.Before(fromTimeFull) && fromTimeFull.Before(toTimeFull)) {
			return &WaveError{
				isCritical: true,
				Err:        ErrIncorrectCutoff,
				Cutoff:     wave[cutoffTime],
				ErrFields:  []string{cutoffTime},
			}
		}
	}

	return nil
}

// validateDifferentStagePeriod validation doesn't allow client to provide
// same from and to time slots
func validateDifferentStagePeriod(waves []rawWave) *WaveError {
	for _, wave := range waves {
		if wave[fromTime] == wave[toTime] {
			return &WaveError{
				isCritical: true,
				Err:        ErrSameFromToTime,
				Cutoff:     wave[cutoffTime],
				ErrFields:  []string{fromTime, toTime},
			}
		}
	}

	return nil
}

// --- Logical Validations ---

// withoutCurrentWave function excludes specified wave from the list of waves
func withoutCurrentWave(wave rawWave, waves []rawWave) []rawWave {
	var rest []rawWave
	for _, w := range waves {
		if wave[cutoffTime] == w[cutoffTime] {
			continue
		}
		rest = append(rest, w)
	}

	return rest
}

// matchMidnightCross  checks if both cutoffs crossing midnight or not
func matchMidnightCross(wave1, wave2 rawWave) bool {
	var match bool

	if wave1[cutoffTime] < wave1[fromTime] && wave2[cutoffTime] < wave2[fromTime] {
		match = true
	}

	if wave1[cutoffTime] > wave1[fromTime] && wave2[cutoffTime] > wave2[fromTime] {
		match = true
	}

	return match
}

func properOrder(wave1, wave2 rawWave) bool {
	var ordered bool

	if matchMidnightCross(wave1, wave2) {
		if (wave1[cutoffTime] < wave2[cutoffTime] && wave1[fromTime] < wave2[fromTime]) ||
			(wave1[cutoffTime] > wave2[cutoffTime] && wave1[fromTime] > wave2[fromTime]) {
			ordered = true
		}
	} else {
		if (wave1[cutoffTime] < wave2[cutoffTime] && wave1[fromTime] > wave2[fromTime]) ||
			(wave1[cutoffTime] > wave2[cutoffTime] && wave1[fromTime] < wave2[fromTime]) {
			ordered = true
		}
	}

	return ordered
}

// validateProperlyOrderedWaves validation doesn't allow client to provide
// wrong cutoff ordering. Later from_time and to_time must have later cutoff
// example: Correct ordering
// Cutoff 1: 08:00, from 10:00 to 12:00
// Cutoff 2: 10:00, from 12:00 to 14:00
//
// Cutoff 1: 23:00, from 02:00 to 04:00
// Cutoff 2: 00:00, from 04:00 to 06:00
// ------------------------------------
// example: Incorrect ordering
// Cutoff 1: 10:00, from 10:00 to 12:00
// Cutoff 2: 08:00, from 14:00 to 16:00
//
// Cutoff 1: 23:00, from 02:00 to 04:00
// Cutoff 2: 22:00, from 04:00 to 06:00"
func validateProperlyOrderedWaves(waves []rawWave) []*WaveError {
	var errs []*WaveError
	for _, i := range waves {
		otherWaves := withoutCurrentWave(i, waves)
		for _, j := range otherWaves {
			if !properOrder(i, j) {
				errs = append(errs, &WaveError{
					Err:       ErrWrongOrderedWaves,
					Cutoff:    i[cutoffTime],
					ErrFields: []string{cutoffTime, fromTime, toTime},
				})
			}
		}
	}

	return errs
}

// isBetween checks that provided time (either from_time or to_time)
// must not overlap with neighbor from_time and to_time
// example:
// targetWave: from_time 20:00, to_time: 06:59
// neighborWave: from_time 19:00, to_time 19:59
// time       from       to_time       overlaps?
// 20:00      19:00      19:59         -
// 06:59      19:00      19:59         -
// 19:00      20:00      06:59         -
// 19:59      19:00      06:59         -
func isBetween(time, from, to string) bool {
	var between bool

	if from < to {
		if (from < time) && (time < to) {
			between = true
		}
	} else {
		if (from < time) || (time < to) {
			between = true
		}
	}

	return between
}

// wavesOverlap
// Windows like 01:00-02:00, 02:00-03:00 are not considered overlapping
// works for ranges crossing midnight
func wavesOverlap(wave1, wave2 rawWave) bool {
	var overlapping bool
	// overlapping considered for the same from_time or to_time
	if (wave1[fromTime] == wave2[fromTime] || wave1[toTime] == wave2[toTime]) ||
		isBetween(wave1[fromTime], wave2[fromTime], wave2[toTime]) ||
		isBetween(wave1[toTime], wave2[fromTime], wave2[toTime]) ||
		isBetween(wave2[fromTime], wave1[fromTime], wave1[toTime]) ||
		isBetween(wave2[toTime], wave1[fromTime], wave1[toTime]) {
		overlapping = true
	}

	return overlapping
}

// validateNoOverlapping validation doesn't allow client to provide
// overlapping between waves
func validateNoOverlapping(waves []rawWave) []*WaveError {
	var errs []*WaveError
	for _, i := range waves {
		otherWaves := withoutCurrentWave(i, waves)
		for _, j := range otherWaves {
			if wavesOverlap(i, j) {
				errs = append(errs, &WaveError{
					Err:       ErrWavesOverlapping,
					Cutoff:    i[cutoffTime],
					ErrFields: []string{cutoffTime, fromTime, toTime},
				})
			}
		}
	}

	return errs
}

func genTime(t time.Time) time.Time {
	return time.Date(2022, 01, 01, t.Hour(), t.Minute(), 0, 0, time.UTC)
}

// validateConfigCorrectness validation doesn't allow client to provide
// isps triggers if module itself turned off.
func validateConfigCorrectness(isISPSEnabled bool) validateLogicalFn {
	return func(waves []rawWave) []*WaveError {
		if isISPSEnabled {
			return nil
		}

		var errs []*WaveError
		for _, wave := range waves {
			var errFields []string
			if wave[prelimSchedule] != "" {
				errFields = append(errFields, prelimSchedule)
			}
			if wave[deltaSchedule] != "" {
				errFields = append(errFields, deltaSchedule)
			}

			if len(errFields) > 0 {
				errs = append(errs, &WaveError{
					Err:       ErrISPSTurnedOff,
					Cutoff:    wave[cutoffTime],
					ErrFields: errFields,
				})
			}
		}

		return errs
	}
}

// validateWavesCover24h validation doesn't allow client to provide
// waves configuration without covering an entire day
func validateWavesCover24h(waves []rawWave) []*WaveError {
	var OneDayInMin = 24 * 60
	var errs []*WaveError

	for _, wave := range waves {
		from, _ := parseTime(wave[fromTime])
		to, _ := parseTime(wave[toTime])

		fromTimeFull := genTime(from)
		toTimeFull := genTime(to)

		// if to_time is crossing midnight then add one more day to count a diff correctly
		if wave[fromTime] > wave[toTime] {
			toTimeFull = toTimeFull.Add(24 * time.Hour)
		}

		deltaMin := int(toTimeFull.Sub(fromTimeFull).Minutes())
		OneDayInMin = OneDayInMin - deltaMin
	}

	// subtracting leftovers minutes, since every wave has 1min gap between next wave
	// example: wave1: to_time 06:04, wave2: from_time 06:05 (1 min gap)
	OneDayInMin = OneDayInMin - len(waves)

	if OneDayInMin != 0 {
		err := &WaveError{Err: ErrFullDayNotCovered}
		errs = append(errs, err)
	}

	return errs
}

// validateDifferentISPSTriggersTime validation doesn't allow client to provide
// same trigger time for different picklists
func validateDifferentISPSTriggersTime(waves []rawWave) []*WaveError {
	var errs []*WaveError
	for _, wave := range waves {
		if wave[prelimSchedule] != "" && (wave[prelimSchedule] == wave[deltaSchedule]) {
			errs = append(errs, &WaveError{
				Err:       ErrSameISPSTriggers,
				Cutoff:    wave[cutoffTime],
				ErrFields: []string{prelimSchedule, deltaSchedule},
			})
		}
	}

	return errs
}

func WithISPS() Options {
	return func(v *Validator) {
		v.isISPSEnabled = true
		v.fieldsToValidate = append(v.fieldsToValidate, ispsFieldsToValidate...)
		v.logicalValidations = append(v.logicalValidations, validateDifferentISPSTriggersTime)
	}
}

func New(opts ...Options) *Validator {
	v := Validator{}

	for _, opt := range opts {
		opt(&v)
	}

	v.fieldsToValidate = append(v.fieldsToValidate, defaultFieldsToValidate...)

	criticalValidations := []validateCriticalFn{
		validateWavesNotEmpty,
		validateRequiredFields,
		validateTimeFormat(v.fieldsToValidate),
		validateAllCutoffUnique,
		validateDifferentStagePeriod,
		validateCutoffBeforeStageTime,
	}

	logicalValidations := []validateLogicalFn{
		validateConfigCorrectness(v.isISPSEnabled),
		validateWavesCover24h,
		validateProperlyOrderedWaves,
		validateNoOverlapping,
	}

	v.criticalValidations = append(v.criticalValidations, criticalValidations...)
	v.logicalValidations = append(v.logicalValidations, logicalValidations...)

	return &v
}

func (v *Validator) Validate(wp dto.wpRequest) *ValidationError {
	waves := make([]rawWave, 0, len(wp.Waves))
	for _, wave := range wp.Waves {
		rw := rawWave{
			cutoffTime:     wave.Cutoff,
			fromTime:       wave.FromTime,
			toTime:         wave.ToTime,
			prelimSchedule: wave.PrelimTime,
			deltaSchedule:  wave.DeltaTime,
		}

		waves = append(waves, rw)
	}

	var wavesErrors []*WaveError
	for _, cFn := range v.criticalValidations {
		if criticalErr := cFn(waves); criticalErr != nil {
			wavesErrors = append(wavesErrors, criticalErr)
			return &ValidationError{WavesErrors: wavesErrors}
		}
	}

	for _, lFn := range v.logicalValidations {
		logicalErrs := lFn(waves)
		if len(logicalErrs) > 0 {
			wavesErrors = append(wavesErrors, logicalErrs...)
		}
	}

	if len(wavesErrors) > 0 {
		return &ValidationError{WavesErrors: wavesErrors}
	}

	return nil
}
