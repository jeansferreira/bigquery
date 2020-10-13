FROM golang:1.14-stretch

ENV GOLANG_CI_LINT_VERSION=v1.18.0
ENV GIT_TERMINAL_PROMPT=1
ENV GO111MODULE=on
ENV GOPROXY=direct
ENV GOSUMDB=off

RUN cd /usr && \
    wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s ${GOLANG_CI_LINT_VERSION}

ARG USER
ARG USER_ID
ARG GROUP_ID

RUN groupadd -f -g ${GROUP_ID} ${USER} && \
    useradd -m -g ${GROUP_ID} -u ${USER_ID} ${USER} || echo "user already exists"

USER ${USER_ID}:${GROUP_ID}

WORKDIR /app
