# CSV Merger CLI

**CSV Merger** is a command-line tool for merging two CSV files based on specified key columns. It allows you to combine data from two CSV files where the rows match based on a key column.

## Installation

Ensure you have Go installed. Clone the repository and build the application:

```sh
$ git clone https://github.com/arvinpaundra/csv-merger.git
$ cd csv-merger
$ go build -o csv-merger
# or
$ make build
```

## Usage

The `csv-merger` command requires the following flags:

- `--source`: Path to the main CSV file.
- `--target`: Path to the CSV file to merge with the source.
- `--source-key`: Column name in the source file to match on.
- `--target-key`: Column name in the target file to match on.
- `--out`: Specify the output filename (default result.csv)

### Example

To merge two CSV files where `source.csv` and `target.csv` are matched based on the columns `foo` in the source file and `bar` in the target file:

```sh
$ ./csv-merger --source /path/to/source.csv \
               --target /path/to/target.csv \
               --source-key foo \
               --target-key bar
```

## Flags

- `--source`:
  - **Description**: Path to the main CSV file.
  - **Usage**: `--source /path/to/source.csv`
- `--target`:
  - **Description**: Path to the CSV file to merge with the source.
  - **Usage**: `--target /path/to/target.csv`
- `--source-key`:
  - **Description**: Column name in the source file to match on.
  - **Usage**: `--source-key column_name`
- `--target-key`:
  - **Description**: Column name in the target file to match on.
  - **Usage**: `--target-key column_name`
- `--out`:
  - **Description**: Specify the output filename (default result.csv)
  - **Usage**: `--out merged.csv
