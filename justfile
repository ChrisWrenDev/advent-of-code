set shell := ["bash", "-eu", "-o", "pipefail", "-c"]
alias m := mise

default:
  @just --list

# -----------------------
# Scaffold AoC day
# -----------------------
aoc-new year day *langs:
  @m exec -- bash -eu -o pipefail -c '\
    root="aoc/{{year}}/{{day}}"; \
    mkdir -p "$$root"; \
    : > "$$root/input.txt"; \
    if [ {{#langs}} -eq 0 ]; then langs=(rust go zig); else langs=({{langs}}); fi; \
    for lang in "$${langs[@]}"; do \
      src="templates/aoc/$${lang}"; \
      dst="$$root/$${lang}"; \
      test -d "$$src" || { echo "Missing template: $$src" >&2; exit 1; }; \
      mkdir -p "$$dst"; \
      rsync -a --exclude ".gitkeep" "$$src/" "$$dst/"; \
      find "$$dst" -type f -print0 | while IFS= read -r -d "" f; do \
        perl -0777 -pe "s/\\{\\{YEAR\\}\\}/{{year}}/g; s/\\{\\{DAY\\}\\}/{{day}}/g" -i "$$f"; \
      done; \
      echo "Scaffolded $$lang -> $$dst"; \
    done'

# -----------------------
# Unified commands
# -----------------------
fmt lang path:
  @m exec -- bash -eu -o pipefail -c '\
    case "{{lang}}" in \
      rust) cd "{{path}}/rust" && cargo fmt ;; \
      go)   cd "{{path}}/go"   && gofmt -w . ;; \
      zig)  cd "{{path}}/zig"  && zig fmt . ;; \
      *) echo "Unknown lang: {{lang}}" >&2; exit 2 ;; \
    esac'

test lang path:
  @m exec -- bash -eu -o pipefail -c '\
    case "{{lang}}" in \
      rust) cd "{{path}}/rust" && cargo test ;; \
      go)   cd "{{path}}/go"   && go test ./... ;; \
      zig)  cd "{{path}}/zig"  && zig build test ;; \
      *) echo "Unknown lang: {{lang}}" >&2; exit 2 ;; \
    esac'

run lang path:
  @m exec -- bash -eu -o pipefail -c '\
    case "{{lang}}" in \
      rust) cd "{{path}}/rust" && cargo run --quiet ;; \
      go)   cd "{{path}}/go"   && go run ./... ;; \
      zig)  cd "{{path}}/zig"  && zig build run ;; \
      *) echo "Unknown lang: {{lang}}" >&2; exit 2 ;; \
    esac'

bench lang path *args:
  @m exec -- bash -eu -o pipefail -c '\
    case "{{lang}}" in \
      rust) \
        cd "{{path}}/rust" && \
        cargo bench {{args}} \
        ;; \
      go) \
        cd "{{path}}/go" && \
        go test -bench=. -benchmem {{args}} \
        ;; \
      zig) \
        cd "{{path}}/zig" && \
        zig build benchmark {{args}} \
        ;; \
      *) \
        echo "Unknown lang: {{lang}}" >&2; exit 2 ;; \
    esac'
