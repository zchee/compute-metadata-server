module github.com/zchee/compute-metadata-server/fakemetadata

go 1.18

replace cloud.google.com/go/compute => cloud.google.com/go/compute v1.8.0

require (
	github.com/goccy/go-json v0.9.10
	github.com/google/go-safeweb v0.0.0-20220125171915-eb79df54b8bc
	github.com/google/safehtml v0.0.3-0.20220430015336-00016cfeca15
	github.com/klauspost/cpuid/v2 v2.1.1-0.20220725114759-b27ab7bf7451
	golang.org/x/net v0.0.0-20220811182439-13a9a731de15
	golang.org/x/oauth2 v0.0.0-20220808172628-8227340efae7
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab
	google.golang.org/api v0.92.0
)

require (
	cloud.google.com/go/compute v1.8.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.1.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220804142021-4e6b2dfa6612 // indirect
	google.golang.org/grpc v1.48.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
