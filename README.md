/*builder pattern*/
- product -> House
- Builder Interface
- Concrete Builder 1 -> Mormal House
- Concrete Builder 2 -> Igloo House
- Director

/*Abstract factory*/
- Abstract Product Interface 1 Shoe
- Abstract Product Interface 2 Short
- Concrete Product 1
- Concrete Product 2
- Abstract Factory Interface 1 Adidas
- Abstract Factory Interface 2 Nike
- Concrete Factory 1
- Concrete Factory 2

/*Prototype*/
- Prototype interface
  + Clone
- Concrete prototype 1 - File
- Concrete prototype 2 - Folder

/* Chain of responsibility*/
- Client -> patient
- Handler interface -> Department
  + Execute
  + SetNext
- Concrete handler 1 -> Reception
- Concrete handler 2 -> Doctor
- Concrete handler 3 -> Medicine
- Concrete handler 4 -> Cashier