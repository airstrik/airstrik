package route

import "github.com/go-chi/chi"

func PublicRoute(r chi.Router) {
	r.Post("/signup", CreateAccount)
}

func AccountRoute(r chi.Router) {
	r.Post("/signup", CreateAccount)
	r.Get("/", GetAccount)
	//r.Get("/", GetAccount)
	//r.With(paginate).Get("/", listArticles)                           // GET /articles
	//r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017
	//
	//r.Post("/", createArticle)                                        // POST /articles
	//r.Get("/search", searchArticles)                                  // GET /articles/search
	//
	//// Regexp url parameters:
	//r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug)                // GET /articles/home-is-toronto
	//
	//// Subrouters:
	//r.Route("/{articleID}", func(r chi.Router) {
	//	r.Use(ArticleCtx)
	//	r.Get("/", getArticle)                                          // GET /articles/123
	//	r.Put("/", updateArticle)                                       // PUT /articles/123
	//	r.Delete("/", deleteArticle)                                    // DELETE /articles/123
	//})
}

// AdminRouter : will redirect all the Account Based service route file
func AdminRoute(r chi.Router) {
}

// SignInRoute : will redirect all the sign in Based service route file
func SignInRoute(r chi.Router) {
	//Sign up Url for the Application (New users)
	r.Post("/signup/mobile", SignUpMobile)
	r.Post("/signup/email", SignUpMobile)

	r.Post("/", SignUserNamePassword)
	r.Post("/signup", SignUserNamePassword)
	r.Get("/google", SignGoogle)
	r.Post("/google/callback", CallbackGoogle)
}

// UserRoute : will redirect all the user Based service route file
func UserRoute(r chi.Router) {
	r.Post("/", CreateUser)
	r.Get("/", GetUserList)
	r.Get("/{UserId}", GetUser)
	r.Put("/{UserId}", UpdateUser)
	r.Delete("/{UserId}", DeleteUser)
}
