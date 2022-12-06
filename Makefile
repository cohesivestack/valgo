gen:


# ----
## LINTER stuff start

linter_include_check:
	@[ -f linter.mk ] && echo "linter.mk include exists" || (echo "getting linter.mk from github.com" && curl -sO https://raw.githubusercontent.com/spacetab-io/makefiles/master/golang/linter.mk)

.PHONY: lint
lint: linter_include_check
	@make -f linter.mk go_lint

## LINTER stuff end
# ----

# ----
## TESTS stuff start

tests_include_check:
	@[ -f tests.mk ] && echo "tests.mk include exists" || (echo "getting tests.mk from github.com" && curl -sO https://raw.githubusercontent.com/spacetab-io/makefiles/master/golang/tests.mk)

tests: tests_include_check
	@make -f tests.mk go_tests
.PHONY: tests

tests_html: tests_include_check
	@make -f tests.mk go_tests_html
.PHONY: tests_html

## TESTS stuff end
# ----
