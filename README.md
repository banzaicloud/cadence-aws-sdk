# Cadence AWS SDK

**Cadence backport for [Temporal AWS SDK](https://github.com/temporalio/temporal-aws-sdk-go).**


## Development

Most of the code inside this library is generated. If you need to change the generated code,
look into the [templates](templates) directory and modify the templates there.

After making the changes, regenerate the code with the following command:

```bash
./pleasew generate
```

**Important: template changes and code generation MUST be end up in separate commits**

```bash
git add templates
git commit -m 'feat: I changed X to behave Y'
git add clients activities
git commit -m 'chore: regenerate code'
```


## License

Apache 2.0 License. Please see [License File](LICENSE) for more information.
