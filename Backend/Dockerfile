FROM public.ecr.aws/lambda/go:1.x

COPY main.go ${LAMBDA_TASK_ROOT}
RUN go build -o app main.go

CMD ["app"]