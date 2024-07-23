# Define the output directory for the report
REPORT_DIR = ./reports
REPORT_FILE = $(REPORT_DIR)/report.json

# Ensure the report directory exists
$(REPORT_DIR):
	mkdir -p $(REPORT_DIR)

# Run godog with the specified format and output the report
tests: $(REPORT_DIR)
	godog --format=cucumber:$(REPORT_FILE)

# Clean the report directory
clean:
	rm -rf $(REPORT_DIR)

.PHONY: run-tests clean
