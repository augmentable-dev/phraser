[![Build Status](https://cloud.drone.io/api/badges/augmentable-opensource/phraser/status.svg)](https://cloud.drone.io/augmentable-opensource/phraser)
[![Go Report Card](https://goreportcard.com/badge/github.com/augmentable-opensource/phraser)](https://goreportcard.com/report/github.com/augmentable-opensource/phraser)
[![Maintainability](https://api.codeclimate.com/v1/badges/7849d3b904e0249f6402/maintainability)](https://codeclimate.com/github/augmentable-opensource/phraser/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/7849d3b904e0249f6402/test_coverage)](https://codeclimate.com/github/augmentable-opensource/phraser/test_coverage)

## phraser

Open source phrase management microservice for your applications, delivered over gRPC. A centralized phrase database (key-value store) for your APIs, frontends, or anything else.

Hardcoding text (error messages, help text, feature explanations, etc.) into your source code is quick and dirty, but can be inefficient to maintain and especially difficult for non-developers (i.e. marketing) to modify. Use `phraser` to centrally manage as much (or as little) human-read text across your applications (frontends, API services, microservices, background jobs, etc).


### FAQs

_So it's a CMS?_

Nope, `phraser` is just a string store for your applications (frontend/backend) that doesn't make assumptions about _how_ those strings are used. You can utilize it for CMS-like functionality (though there's no support for rich text, just plaintext strings).

