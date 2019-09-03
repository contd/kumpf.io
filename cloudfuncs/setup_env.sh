#!/usr/bin/env bash

export PROJECT_ID="tranquil-buffer-251815"

gcloud functions deploy api --entry-point Send --runtime go111 --trigger-http --set-env-vars PROJECT_ID=$PROJECT_ID

# Consumer trigger
gcloud functions deploy consumer --entry-point Receive --runtime go111 --trigger-topic=randomNumbers
gcloud functions logs read consumer

# Cleanup
gcloud functions delete api
gcloud functions delete consumer
gcloud pubsub topics delete randomNumbers
