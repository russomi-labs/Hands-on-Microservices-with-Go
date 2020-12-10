module main

go 1.15

require (
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-3/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.34.0
)

replace github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-3/proto => ../proto
