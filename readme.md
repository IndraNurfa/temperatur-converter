# Temperature Conversion Program

This Go program provides a command-line interface for converting temperatures between three units: Celsius, Fahrenheit, and Kelvin. It supports bidirectional conversion between all three units.

## Features

- Convert temperatures from:
  - Celsius to Fahrenheit and Kelvin
  - Fahrenheit to Celsius and Kelvin
  - Kelvin to Celsius and Fahrenheit
- User-friendly prompts to select the desired input and output temperature units.
- Validates user input for both unit selection and temperature values.

## How It Works

The program is implemented using a combination of structs and an interface:

### Structs

1. **`celsius`**: Represents temperatures in Celsius.
2. **`farenheit`**: Represents temperatures in Fahrenheit.
3. **`kelvin`**: Represents temperatures in Kelvin.

Each struct has methods to convert the temperature to the other two units.

### Interface

- **`hitungSuhu`**: An interface implemented by the structs, which defines the following methods:
  - `toCelsius() float64`
  - `toFarenheit() float64`
  - `toKelvin() float64`

### Conversion Logic

Each struct implements the `hitungSuhu` interface by defining the appropriate conversion formulas.

## Usage

1. Compile and run the program.
2. Follow the prompts to:
   - Select the initial temperature unit.
   - Select the target temperature unit.
   - Enter the temperature value.
3. The program will display the converted temperature value.

### Example

```bash
Masukkan opsi suhu awal:
1. Celsius
2. Farenheit
3. Kelvin
> 1

Masukkan opsi suhu akhir:
1. Celsius
2. Farenheit
3. Kelvin
> 2

Masukkan suhu:
> 25

Suhu yang dikonversi: 77.00
```

## Code Overview

### Struct Definitions

```go
type celsius struct {
    suhu float64
}

type farenheit struct {
    suhu float64
}

type kelvin struct {
    suhu float64
}
```

### Interface Definition

```go
type hitungSuhu interface {
    toCelsius() float64
    toFarenheit() float64
    toKelvin() float64
}
```

### Methods

#### Celsius Methods

- `toCelsius()`: Returns the same value.
- `toFarenheit()`: Converts Celsius to Fahrenheit using the formula:

  ```
  (9.0 / 5.0) * suhu + 32
  ```

- `toKelvin()`: Converts Celsius to Kelvin using the formula:

  ```
  suhu + 273.15
  ```

#### Fahrenheit Methods

- `toCelsius()`: Converts Fahrenheit to Celsius using the formula:

  ```
  (suhu - 32) * (5.0 / 9.0)
  ```

- `toFarenheit()`: Returns the same value.
- `toKelvin()`: Converts Fahrenheit to Kelvin using the formula:

  ```
  (suhu + 459.67) * (5.0 / 9.0)
  ```

#### Kelvin Methods

- `toCelsius()`: Converts Kelvin to Celsius using the formula:

  ```
  suhu - 273.15
  ```

- `toFarenheit()`: Converts Kelvin to Fahrenheit using the formula:

  ```
  (suhu * (9.0 / 5.0)) - 459.67
  ```

- `toKelvin()`: Returns the same value.

### Main Function

The `main` function:

1. Prompts the user to select input and output temperature units.
2. Validates the input.
3. Prompts the user to enter the temperature value.
4. Converts the temperature using the appropriate struct and methods.
5. Displays the result.

## Notes

- **Input Validation**: Ensures the user enters valid options and numeric values.
- **Floating-Point Precision**: Conversion results are formatted to two decimal places.

## How to Run

1. Save the code to a file, e.g., `temperature_converter.go`.
2. Open a terminal and navigate to the file's directory.
3. Run the program using:

   ```bash
   go run temperature_converter.go
   ```

## License

This program is provided as-is for educational purposes. Feel free to modify and use it for personal or commercial projects.

---

Feel free to report any issues or suggest improvements!
