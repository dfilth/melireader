package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"melireader/internal/core/port"
	"mime/multipart"
	"net/http"
	"strconv"
)

// FileItemHandler represents the HTTP handler for items file requests
type FileItemHandler struct {
	svc port.ItemService
}

// NewFileItemHandler creates a new FileItemHandler instance
func NewFileItemHandler(svc port.ItemService) *FileItemHandler {
	return &FileItemHandler{
		svc,
	}
}

func (fh *FileItemHandler) ItemFileRegister(ctx *gin.Context) {

	file, fileHeader, err := ctx.Request.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error al cerrar el archivo:", err)
		}
	}(file)

	ctx.JSON(http.StatusOK, gin.H{"message": "Archivo recibido y procesamiento iniciado"})

	go func() {
		_, err = fh.svc.Register(ctx, file, fileHeader)
		if err != nil {
			log.Println("Error al procesar el archivo y escribir en la base de datos:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el archivo y " +
				"escribir en la base de datos"})
			return
		}
	}()
}

func (fh *FileItemHandler) GetItemsPage(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	pageSizeStr := ctx.Query("pageSize")
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize parameter"})
		return
	}

	items, err := fh.svc.GetItemsPage(ctx, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"items": items})
}
