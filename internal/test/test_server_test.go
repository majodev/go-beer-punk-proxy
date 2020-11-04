package test_test

import (
	"net/http"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/labstack/echo/v4"
	"github.com/majodev/go-beer-punk-proxy/internal/api"
	"github.com/majodev/go-beer-punk-proxy/internal/test"
	"github.com/majodev/go-beer-punk-proxy/internal/util"
	"github.com/stretchr/testify/assert"
)

type TestRequestPayload struct {
	Name string `json:"name"`
}

type TestResponsePayload struct {
	Hello string `json:"hello"`
}

// Validate validates this request payload
func (m *TestRequestPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestRequestPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestRequestPayload) UnmarshalBinary(b []byte) error {
	var res TestRequestPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Validate validates this response payload
func (m *TestResponsePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestResponsePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestResponsePayload) UnmarshalBinary(b []byte) error {
	var res TestResponsePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func TestWithTestServer(t *testing.T) {

	t.Parallel()

	test.WithTestServer(t, func(s1 *api.Server) {
		test.WithTestServer(t, func(s2 *api.Server) {

			path := "/testing-f679dbac-62bb-445d-b7e8-9f2c71ca382c"

			// add an new route to s1 for the purpose of this test.
			s1.Echo.POST(path, func(c echo.Context) error {

				var body TestRequestPayload
				if err := util.BindAndValidate(c, &body); err != nil {
					t.Fatal(err)
				}

				response := TestResponsePayload{
					Hello: body.Name,
				}

				return util.ValidateAndReturn(c, http.StatusOK, &response)
			})

			payload := test.GenericPayload{
				"name": "Mario",
			}

			res1 := test.PerformRequest(t, s1, "POST", path, payload, nil)
			assert.Equal(t, http.StatusOK, res1.Result().StatusCode)

			var response1 TestResponsePayload
			test.ParseResponseAndValidate(t, res1, &response1)

			assert.Equal(t, "Mario", response1.Hello)

			res2 := test.PerformRequest(t, s2, "POST", path, payload, nil)
			assert.Equal(t, http.StatusNotFound, res2.Result().StatusCode)

		})
	})

}
