from sentry_sdk import init, capture_message

init("https://mydsn@sentry.io/123")
capture_message("Hello World")
