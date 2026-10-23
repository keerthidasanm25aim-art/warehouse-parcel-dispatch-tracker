# Warehouse Parcel Dispatch Tracker

## About the Project

The Warehouse Parcel Dispatch Tracker is a simple command-line application developed using the Go programming language. It helps manage parcel registration, tracking, status updates, and deletion of delivered parcel records.

Parcel details are stored in a text file named `parcel.txt`.

## Features

- Register a new parcel.
- Store tracking number, weight, and destination postal code.
- Track a parcel using its tracking number.
- Update parcel status to `In-Transit` and `Delivered`.
- Save registered parcel details in a text file.
- Delete a parcel record after delivery.
- Handle errors such as invalid weight and missing parcels.

## Technologies Used

- **Programming Language:** Go (Golang)
- **File Handling:** Go `os` package
- **Input Conversion:** Go `strconv` package
- **Data Storage:** Go map and text file
- **Version Control:** Git
- **Code Hosting:** GitHub

## Project Structure

```text
go_program/
├── main.go
├── go.mod
├── parcel.txt
└── logistics/
    └── logistics.go
```

- `main.go` — Main program and user interaction.
- `go.mod` — Go module configuration.
- `logistics/logistics.go` — Parcel registration, tracking, status updates, and deletion.
- `parcel.txt` — Stores registered parcel details.

## Installation

### 1. Install Go

Download and install Go from the official website:

https://go.dev/dl/

Check your installation:

```bash
go version
```

### 2. Clone the Repository

```bash
git clone https://github.com/keerthidasanm25aim-art/warehouse-parcel-dispatch-tracker.git
```

### 3. Open the Project Folder

```bash
cd warehouse-parcel-dispatch-tracker
```

Make sure `main.go`, `go.mod`, and the `logistics` folder are in the project root.

## How to Run

### 1. Build the Project

```bash
go build ./...
```

### 2. Run the Program

The program accepts three command-line arguments:

1. Tracking number
2. Parcel weight in kilograms
3. Destination postal code

Example:

```bash
go run . TRK1001 4.5 641602
```

## Example Output

```text
Parcel Registered Successfully

Parcel Details
Tracking Number : TRK1001
Weight (KG)     : 4.50
Destination     : 641602
Status          : Received

After Update
Status : In-Transit

After Delivery
Status : Delivered

Parcel record deleted successfully.
```

*Note: The exact output depends on the program version and whether the parcel registration succeeds.*

## Sample Parcel Record

The registered parcel details are appended to `parcel.txt`.

```text
Tracking Number : TRK1001
Weight : 4.50 KG
Destination : 641602
Status : Received
--------------------------
```

## Parcel Status Flow

```text
Received
   |
   v
In-Transit
   |
   v
Delivered
   |
   v
Record Deleted
```

## Error Handling

The application handles common errors, including:

- Invalid weight input.
- Duplicate parcel registration.
- Parcel not found.
- Attempting to delete a parcel before delivery.
- File access errors.

## Limitations

- Parcel records are held in memory while the program runs.
- Registered parcel details are appended to `parcel.txt`.
- Existing parcels are not automatically loaded back into memory after restarting the application.
- Deleting a parcel from memory does not remove its previous entry from the text file.

## Future Improvements

- Add permanent parcel storage.
- Load saved parcel records when the program starts.
- Add a menu-driven interface.
- Validate postal codes and parcel weights.
- Add automated tests.

## Author

**KEERTHIDASAN**

## License

This project is intended for learning and educational purposes.
