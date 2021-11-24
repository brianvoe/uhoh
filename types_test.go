package uhoh

import (
	"errors"
	"testing"
)

func TestError_Types(t *testing.T) {
	var scenarios = []struct {
		err     *Err
		typeErr error
	}{
		{
			err:     New(errors.New("something error happened")).SetType(ErrGeneral),
			typeErr: ErrGeneral,
		},
		{
			err:     New(errors.New("the request of not correct")).SetType(ErrBadRequest),
			typeErr: ErrBadRequest,
		},
		{
			err:     New(errors.New("validation of payment method failed")).SetType(ErrValidation),
			typeErr: ErrValidation,
		},
		{
			err:     New(errors.New("merchant not found")).SetType(ErrForbidden),
			typeErr: ErrForbidden,
		},
		{
			err:     New(errors.New("merchant not allowed to do this")).SetType(ErrPermission),
			typeErr: ErrPermission,
		},
		{
			err:     New(errors.New("unknown api key")).SetType(ErrUnauthorized),
			typeErr: ErrUnauthorized,
		},
		{
			err:     New(errors.New("sql connection error")).SetType(ErrDBConnection),
			typeErr: ErrDBConnection,
		},
		{
			err:     New(errors.New("sql query error")).SetType(ErrDBQuery),
			typeErr: ErrDBQuery,
		},
		{
			err:     New(errors.New("there is no row for that transaction id")).SetType(ErrDBNoRows),
			typeErr: ErrDBNoRows,
		},
		{
			err:     New(errors.New("invalid memory address")).SetType(ErrInternal),
			typeErr: ErrInternal,
		},
	}

	for _, scenario := range scenarios {
		if !errors.Is(scenario.err, scenario.typeErr) {
			t.Errorf("expected error type %s, got %s", scenario.typeErr, scenario.err.Type)
		}
	}
}
