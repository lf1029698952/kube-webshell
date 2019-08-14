all: binary

.PHONY: binary
binary:
	@./scripts/build/binary.sh

.PHONY: cross
cross:
	@./scripts/build/cross.sh
	
clean:
	@-rm -rf build/sahaba*
	@-rm -rf vendor
	@-git worktree prune