# Mood Tracker

A simple command-line tool to track your daily moods with emoji support.

## Features

- Record your mood with emoji visualization
- View your mood timeline
- Simple and intuitive command-line interface
- SQLite database for persistent storage

## Installation

1. Make sure you have Go installed on your system
2. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/mood-tracker.git
   cd mood-tracker
   ```
3. Run the build script:
   ```bash
   chmod +x build.sh
   ./build.sh
   ```
4. Install the package in your machine:
   ```bash
   go install
   ```

## Usage

### Recording a Mood

```bash
mood-tracker record <mood>
```

Available moods:
- happy 😊
- sad 😔
- neutral 😐
- angry 😠
- excited 🤩

### Viewing Timeline

```bash
mood-tracker timeline
```

## Development

### Building from Source

```bash
go build
```

### Versioning

The project uses git tags for versioning. To create a new version:

```bash
git tag v1.0.0
git push origin v1.0.0
```

## License

MIT License 