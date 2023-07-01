package controller

import (
	"lbc/fizzbuzz/model"
	"lbc/fizzbuzz/stats"
	"lbc/fizzbuzz/usecase"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Constant storing paths of services
const (
	POSTFizzbuzzURI = "/fizzbuzz"
	GETStatsURI     = "/stats"
)

// Callback for POST /fizzbuzz endpoint
func postFizzbuzz(ctx *gin.Context) {
	var fizzBuzzRequest model.FizzbuzzRequest

	if err := ctx.BindJSON(&fizzBuzzRequest); err != nil {
		log.Errorf("Bad Request received \n %v", fizzBuzzRequest)
		log.Debug(err)
		return
	}

	stats.GetInstance().AddFor(fizzBuzzRequest)

	if err := model.ValidateFizzbuzzRequest(fizzBuzzRequest); err != nil {
		log.Errorf("Bad Request received \n %v", fizzBuzzRequest)
		log.Debug(err)
		ctx.JSON(400, nil)
		return
	}

	fizzbuzzed := usecase.Fizzbuzz(fizzBuzzRequest.FirstInteger, fizzBuzzRequest.SecondInteger, fizzBuzzRequest.Limit, fizzBuzzRequest.FirstString, fizzBuzzRequest.SecondString)
	ctx.JSON(200, fizzbuzzed)
}

// Callback for GET /stats endpoint
func getStats(ctx *gin.Context) {
	request, count, err := usecase.GetMostFrequentRequest(stats.GetInstance())

	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	var statsReply = &model.StatsReply{
		Request: request,
		Count:   count,
	}

	ctx.JSON(200, statsReply)
}

// Register in the router services
func RegisterEndpoints(router *gin.Engine) {
	log.Debugln("Registering " + POSTFizzbuzzURI)
	router.POST(POSTFizzbuzzURI, postFizzbuzz)

	log.Debugln("Registering " + GETStatsURI)
	router.GET(GETStatsURI, getStats)

}
