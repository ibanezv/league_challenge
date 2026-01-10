package integrationtest

import (
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/league/league_challenge/internal/domain/usecase"
	httpApp "github.com/league/league_challenge/internal/infrastructure/http"
	"github.com/league/league_challenge/internal/infrastructure/http/middleware"
	"github.com/league/league_challenge/internal/server"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx/fxtest"
)

// Given a valid file request
// When Post to /echo
// Then It retrieves 200 status-code and valid body response
func TestHttpEchoSuccess(t *testing.T) {
	// Arrange
	expectedBody := "1,2,3\n4,5,6\n7,8,9\n"
	appTest := fxtest.New(t, server.NewServer())
	appTest.RequireStart()
	defer appTest.RequireStop()

	body, writer, _ := createFileRequest("matrixtest.csv")

	echoCtrl := httpApp.NewEchoController()
	test := NewFiberTest(t, middleware.ValidationInput, echoCtrl.Handler)
	req, _ := http.NewRequest(
		fiber.MethodPost,
		"/echo",
		body,
	)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Act
	res, err := test.TestRequest(req)

	// Assert
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, res.StatusCode)
	assert.Equal(t, expectedBody, bodyString)
}

// Given A not valid file request
// When Post to /echo
// Then It retrieves 400 Bad Request status-code
func TestHttpEchoFail(t *testing.T) {
	// Arrange
	appTest := fxtest.New(t, server.NewServer())
	appTest.RequireStart()
	defer appTest.RequireStop()

	body, writer, _ := createFileRequest("matrixinvalid.txt")

	echoCtrl := httpApp.NewEchoController()
	test := NewFiberTest(t, middleware.ValidationInput, echoCtrl.Handler)
	req, _ := http.NewRequest(
		fiber.MethodPost,
		"/echo",
		body,
	)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Act
	res, err := test.TestRequest(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}

// Given a valid file request
// When Post to /flatten
// Then It retrieves 200 status-code and valid body response
func TestHttpFlattenSuccess(t *testing.T) {
	// Arrange
	expectedBody := "1,2,3,4,5,6,7,8,9"
	appTest := fxtest.New(t, server.NewServer())
	appTest.RequireStart()
	defer appTest.RequireStop()

	body, writer, _ := createFileRequest("matrixtest.csv")

	uc := usecase.NewFlattenUseCase()
	flattenCtrl := httpApp.NewFlattenController(uc)
	test := NewFiberTest(t, flattenCtrl.Handler)
	req, _ := http.NewRequest(
		fiber.MethodPost,
		"/flatten",
		body,
	)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Act
	res, err := test.TestRequest(req)

	// Assert
	bodyBytes, errRead := io.ReadAll(res.Body)
	if errRead != nil {
		log.Fatal(errRead)
	}
	bodyString := string(bodyBytes)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, res.StatusCode)
	assert.Equal(t, expectedBody, bodyString)
}

// Given a valid file request
// When Post to /invert
// Then It retrieves 200 status-code and valid body response
func TestHttpInvertSuccess(t *testing.T) {
	// Arrange
	expectedBody := "1,4,7\n2,5,8\n3,6,9\n"
	appTest := fxtest.New(t, server.NewServer())
	appTest.RequireStart()
	defer appTest.RequireStop()

	body, writer, _ := createFileRequest("matrixtest.csv")

	uc := usecase.NewInvertUseCase()
	invertCtrl := httpApp.NewInvertController(uc)
	test := NewFiberTest(t, invertCtrl.Handler)
	req, _ := http.NewRequest(
		fiber.MethodPost,
		"/invert",
		body,
	)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Act
	res, err := test.TestRequest(req)

	// Assert
	bodyBytes, errRead := io.ReadAll(res.Body)
	if errRead != nil {
		log.Fatal(errRead)
	}
	bodyString := string(bodyBytes)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, res.StatusCode)
	assert.Equal(t, expectedBody, bodyString)
}

// Given a valid file request
// When Post to /multiply
// Then It retrieves 200 status-code and valid body response
func TestHttpMultiplySuccess(t *testing.T) {
	// Arrange
	expectedBody := "362880"
	appTest := fxtest.New(t, server.NewServer())
	appTest.RequireStart()
	defer appTest.RequireStop()

	body, writer, _ := createFileRequest("matrixtest.csv")

	uc := usecase.NewMultiplyUseCase()
	multiplyCtrl := httpApp.NewMultiplyController(uc)
	test := NewFiberTest(t, middleware.ValidationInput, middleware.ValidationMatrix, multiplyCtrl.Handler)
	req, _ := http.NewRequest(
		fiber.MethodPost,
		"/multiply",
		body,
	)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Act
	res, err := test.TestRequest(req)

	// Assert
	bodyBytes, errRead := io.ReadAll(res.Body)
	if errRead != nil {
		log.Fatal(errRead)
	}
	bodyString := string(bodyBytes)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, res.StatusCode)
	assert.Equal(t, expectedBody, bodyString)
}

// Given a not valid file request
// When Post to /multiply
// Then It retrieves 400 status-code
func TestHttpMultiplyFail(t *testing.T) {
	// Arrange
	appTest := fxtest.New(t, server.NewServer())
	appTest.RequireStart()
	defer appTest.RequireStop()

	body, writer, _ := createFileRequest("matrixinvalid.csv")

	uc := usecase.NewMultiplyUseCase()
	multiplyCtrl := httpApp.NewMultiplyController(uc)
	test := NewFiberTest(t, middleware.ValidationInput, middleware.ValidationMatrix, multiplyCtrl.Handler)
	req, _ := http.NewRequest(
		fiber.MethodPost,
		"/multiply",
		body,
	)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Act
	res, err := test.TestRequest(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}

// Given a valid file request
// When Post to /sum
// Then It retrieves 200 status-code and valid body response
func TestHttpSumSuccess(t *testing.T) {
	// Arrange
	expectedBody := "45"
	appTest := fxtest.New(t, server.NewServer())
	appTest.RequireStart()
	defer appTest.RequireStop()

	body, writer, _ := createFileRequest("matrixtest.csv")

	uc := usecase.NewSumUseCase()
	sumCtrl := httpApp.NewSumController(uc)
	test := NewFiberTest(t, middleware.ValidationInput, middleware.ValidationMatrix, sumCtrl.Handler)
	req, _ := http.NewRequest(
		fiber.MethodPost,
		"/sum",
		body,
	)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Act
	res, err := test.TestRequest(req)

	// Assert
	bodyBytes, errRead := io.ReadAll(res.Body)
	if errRead != nil {
		log.Fatal(errRead)
	}
	bodyString := string(bodyBytes)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, res.StatusCode)
	assert.Equal(t, expectedBody, bodyString)
}

// Given a valid not file request
// When Post to /sum
// Then It retrieves 400 status-code
func TestHttpSumFail(t *testing.T) {
	// Arrange
	appTest := fxtest.New(t, server.NewServer())
	appTest.RequireStart()
	defer appTest.RequireStop()

	body, writer, _ := createFileRequest("matrixinvalid.csv")

	uc := usecase.NewSumUseCase()
	sumCtrl := httpApp.NewSumController(uc)
	test := NewFiberTest(t, middleware.ValidationInput, middleware.ValidationMatrix, sumCtrl.Handler)
	req, _ := http.NewRequest(
		fiber.MethodPost,
		"/sum",
		body,
	)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Act
	res, err := test.TestRequest(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}
