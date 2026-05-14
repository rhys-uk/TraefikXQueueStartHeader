# Traefik Plugin - X-Queue-Start Timestamp Injector
Fork of https://plugins.traefik.io/plugins/68b947762c6511271c0f0f45/timestamp-injector
- https://github.com/helpshift/timestamp-injector#

A Traefik middleware plugin that adds a timestamp header to **HTTP requests** before they reach the backend services.

It injects the current Unix time in **microseconds** (e.g., `1715691234567890`).

**Why this plugin?**
This plugin is specifically optimized for Application Performance Monitoring (APM) tools like New Relic, Datadog, etc. By injecting the timestamp at the ingress layer (Traefik), the APM agents can calculate the exact "Request Queue Time" (the time a request spent traveling from the ingress controller to the application).

NOTE: The default header name is `X-Queue-Start`, but it is fully configurable.

## Static Configuration

Enable experimental plugins and specify this plugin (traefik pulls from github, the path below does not currently exists but is present as an example):

```yaml
experimental:
  plugins:
    timestamp-injector:
      moduleName: github.com/rhys-uk/TraefikXQueueStartHeader
      version: v0.9.0
```

## Dynamic Configuration

Create the middleware:

```yaml
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
    name: mytimestamp-middleware
    namespace: my-namespace
spec:
    plugin:
        timestamp-injector: {}
        # OR override the default header (X-Queue-Start)
        #timestamp-injector:
        #  HeaderName: x-custom-start
```

## How It Works

This plugin injects a header into each incoming request with the current Unix time in microseconds. It uses `strconv.FormatInt` under the hood to ensure minimal overhead and maximum performance on every request.

`strconv.FormatInt` was chosen over the source project's original `fmt.Sprintf` as it is over twice as fast.
