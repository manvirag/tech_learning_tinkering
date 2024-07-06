package models

import (
	"fmt"
)

type Atm struct{
   AtmState AtmStateInterface
   BankRepo *BankRepo
   CurrentAccount string
}
func NewAtm(bankRepo *BankRepo) *Atm{
  return &Atm{
    BankRepo: bankRepo,
    AtmState: &InsertState{},
  }
}

func(atm *Atm) Process() {
  atm.AtmState.Process(atm)
}

type AtmStateInterface interface{
  Process(atm *Atm)
}


type InsertState struct{}
func (is *InsertState)Process(atm *Atm){
  fmt.Println("Insert Your Card Please") // account id
  var account_id string
  fmt.Scan(&account_id)
  _, err := atm.BankRepo.GetBalance(account_id)
 if(err != nil){
   fmt.Println("Remove your card, Your account didn't Find.\n. Please insert valid card.")
   return ;
 }
  atm.CurrentAccount = account_id
  atm.AtmState = &ValidateState{}
}

type ValidateState struct{}
func (is *ValidateState)Process(atm *Atm){
  fmt.Println("Enter Your Pin Please")
  var pin string;
  fmt.Scan(&pin)
  fmt.Println(pin)
  fmt.Println(atm.CurrentAccount)
  err := atm.BankRepo.ValidateAtm(atm.CurrentAccount , pin)
  if(err != nil){
    fmt.Println("Remove your card, Your pin didn't match.")
    atm.AtmState = &InsertState{}
    return; 
  }
  atm.AtmState = &DisplayState{}
  
}
type DisplayState struct{}
func (is *DisplayState)Process(atm *Atm){
  
  fmt.Println("Select Your Operation")
  fmt.Println("1. Balance Enquiry")
  fmt.Println("2. Cash Deposit")
  fmt.Println("3. Cash Withdrawl")
  fmt.Println("4. Exit")
  var choice int
  fmt.Scan(&choice)
  switch choice {
  case 1:
    atm.AtmState = &BalanceQuery{}
  case 2:
    atm.AtmState = &CashDeposit{}
  case 3:
    atm.AtmState = &CashWithdrawl{}
  case 4:
    fmt.Println("Thank You for using our ATM")
    atm.AtmState = &InsertState{}
    return
  default:
    fmt.Println("Invalid Choice")
    atm.AtmState = &InsertState{}
    return
  }
  
}
type BalanceQuery struct{}
func (is *BalanceQuery)Process(atm *Atm){
  amoumt , err := atm.BankRepo.GetBalance(atm.CurrentAccount)
  if(err != nil){
    fmt.Println(err)
    atm.AtmState = &InsertState{}
    return;
  }
  fmt.Println("Your Balance is : ", amoumt)  
  atm.AtmState = &InsertState{}
}
type CashDeposit struct{}
func (is *CashDeposit)Process(atm *Atm){
  
  fmt.Println("Enter the amount you want to deposit")
  var amount int
  fmt.Scan(&amount)
  err := atm.BankRepo.Update(atm.CurrentAccount, int32(amount))
  if(err != nil){
    fmt.Println("Deposit Failed")
    atm.AtmState = &InsertState{}
    return
  }
  fmt.Println("Cash Deposit successful")
  atm.AtmState = &InsertState{}
}
type CashWithdrawl struct{}
func (is *CashWithdrawl)Process(atm *Atm){
  fmt.Println("Enter the amount you want to withdraw")
  var amount int
  fmt.Scan(&amount)
  err := atm.BankRepo.Update(atm.CurrentAccount, -int32(amount))
  if(err != nil){
    fmt.Println("Withdraw Failed")
    atm.AtmState = &InsertState{}
    return
  }
  fmt.Println("Cash Withdrawl successful")
  atm.AtmState = &DispenseState{}
}
type DispenseState struct{}
func (is *DispenseState)Process(atm *Atm){
  fmt.Println("Please collect your cash")
  atm.AtmState = &InsertState{}
}





