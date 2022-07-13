#!/usr/bin/env bash
ci:
	git add . && (git commit || echo ) && git pull
