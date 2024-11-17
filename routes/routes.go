package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"backend-tugas-reactjs/controllers"
	"backend-tugas-reactjs/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB, r *gin.Engine) *gin.Engine {
	// cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"}

	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	r.Use(cors.New(corsConfig))

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	authMiddlewareRoute := r.Group("/auth")
	authMiddlewareRoute.Use(middlewares.JwtAuthMiddleware("all-user"))
	authMiddlewareRoute.GET("/me", controllers.GetCurrentUser)

	booksMiddlewareRoute := r.Group("/books")
	booksMiddlewareRoute.Use(middlewares.JwtAuthMiddleware("admin"))
	booksMiddlewareRoute.POST("", controllers.CreateBook)
	booksMiddlewareRoute.PATCH("/:id", controllers.UpdateBook)
	booksMiddlewareRoute.DELETE("/:id", controllers.DeleteBook)
	r.GET("/books", controllers.GetAllBook)
	r.GET("/books/:id", controllers.GetBookById)

	mahasiswaMiddlewareRoute := r.Group("/mahasiswa")
	mahasiswaMiddlewareRoute.Use(middlewares.JwtAuthMiddleware("admin"))
	mahasiswaMiddlewareRoute.POST("", controllers.CreateMahasiswa)
	mahasiswaMiddlewareRoute.PATCH("/:id", controllers.UpdateMahasiswa)
	mahasiswaMiddlewareRoute.DELETE("/:id", controllers.DeleteMahasiswa)
	r.GET("/mahasiswa", controllers.GetAllMahasiswa)
	r.GET("/mahasiswa/:id", controllers.GetMahasiswaById)

	mataKuliahMiddlewareRoute := r.Group("/mata-kuliah")
	mataKuliahMiddlewareRoute.Use(middlewares.JwtAuthMiddleware("admin"))
	mataKuliahMiddlewareRoute.POST("", controllers.CreateMataKuliah)
	mataKuliahMiddlewareRoute.PATCH("/:id", controllers.UpdateMataKuliah)
	mataKuliahMiddlewareRoute.DELETE("/:id", controllers.DeleteMataKuliah)
	r.GET("/mata-kuliah", controllers.GetAllMataKuliah)
	r.GET("/mata-kuliah/:id", controllers.GetMataKuliahById)

	dosenMiddlewareRoute := r.Group("/dosen")
	dosenMiddlewareRoute.Use(middlewares.JwtAuthMiddleware("admin"))
	dosenMiddlewareRoute.POST("/", controllers.CreateDosen)
	dosenMiddlewareRoute.PATCH("/:id", controllers.UpdateDosen)
	dosenMiddlewareRoute.DELETE("/:id", controllers.DeleteDosen)
	r.GET("/dosen", controllers.GetAllDosen)
	r.GET("/dosen/:id", controllers.GetDosenById)

	nilaiMiddlewareRoute := r.Group("/nilai")
	nilaiMiddlewareRoute.Use(middlewares.JwtAuthMiddleware("admin"))
	nilaiMiddlewareRoute.POST("/", controllers.CreateNilai)
	nilaiMiddlewareRoute.PATCH("/:id", controllers.UpdateNilai)
	nilaiMiddlewareRoute.DELETE("/:id", controllers.DeleteNilai)
	r.GET("/nilai", controllers.GetAllNilai)
	r.GET("/nilai/:id", controllers.GetNilaiById)

	jadwalKuliahMiddlewareRoute := r.Group("/jadwal-kuliah")
	jadwalKuliahMiddlewareRoute.Use(middlewares.JwtAuthMiddleware("admin"))
	jadwalKuliahMiddlewareRoute.POST("", controllers.CreateJadwalKuliah)
	jadwalKuliahMiddlewareRoute.PATCH("/:id", controllers.UpdateJadwalKuliah)
	jadwalKuliahMiddlewareRoute.DELETE("/:id", controllers.DeleteJadwalKuliah)
	r.GET("/jadwal-kuliah/:id", controllers.GetJadwalKuliahById)
	r.GET("/jadwal-kuliah", controllers.GetAllJadwalKuliah)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
