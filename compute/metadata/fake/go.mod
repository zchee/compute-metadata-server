module github.com/zchee/gce-metadata-server/compute/metadata/fake

go 1.18

replace github.com/zchee/gce-metadata-server/compute => ../../

require (
	cloud.google.com/go/iam v0.3.0
	github.com/google/go-safeweb v0.0.0-20220125171915-eb79df54b8bc
	github.com/google/safehtml v0.0.2
	github.com/zchee/gce-metadata-server/compute v0.0.0-20220614120728-0750052a9ff2
	golang.org/x/net v0.0.0-20220607020251-c690dde0001d
	golang.org/x/oauth2 v0.0.0-20220608161450-d0670ef3b1eb
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d
	google.golang.org/api v0.84.0
	google.golang.org/genproto v0.0.0-20220616135557-88e70c0c3a90
)

require (
	cloud.google.com/go/compute v1.7.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.0.0-20220520183353-fd19c99a87aa // indirect
	github.com/googleapis/gax-go/v2 v2.4.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
