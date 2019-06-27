<img src='https://s3-us-west-2.amazonaws.com/cortex-public/logo.png' height='88'>

<br>

**Get started:** [Install](https://docs.cortex.dev/install) • [Tutorial](https://docs.cortex.dev/tutorial) • <!-- CORTEX_VERSION_MINOR_STABLE e.g. https://docs.cortex.dev/v/0.2/ -->[Docs](https://docs.cortex.dev/v/0.5/) • <!-- CORTEX_VERSION_MINOR_STABLE -->[Examples](https://github.com/cortexlabs/cortex/tree/0.5/examples)

**Learn more:** [Website](https://cortex.dev) • [Blog](https://blog.cortex.dev) • [Subscribe](https://cortexlabs.us20.list-manage.com/subscribe?u=a1987373ab814f20961fd90b4&id=ae83491e1c) • [Twitter](https://twitter.com/cortex_deploy) • [Contact](mailto:hello@cortex.dev)

<br>

Cortex deploys your machine learning models to your cloud infrastructure. You define your deployment with simple declarative configuration, Cortex containerizes your models, deploys them as scalable JSON APIs, and manages their lifecycle in production.

Cortex is actively maintained by Cortex Labs. We're a venture-backed team of infrastructure engineers and [we're hiring](https://angel.co/cortex-labs-inc/jobs).

<br>

## How it works

**Define** your deployment using declarative configuration:

```yaml
# cortex.yaml

- kind: api
  name: my-api
  model: s3://my-bucket/my-model.zip
  request_handler: handler.py
  compute:
    replicas: 4
    gpu: 2
```

**Deploy** to your cloud infrastructure:

```
$ cortex deploy

Deploying ...
Ready! https://amazonaws.com/my-api
```

**Serve** real time predictions via scalable JSON APIs:

```
$ curl -d '{"a": 1, "b": 2, "c": 3}' https://amazonaws.com/my-api

{ prediction: "def" }
```

<br>

## Key features

- **Machine learning deployments as code:** Cortex deployments are defined using declarative configuration.

- **Multi framework support (coming soon):** Cortex will support TensorFlow, Keras, PyTorch, Scikit-learn, XGBoost, and more.

- **CPU / GPU support:** Cortex can run inference on CPU or GPU infrastructure.

- **Scalability:** Cortex can scale APIs to handle production workloads.

- **Rolling updates:** Cortex updates deployed APIs without any downtime.

- **Cloud native:** Cortex can be deployed on any AWS account in minutes.
