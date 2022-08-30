Go := go

.PHONY: generate
generate:
	@echo "  >  "$@"ing $(DESTDIR)"
	${MAKE} -C api generate