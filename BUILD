CONFIG.setdefault("TEMPORAL_AWS_SDK_GENERATOR_VERSION", "0.0.3")

genrule(
    name = "temporal-aws-sdk-generator",
    srcs = [remote_file(
        name = "temporal-aws-sdk-generator",
        _tag = "download",
        url = f"https://github.com/temporalio/temporal-aws-sdk-generator/releases/download/v{CONFIG.TEMPORAL_AWS_SDK_GENERATOR_VERSION}/temporal-aws-sdk-generator_{CONFIG.HOSTOS}_{CONFIG.HOSTARCH}.tar.gz",
        out = f"temporal-aws-sdk-generator-{CONFIG.TEMPORAL_AWS_SDK_GENERATOR_VERSION}-{CONFIG.HOSTOS}-{CONFIG.HOSTARCH}.tar.gz",
    )],
    outs = ["temporal-aws-sdk-generator"],
    binary = True,
    cmd = f'"$TOOL" x "$SRCS" temporal-aws-sdk-generator -o tmp && mv $(find tmp -name "temporal-aws-sdk-generator") "$OUT"',
    tools = [CONFIG.JARCAT_TOOL],
)
