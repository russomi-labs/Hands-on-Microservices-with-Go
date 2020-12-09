module main

go 1.15

require (
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-1/video-6/src/api/handlers v0.0.0-00010101000000-000000000000
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-1/video-6/src/api/repository v0.0.0-00010101000000-000000000000
	github.com/globalsign/mgo v0.0.0-20180424091348-efe0945164a7 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.6.2
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace (
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-1/video-6/src/api/handlers => ./handlers
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-1/video-6/src/api/repository => ./repository
)
