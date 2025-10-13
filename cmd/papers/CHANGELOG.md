# Changelog

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Parsing text within points as Markdown.
- Parsing Bullet Paper types (Point, Talking, and Bullet Background)

### Changed
- YAML header needs to be terminated with the YAML "end of documents" symbol (...) to prevent
  misinterpreting bullet dashes later in the document.
- BREAKING CHANGE: YAML header key "type" changed to "paper_type"

### Deprecated

### Removed

### Fixed

### Security

## [0.1.1]

### Fixed
- Incorrectly interpreting LaTeX em-dashes (---) as YAML document separators

## [0.1.0]

### Added
- Parsing Outline PMD files

[unreleased]: https://github.com/DerHabicht/writing-tools/compare/papers-v0.1.1...HEAD
[0.1.1]: https://github.com/DerHabicht/writing-tools/compare/papers-v0.1.0...papers-v0.1.1
[0.1.0]: https://github.com/DerHabicht/writing-tools/releases/tag/papers-v0.1.0