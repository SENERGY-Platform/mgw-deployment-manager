module github.com/SENERGY-Platform/mgw-container-engine-wrapper

go 1.23

require (
	github.com/SENERGY-Platform/gin-middleware v0.4.4
	github.com/SENERGY-Platform/go-cc-job-handler v0.1.2
	github.com/SENERGY-Platform/go-service-base/config-hdl v0.1.1
	github.com/SENERGY-Platform/go-service-base/job-hdl v1.1.1
	github.com/SENERGY-Platform/go-service-base/job-hdl/lib v0.1.0
	github.com/SENERGY-Platform/go-service-base/logger v0.2.0
	github.com/SENERGY-Platform/go-service-base/srv-info-hdl v0.0.3
	github.com/SENERGY-Platform/go-service-base/srv-info-hdl/lib v0.0.2
	github.com/SENERGY-Platform/go-service-base/util v1.1.0
	github.com/SENERGY-Platform/go-service-base/watchdog v0.4.2
	github.com/SENERGY-Platform/mgw-container-engine-wrapper/lib v0.0.0-00010101000000-000000000000
	github.com/docker/docker v27.4.1+incompatible
	github.com/docker/go-connections v0.5.0
	github.com/gin-contrib/requestid v1.0.4
	github.com/gin-gonic/gin v1.10.0
	github.com/y-du/go-env-loader v0.5.2
	github.com/y-du/go-log-level v1.0.0
)

require (
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/bytedance/sonic v1.12.6 // indirect
	github.com/bytedance/sonic/loader v0.2.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/gin-contrib/sse v1.0.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.23.0 // indirect
	github.com/goccy/go-json v0.10.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.58.0 // indirect
	go.opentelemetry.io/otel v1.33.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.33.0 // indirect
	go.opentelemetry.io/otel/sdk v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.33.0 // indirect
	golang.org/x/arch v0.13.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/protobuf v1.36.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.5.1 // indirect
)

replace github.com/SENERGY-Platform/mgw-container-engine-wrapper/lib => ./lib
