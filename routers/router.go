// routers/router.go
// @APIVersion 1.0.0
// @Title Booking Cities API
// @Description API for fetching city information from Booking.com with rate limiting
// @Contact yourname@example.com
// @TermsOfServiceUrl http://example.com/terms/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
    beego "github.com/beego/beego/v2/server/web"
    "backend_rental/controllers"
    "backend_rental/service"
    "backend_rental/utils"
)

func init() {
    // Initialize dependencies
    bookingService := &services.BookingService{}
    rateLimiter := utils.NewRateLimiter()
    
    // Create the city controller with dependencies
    cityController := &controllers.CityController{
        Controller:  beego.Controller{},
        RateLimiter: rateLimiter,
        Service:     bookingService,
    }

    // Create API namespace
    ns := beego.NewNamespace("/v1",
        beego.NSNamespace("/cities",
            beego.NSInclude(
                cityController,
            ),
            // Define specific routes within the cities namespace
            beego.NSRouter("/search", cityController, "get:Get"),
            beego.NSRouter("/byId/:cityId", cityController, "get:GetById"),
            // You can add more routes here as needed
        ),
        // You can add more namespaces here for other features
    )

    // Register the namespace
    beego.AddNamespace(ns)

    // Optional: Add custom filters/middleware for the entire API
    beego.InsertFilter("/v1/*", beego.BeforeRouter, func(ctx *beego.Context) {
        // Add CORS headers or other middleware logic here
        ctx.Output.Header("Access-Control-Allow-Origin", "*")
    })
}