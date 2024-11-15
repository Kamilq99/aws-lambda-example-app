FROM public.ecr.aws/lambda/go:1.2024.10.04.19 

COPY generate_pass.go ${LAMBDA_TASK_ROOT}

CMD [ "generate_pass" ]