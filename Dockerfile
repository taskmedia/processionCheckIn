FROM golang:alpine AS build

ARG RELEASE_COMMIT=""
ARG RELEASE_VERSION="dev"

WORKDIR /build
COPY . /build

RUN go build \
  -ldflags "-X github.com/taskmedia/processionCheckIn/pkg/version.VERSION=${RELEASE_VERSION}" \
  -ldflags "-X github.com/taskmedia/processionCheckIn/pkg/version.COMMIT=${RELEASE_COMMIT}" \
  -o bin/pci \
  cmd/backend/main.go

FROM alpine

RUN addgroup -S tm && \
  adduser -S tm -G tm

COPY --from=build /build/bin/pci /bin/pci

USER tm
EXPOSE 8080

ENTRYPOINT [ "/bin/pci" ]
