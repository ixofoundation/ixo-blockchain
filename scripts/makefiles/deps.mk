###############################################################################
###                           Dependency Updates                            ###
###############################################################################
deps-help:
	@echo "Dependency Update subcommands"
	@echo ""
	@echo "Usage:"
	@echo "  make deps-[command]"
	@echo ""
	@echo "Available Commands:"
	@echo "  clean                    Remove artifacts"
	@echo "  distclean                Remove vendor directory"
	@echo "  draw                     Create a dependency graph"
	@echo "  go-mod-cache             Download go modules to local cache"
	@echo "  go.sum                   Ensure dependencies have not been modified"
	@echo "  tidy-workspace           Tidy workspace (no workspaces currently)"
deps: deps-help

deps-go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

deps-go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@GOWORK=off go mod verify
# TODO test to make sure @GOWORK=off works
# @go mod verify

deps-draw:
	@# requires brew install graphviz or apt-get install graphviz
	go get github.com/RobotsAndPencils/goviz
	@goviz -i ./cmd/ixod -d 2 | dot -Tpng -o dependency-graph.png
# TODO test or use below
# @goviz -i github.com/ixofoundation/ixo-blockchain/cmd/ixod -d 2 | dot -Tpng -o dependency-graph.png

deps-clean:
	rm -rf $(CURDIR)/artifacts/

deps-distclean: clean
	rm -rf vendor/

deps-tidy-workspace:
	@./scripts/tidy_workspace.sh
