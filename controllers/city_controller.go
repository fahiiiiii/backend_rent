// controllers/city_controller.go
package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "backend_rental/service"
    "backend_rental/utils"
)

// CityController operations for Cities
// @Title Cities API
// @Description City management with rate limiting
type CityController struct {
    web.Controller
    RateLimiter *utils.RateLimiter
    Service     *services.BookingService
}

// @Title Search Cities
// @Description Search cities based on query string
// @Param query query string true "Search query"
// @Success 200 {array} models.City
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 429 {object} controllers.ErrorResponse
// @Failure 500 {object} controllers.ErrorResponse
// @router /search [get]
func (c *CityController) Get() {
    // Apply rate limiting
    limiter := c.RateLimiter.GetLimiter(c.Ctx.Input.IP())
    if !limiter.Allow() {
        c.Ctx.Output.SetStatus(429)
        c.Data["json"] = ErrorResponse{Error: "Too many requests"}
        c.ServeJSON()
        return
    }

    query := c.GetString("query")
    if query == "" {
        c.Ctx.Output.SetStatus(400)
        c.Data["json"] = ErrorResponse{Error: "Query parameter is required"}
        c.ServeJSON()
        return
    }

    cities, err := c.Service.FetchCities(query)
    if err != nil {
        c.Ctx.Output.SetStatus(500)
        c.Data["json"] = ErrorResponse{Error: err.Error()}
        c.ServeJSON()
        return
    }

    c.Data["json"] = cities
    c.ServeJSON()
}

// @Title Get City by ID
// @Description Get city information by city ID
// @Param cityId path string true "City ID"
// @Success 200 {object} models.City
// @Failure 404 {object} controllers.ErrorResponse
// @Failure 429 {object} controllers.ErrorResponse
// @Failure 500 {object} controllers.ErrorResponse
// @router /byId/:cityId [get]
func (c *CityController) GetById() {
    // Apply rate limiting
    limiter := c.RateLimiter.GetLimiter(c.Ctx.Input.IP())
    if !limiter.Allow() {
        c.Ctx.Output.SetStatus(429)
        c.Data["json"] = ErrorResponse{Error: "Too many requests"}
        c.ServeJSON()
        return
    }

    // cityId := c.Ctx.Input.Param(":cityId")
    // Implement the logic to fetch city by ID
    // This would need to be added to your service layer
}

// ErrorResponse represents an error response
type ErrorResponse struct {
    Error string `json:"error"`
}