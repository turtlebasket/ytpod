# ytpod

`ytpod` is a YouTube channel RSS generator compatible with Apple Podcasts or other players like AntennaPod.

## Usage

There’s a [public instance](https://ytpod.fly.dev/) hosted on `fly.io` that can serve as an example. Basically, to obtain a feed of a YouTube source, one can just forward the YouTube path, for example, a feed from the author’s channel can be obtained via:
```
https://ytpod.fly.dev/youtube/feed/@lordawesomeguy
https://ytpod.fly.dev/youtube/feed/user/lordawesomeguy
https://ytpod.fly.dev/youtube/feed/channel/UCGFSAnRMBuCeEIiAAlKzdfg
```

This service fetches RSS feed from YouTube and enriches it with direct audio streams.
Keep in mind that the code quality is rather poor because it is nothing more than a hacky weekend project.

## Hosting

### Environment Variables

The `YTPOD_URL` environment variable must be set for the homepage, e.g. `https://ytpod.my-site.com`.

### Docker

[There's an image here](https://github.com/turtlebasket/ytpod/pkgs/container/ytpod). To run it:

```bash
# Minimal setup example
cat 'YTPOD_URL=https://ytpod.my-site.com' > .env
docker run -t -i -p 8080:8080 --env-file .env ghcr.io/turtlebasket/ytpod:latest
```

## License

Source code is available under the MIT [License](/LICENSE).
