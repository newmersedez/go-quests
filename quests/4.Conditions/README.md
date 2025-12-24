# Go Conditionals Quest

Your task is to implement a single function using **`if`, `else if`, and `else`** in Go.

This quest is designed to test **real conditional reasoning**, not just syntax.

---

## Reference

- [https://gobyexample.com/if-else](https://gobyexample.com/if-else)
- [https://gobyexample.com/booleans](https://gobyexample.com/booleans)

---

## Function: `ClassifyRequest`

```go
func ClassifyRequest(age int, hasID bool, balance float64, isVIP bool) string
```

---

## Description

You are building a **request classification engine**.

Given user attributes (`age`, `hasID`, `balance`, `isVIP`), determine the **access level** based on a strict set of rules.

Exactly **one** classification string must be returned.

---

## Valid Return Values

The function must return **one of the following strings**:

- `"INVALID"`
- `"REJECTED"`
- `"VIP_ACCESS"`
- `"STANDARD_ACCESS"`
- `"LIMITED_ACCESS"`

No other values are allowed.

---

## Business Rules (Order Matters)

Rules must be evaluated **top to bottom**.
Once a condition matches, return immediately.

---

### 1. Invalid Request

Return `"INVALID"` if **any** of the following are true:

- `age <= 0`
- `balance < 0`

---

### 2. Rejected

Return `"REJECTED"` if **either** of the following are true:

- `age < 18`
- `hasID == false`

---

### 3. VIP Access

Return `"VIP_ACCESS"` if **both** are true:

- `isVIP == true`
- `balance >= 10000`

---

### 4. Standard Access

Return `"STANDARD_ACCESS"` if:

- `balance >= 1000`

---

### 5. Limited Access

All remaining **valid** requests fall into this category.

Return `"LIMITED_ACCESS"`.

---

## Requirements

1. Use `if / else if / else`
2. Use **compound conditions** (`&&`, `||`)
3. Use **early returns**
4. Do **not** use `switch`
5. Do **not** reorder the rules
6. Do **not** nest `if` blocks more than one level deep
7. Do not print anything
8. The function must be deterministic and side-effect free

---

## Examples

| age | hasID | balance | isVIP | Output          |
| --: | :---: | ------: | :---: | --------------- |
|  -1 | true  |     500 | false | INVALID         |
|  16 | true  |    2000 | false | REJECTED        |
|  25 | false |    5000 | true  | REJECTED        |
|  30 | true  |   15000 | true  | VIP_ACCESS      |
|  40 | true  |    3000 | false | STANDARD_ACCESS |
|  22 | true  |     200 | false | LIMITED_ACCESS  |

---

## Common Mistakes (Tests Will Catch These)

- Using multiple independent `if` statements
- Checking VIP before validating age/ID
- Forgetting early returns
- Allowing invalid inputs to pass through
- Returning default values incorrectly

---

## Key Ideas

- Condition **ordering is logic**
- `else if` enforces exclusivity
- Short-circuit logic (`&&`, `||`) matters
- Real-world validation is sequential, not parallel

---

## Run Tests

```bash
go test ./conditionals/solution
```

Do not add logging or printing.
Only return the correct classification string.
