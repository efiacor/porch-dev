#  Copyright 2023 The Nephio Authors.
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

GO_VERSION ?= 1.22.2
TEST_COVERAGE_FILE=lcov.info
TEST_COVERAGE_TMP_FILE=lcov.tmp.info
TEST_COVERAGE_HTML_FILE=coverage_unit.html
TEST_COVERAGE_FUNC_FILE=func_coverage.out
GIT_ROOT_DIR ?= $(dir $(lastword $(MAKEFILE_LIST)))
include $(GIT_ROOT_DIR)/detect-container-runtime.mk

.PHONY: unit
unit: test

.PHONY: test
test: ## Run unit tests (go test)
ifeq ($(CONTAINER_RUNNABLE), 0)
		$(RUN_CONTAINER_COMMAND) docker.io/nephio/gotests:1782782171367346176 \
        sh -e -c "git config --global --add user.name test; \
	 	git config --global --add user.email test@nephio.org; \
	 	go test ./... -v -coverprofile ${TEST_COVERAGE_TMP_FILE}; \
		cat ${TEST_COVERAGE_TMP_FILE} | grep -v "test" > ${TEST_COVERAGE_FILE}; \
        go tool cover -html=${TEST_COVERAGE_FILE} -o ${TEST_COVERAGE_HTML_FILE}; \
        go tool cover -func=${TEST_COVERAGE_FILE} -o ${TEST_COVERAGE_FUNC_FILE}"
else
		go test ./... -v -coverprofile ${TEST_COVERAGE_FILE}
		go tool cover -html=${TEST_COVERAGE_FILE} -o ${TEST_COVERAGE_HTML_FILE}
		go tool cover -func=${TEST_COVERAGE_FILE} -o ${TEST_COVERAGE_FUNC_FILE}
endif

.PHONY: unit-clean
unit-clean: ## Clean up the artifacts created by the unit tests
ifeq ($(CONTAINER_RUNNABLE), 0)
		$(CONTAINER_RUNTIME) system prune -f
endif
		rm -f ${TEST_COVERAGE_FILE} ${TEST_COVERAGE_HTML_FILE} ${TEST_COVERAGE_FUNC_FILE} > /dev/null 2>&1
