#!/usr/bin/env bash
ci:
	git add . && (git commit || echo )
push:
	git pull && git push
