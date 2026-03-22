const fs = require('fs');
const path = require('path');

function parseYamlFile(filePath) {
    try {
        const yamlData = fs.readFileSync(filePath, 'utf8');
        const yaml = require('js-yaml');
        const data = yaml.safeLoad(yamlData);
        return data;
    } catch (error) {
        console.error(`Error parsing YAML file: ${error}`);
        return null;
    }
}

function parseJsonFile(filePath) {
    try {
        const jsonData = fs.readFileSync(filePath, 'utf8');
        return JSON.parse(jsonData);
    } catch (error) {
        console.error(`Error parsing JSON file: ${error}`);
        return null;
    }
}

function getParser(filePath) {
    const fileExtension = path.extname(filePath);
    switch (fileExtension) {
        case '.yaml':
        case '.yml':
            return parseYamlFile(filePath);
        case '.json':
            return parseJsonFile(filePath);
        default:
            return null;
    }
}

module.exports = {
    parseYamlFile,
    parseJsonFile,
    getParser
};