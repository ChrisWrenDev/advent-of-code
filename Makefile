# https://gist.github.com/prwhite/8168133
help: ## Show this help
	@ echo 'Usage: make <target>'
	@ echo
	@ echo 'Available targets:'
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

check-cookie:  ## ensures $AOC_SESSION_COOKIE env var is set
	@ test $${AOC_SESSION_COOKIE?env var not set}

prompt: check-cookie ## get prompt, requires $AOC_SESSION_COOKIE, optional: $DAY and $YEAR
	@ if [[ -n $$DAY && -n $$YEAR ]]; then \
		go run main.go -get-prompt=true -day $(DAY) -year $(YEAR) -cookie $(AOC_SESSION_COOKIE); \
	elif [[ -n $$DAY ]]; then \
		go run main.go -get-prompt=true -day $(DAY) -cookie $(AOC_SESSION_COOKIE); \
	else \
		go run main.go -get-prompt=true -cookie $(AOC_SESSION_COOKIE); \
	fi

input: check-cookie ## get input, requires $AOC_SESSION_COOKIE, optional: $DAY and $YEAR
	@ if [[ -n $$DAY && -n $$YEAR ]]; then \
		go run main.go -get-input=true -day $(DAY) -year $(YEAR) -cookie $(AOC_SESSION_COOKIE); \
	elif [[ -n $$DAY ]]; then \
		go run main.go -get-input=true -day $(DAY) -cookie $(AOC_SESSION_COOKIE); \
	else \
		go run main.go -get-input=true -cookie $(AOC_SESSION_COOKIE); \
	fi

template: ## create initial files for aoc, optional: $DAY and $YEAR
	@ if [[ -n $$DAY && -n $$YEAR ]]; then \
		go run main.go -generate-template=true -day $(DAY) -year $(YEAR) ; \
	elif [[ -n $$DAY ]]; then \
		go run main.go -generate-template=true -day $(DAY); \
	else \
		go run main.go -generate-template=true; \
	fi