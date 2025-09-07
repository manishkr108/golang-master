👉 What is a struct in Go, and why would you use it instead of a simple variable or map?

A struct in Go is a composite data type that groups together fields, possibly of different types, under a single name. Unlike simple variables that hold only one value, or maps which are more dynamic but less type-safe, structs provide a strongly-typed way to model real-world entities. For example, a User struct can hold fields like ID int, Name string, and Active bool together in one logical unit.


Question 2 (Comparison with Classes in OOP)

👉 Since Go doesn’t have classes, how do structs + methods in Go compare to classes in other languages like Java or PHP?

In Go, we don’t have classes like in Java or PHP. Instead, we use structs to hold data (fields), and then we attach methods to structs to define behavior. This combination makes structs behave very similar to classes in OOP.

Struct = fields (like class properties)

Methods = functions attached to struct (like class methods)

Interfaces = define behavior contracts (like abstract classes)

type User struct {
    ID   int
    Name string
}

// Method attached to struct
func (u *User) UpdateName(newName string) {
    u.Name = newName
}

The difference from classes is that Go doesn’t support inheritance. Instead, Go uses composition (structs inside structs) and interfaces for polymorphism.


Question 3 (Polymorphism with Structs)

👉 Since Go doesn’t have classes and inheritance, how can we achieve polymorphism using structs + interfaces? Can you show me an example with a Notifier interface and a User struct?

🟢 What is Polymorphism?

In OOP languages like Java/PHP, polymorphism means one interface, many implementations.
For example: different classes (EmailNotifier, SMSNotifier) can implement the same interface (Notifier) but behave differently.

👉 In Go, we achieve this using interfaces + structs

🟢 Example in Go
Step 1: Define an Interface
type Notifier interface {
    Notify() // any struct with Notify() method becomes a Notifier
}

Step 2: Define Structs
type User struct {
    Name  string
    Email string
}

type Admin struct {
    Name  string
    Phone string
}

Step 3: Attach Methods
func (u User) Notify() {
    fmt.Println("Sending email to:", u.Email)
}

func (a Admin) Notify() {
    fmt.Println("Sending SMS to:", a.Phone)
}

Step 4: Write a Function that Accepts the Interface
func SendNotification(n Notifier) {
    n.Notify() // calls the correct method based on struct type
}

Step 5: Use in main
func main() {
    u := User{Name: "Manish", Email: "manish@example.com"}
    a := Admin{Name: "Raj", Phone: "9999999999"}

    SendNotification(u) // "Sending email to: manish@example.com"
    SendNotification(a) // "Sending SMS to: 9999999999"
}


🟢 Why is this powerful?

You don’t care which struct is passed (User or Admin), as long as it implements Notify().

This is polymorphism in Go: different structs can behave differently under the same interface.

How do you achieve polymorphism in Go?

Go achieves polymorphism through interfaces. A struct can implement an interface by simply defining the required methods, without explicitly declaring it. This allows different structs to be used interchangeably where the interface is expected.