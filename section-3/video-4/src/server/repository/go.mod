module repository

go 1.15

require (
	github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-4/src/server/entities v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.9.0
)

replace github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-4/src/server/entities => ../entities
