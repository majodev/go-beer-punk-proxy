package beers

// func GetBeersRoute(s *api.Server) *echo.Route {

// 	url, _ := url.Parse("https://api.punkapi.com/v2/beers")

// 	proxy := httputil.NewSingleHostReverseProxy(url)

// 	return s.Router.APIV1Beers.GET("", echo.WrapHandler(proxy))

// 	// return s.Router.APIV1Beers.GET("/", getBeersHandler(s))
// }

// func GetBeersRoute(s *api.Server) *echo.Route {

// 	url, _ := url.Parse("https://api.punkapi.com/v2/beers")

// 	targets := []*echoMiddleware.ProxyTarget{
// 	{
// 		URL: url,
// 	},
// 	}

// 	return s.Router.APIV1Beers.Use(echoMiddleware.Proxy(echoMiddleware.NewRoundRobinBalancer(targets)))

// 	// return s.Router.APIV1Beers.GET("/", getBeersHandler(s))
// }

// func GetBeerRoute(s *api.Server) *echo.Route {
// 	return s.Router.APIV1Beers.GET("/:id", getBeersHandler(s))
// }

// func GetRandomBeerRoute(s *api.Server) *echo.Route {
// 	return s.Router.APIV1Beers.GET("/random", getBeersHandler(s))
// }

// func getBeersHandler(s *api.Server) echo.HandlerFunc {

// 	url, _ := url.Parse("https://api.punkapi.com/v2/beers")

// 	proxy := httputil.NewSingleHostReverseProxy(url)

// 	// return s.Router.APIV1Beers.GET("", echo.WrapHandler(proxy))

// 	return func(c echo.Context) error {

// 		return echo.WrapHandler(proxy)

// 		// ctx := c.Request().Context()
// 		// params := beers.NewGetBeersRouteParams()

// 		// err := util.BindAndValidate(c, &params)
// 		// if err != nil {
// 		// 	return err
// 		// }

// 		// log := util.LogFromContext(ctx).With().Logger()

// 	}
// }
