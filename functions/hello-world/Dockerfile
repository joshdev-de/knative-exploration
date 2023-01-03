FROM gcr.io/distroless/static:nonroot
ARG COMMIT
ENV COMMIT=${COMMIT}

USER nonroot:nonroot

COPY app /bin/app

ENTRYPOINT [ "app" ]
