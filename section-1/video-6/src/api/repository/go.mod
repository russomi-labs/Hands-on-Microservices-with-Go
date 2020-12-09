module repository

go 1.15

replace (
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-1/video-6/src/api/handlers => ./handlers
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-1/video-6/src/api/repository => ./repository
)

require (
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
