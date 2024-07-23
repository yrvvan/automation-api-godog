const reporter = require('cucumber-html-reporter');

const options = {
    theme: 'bootstrap',
    jsonFile: 'reports/report.json',
    output: 'reports/report.html',
    reportSuiteAsScenarios: true,
    launchReport: true,
    metadata: {
        "App Version": "0.0.1",
        "Test Environment": "LOCAL",
        "Browser": "Chrome@latest",
        "Platform": "MacOs",
        "Parallel": "Scenarios",
        "Executed": "Locally"
    }
};

reporter.generate(options);
