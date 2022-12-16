package handler

import (
	"net/http"

	"flyme-backend/app/interfaces/response"
	"flyme-backend/app/usecase"

	"github.com/labstack/echo/v4"
)

type ImageHandler struct {
	imageUseCase *usecase.ImageUseCase
}

func NewImageHandler(u *usecase.ImageUseCase) *ImageHandler {
	return &ImageHandler{u}
}

func (h *ImageHandler) UploadIcon(c echo.Context) error {
	userID := c.Param("user_id")

	file, err := c.FormFile("icon")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 50, Message: err.Error()})
	}

	f, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 50, Message: err.Error()})
	}
	defer f.Close()

	// if err := h.imageUseCase.ValidateImg(f); err != nil {
	// 	return c.JSON(http.StatusBadRequest, response.Error{Code: 40, Message: err.Error()})
	// }

	err = h.imageUseCase.UploadIconImg(f, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 50, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &response.ImageResponse{Message: "uploaded"})
}

func (h *ImageHandler) DownloadIcon(c echo.Context) error {
	userID := c.Param("user_id")

	res := c.Response()
	res.Header().Set("Cache-Control", "no-store")
	res.Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)
	res.Header().Set(echo.HeaderAccessControlExposeHeaders, "Content-Disposition")
	res.Header().Set(echo.HeaderContentDisposition, "attachment; filename="+userID+"_icon.png")
	res.WriteHeader(http.StatusOK)

	err := h.imageUseCase.DownloadIconImg(res.Writer, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 50, Message: err.Error()})
	}

	return c.NoContent(http.StatusOK)
}
