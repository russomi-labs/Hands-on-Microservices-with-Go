module main

go 1.15

require (
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-4/src/proto v0.0.0-00010101000000-000000000000
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-4/src/server/WTAServer v0.0.0-00010101000000-000000000000
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-4/src/server/repository v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.3
	github.com/lib/pq v1.9.0
	google.golang.org/grpc v1.34.0
)

replace (
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-4/src/proto => ../proto
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-4/src/server/WTAServer => ./WTAServer
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-4/src/server/entities => ./entities
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-4/src/server/repository => ./repository

)
