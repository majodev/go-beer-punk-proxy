package httperrors_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/go-openapi/swag"
	"github.com/majodev/go-beer-punk-proxy/internal/api/httperrors"
	"github.com/majodev/go-beer-punk-proxy/internal/types"
	"github.com/stretchr/testify/require"
)

func TestHTTPErrorSimple(t *testing.T) {
	t.Parallel()
	e := httperrors.NewHTTPError(http.StatusNotFound, httperrors.HTTPErrorTypeGeneric, http.StatusText(http.StatusNotFound))
	require.Equal(t, "HTTPError 404 (generic): Not Found", e.Error())
}

func TestHTTPErrorDetail(t *testing.T) {
	t.Parallel()
	e := httperrors.NewHTTPErrorWithDetail(http.StatusNotFound, httperrors.HTTPErrorTypeGeneric, http.StatusText(http.StatusNotFound), "ToS violation")
	require.Equal(t, "HTTPError 404 (generic): Not Found - ToS violation", e.Error())
}

func TestHTTPErrorInternalError(t *testing.T) {
	t.Parallel()
	e := httperrors.NewHTTPError(http.StatusInternalServerError, httperrors.HTTPErrorTypeGeneric, http.StatusText(http.StatusInternalServerError))

	e.Internal = sql.ErrConnDone

	require.Equal(t, "HTTPError 500 (generic): Internal Server Error, sql: connection is already closed", e.Error())
}

func TestHTTPErrorAdditionalData(t *testing.T) {
	t.Parallel()
	e := httperrors.NewHTTPError(http.StatusInternalServerError, httperrors.HTTPErrorTypeGeneric, http.StatusText(http.StatusInternalServerError))

	e.AdditionalData = map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	require.Equal(t, "HTTPError 500 (generic): Internal Server Error. Additional: key1=value1, key2=value2", e.Error())
}

var valErrs = append(make([]*types.HTTPValidationErrorDetail, 0, 2), &types.HTTPValidationErrorDetail{
	Key:   swag.String("test1"),
	In:    swag.String("body.test1"),
	Error: swag.String("ValidationError"),
}, &types.HTTPValidationErrorDetail{
	Key:   swag.String("test2"),
	In:    swag.String("body.test2"),
	Error: swag.String("Validation Error"),
})

func TestHTTPValidationErrorSimple(t *testing.T) {
	t.Parallel()

	e := httperrors.NewHTTPValidationError(http.StatusBadRequest, httperrors.HTTPErrorTypeGeneric, http.StatusText(http.StatusBadRequest), valErrs)
	require.Equal(t, "HTTPValidationError 400 (generic): Bad Request - Validation: test1 (in body.test1): ValidationError, test2 (in body.test2): Validation Error", e.Error())
}

func TestHTTPValidationErrorDetail(t *testing.T) {
	t.Parallel()
	e := httperrors.NewHTTPValidationErrorWithDetail(http.StatusBadRequest, httperrors.HTTPErrorTypeGeneric, http.StatusText(http.StatusBadRequest), valErrs, "Did API spec change?")
	require.Equal(t, "HTTPValidationError 400 (generic): Bad Request - Did API spec change? - Validation: test1 (in body.test1): ValidationError, test2 (in body.test2): Validation Error", e.Error())
}

func TestHTTPValidationErrorInternalError(t *testing.T) {
	t.Parallel()
	e := httperrors.NewHTTPValidationError(http.StatusBadRequest, httperrors.HTTPErrorTypeGeneric, http.StatusText(http.StatusBadRequest), valErrs)

	e.Internal = sql.ErrConnDone

	require.Equal(t, "HTTPValidationError 400 (generic): Bad Request, sql: connection is already closed - Validation: test1 (in body.test1): ValidationError, test2 (in body.test2): Validation Error", e.Error())
}

func TestHTTPValidationErrorAdditionalData(t *testing.T) {
	t.Parallel()
	e := httperrors.NewHTTPValidationError(http.StatusBadRequest, httperrors.HTTPErrorTypeGeneric, http.StatusText(http.StatusBadRequest), valErrs)

	e.AdditionalData = map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	require.Equal(t, "HTTPValidationError 400 (generic): Bad Request. Additional: key1=value1, key2=value2 - Validation: test1 (in body.test1): ValidationError, test2 (in body.test2): Validation Error", e.Error())
}
