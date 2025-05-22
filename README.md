# 🃏 Go Deck of Cards

This is a basic implementation of a deck of cards written in Go. The project was developed while learning the Go programming language and demonstrates several core concepts.

Use this project as a reference to understand:

- ✅ What is a slice and how to use it in Go  
- 📦 How to import standard and custom libraries  
- 🔁 Receiver functions (methods with receiver types)  
- 🔄 Returning multiple values from a function  
- 💾 File I/O: reading from and writing to files  
- 🎲 Generating pseudo-random numbers and setting custom seeds  

---

## 📂 File Structure

```bash
deck/
├── deck.go            # Core logic
├── deck_test.go       # Unit tests
├── main.go            # Example usage
├── cards.txt          # Example output file
```
---

## 📘 Features & Functions

### `deal`
Draws a hand of cards from the deck.

### `print`
Prints all cards in the deck to the console.

### `saveToFile` / `readFromFile`
Saves a deck (as a `[]byte`) to a file and reads it back.

### `shuffle`
Shuffles the cards in the deck using a pseudo-random number generator with a seed.

---

## 🧪 Unit Testing

Basic unit tests are included to cover:

- Deck creation  
- File I/O functionality  

Run the tests using:

```bash
go test



