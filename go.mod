module github.com/zchee/compute-metadata-server

go 1.20

replace cloud.google.com/go/compute/metadata => cloud.google.com/go/compute/metadata v0.2.3

require (
	cloud.google.com/go/iam v0.10.0
	github.com/goccy/go-json v0.10.0
	github.com/google/go-safeweb v0.0.0-20221125093303-48c35df3de5e
	github.com/google/safehtml v0.1.0
	github.com/klauspost/cpuid/v2 v2.2.3
	golang.org/x/net v0.7.0
	golang.org/x/oauth2 v0.5.0
	golang.org/x/sys v0.5.0
	google.golang.org/api v0.110.0
	google.golang.org/genproto v0.0.0-20230209215440-0dfe4f8abfcc
)

require (
	cloud.google.com/go/compute v1.18.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
