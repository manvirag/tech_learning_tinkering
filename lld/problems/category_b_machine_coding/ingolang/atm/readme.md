Requirements
- The ATM system should support basic operations such as balance inquiry, cash withdrawal, and cash deposit. -> done
- Users should be able to authenticate themselves using a card and a PIN (Personal Identification Number). -> done
- The system should interact with a bank's backend system to validate user accounts and perform transactions. -> done
- The ATM should have a cash dispenser to dispense cash to users. -> done
- The system should handle concurrent access and ensure data consistency. 
- The ATM should have a user-friendly interface for users to interact with.



- Rough:
- balance inquiry , cash withdrawl , cash deposit,
- Assume atm have infinite money. 


- models
  - AtmState
    - insert
    - validate
    - Options
    - Balance inenquir -> insert
    - cash deposit -> amount -> insert
    - cash withdrawl -> amount -> dispense -> insert
  - BankRepo 

- usecase
  - Atm -> state