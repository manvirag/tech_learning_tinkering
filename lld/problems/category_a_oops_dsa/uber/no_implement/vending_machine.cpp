/*


Was told to create a vending machine with wierd requirements.
Missed one of the requirement so will probably be rejected.

since we do have less time 

45 mins
controller -> servcie -> dao -> mdoel 

controler -> service -> dao -> nomal one calss inmemorydata(dao), vlaidtion + business logic. 
+ model 

single file .


Design a vending machine. The vending machine should allow a user to:

- Select a product from a list of available products.
- Insert money into the vending machine.
- Receive the selected product if enough money has been inserted.
- Receive any change due after a successful purchase.
- Cancel the transaction and receive a refund of the inserted money.

- The vending machine should also allow an administrator to:

- Add new products to the vending machine.
- Restock existing products.

- Withdraw money from the vending machine. -> ignore 
- Set product prices. 


assumption -> one person at a time.

kind of graph one thing depende on other. 

customer controller. 

- initialiastion
- product selection for buy
- money collect
- valdate
- process buy
- process change
- cancel trasaction. 
- end process

different state of vending machine -> like above. 
start with 

- admin controller. 


for less time 

-> one call -> admin and normal user function both 
-> no validation kind of 
-> no user class etc. 
-> simple interfaces


vendinmachine class 
    -> initial
    -> productselction
    -> collectmoney
    -> validtionmoney
    -> process buy 
    -> end process 
    
methods 
    -> BuyProduct(productId) -> continue on console 
    -> CancelTransaction(id)
    -> .. 

phase 1 
    -> covering happy cases -> make it running 
pahse 2
    -> check and fix the edge cases.
*/

#include<iostream> 
#include<map>
#include<thread>
using namespace std; 



// Forward declaration
class VendineMachineController;

class Transaction {
    public: 
        int id ;
        int productId; 
        int quantity;
        int paidAmount; 
    Transaction(){}
    Transaction(int id): id(id), productId(-1), quantity(-1), paidAmount(-1) {}
};
class Product {
    public: 
        int id ; 
        int price;
    Product(){}
    Product(int id  , int price): id(id), price(price) {}
};
class ProductInventory {
    public: 
        Product product ; 
        int currentStock; 
    ProductInventory(){}
    ProductInventory(int stocks, Product product): currentStock(stocks), product(product) {}
};

// State classes - defined before VendineMachineController
class VendineMachineState {
    public: 
        virtual void process(VendineMachineController *vm) = 0;
};  
class InitializeNewTransaction: public VendineMachineState {
    public: 
        void process(VendineMachineController *vm) override;
};
class SelectProduct: public VendineMachineState {
    public: 
        void process(VendineMachineController *vm) override;
};
class CollectMoney: public VendineMachineState {
    public: 
        void process(VendineMachineController *vm) override;
};
class ValidateAndProcess: public VendineMachineState {
    public: 
        void process(VendineMachineController *vm) override;
};
class EndTransaction: public VendineMachineState {
    public: 
        void process(VendineMachineController *vm) override;
};

class VendineMachineController {
    public: 
        VendineMachineState* vmState; 
        map<int,ProductInventory> products; 
        int currentTransactionId; 
        Transaction currentTransaction;
        
        // change -> always possible assume
    VendineMachineController(): currentTransactionId(0) {}


    void addProduct(int id, int stocks, int price) {
        products[id]=ProductInventory(10,Product(id, price));
    }

    void updateStock(int id, int stocks) {
        products[id].currentStock+=stocks;
    }

    void updatePrice(int id, int price) {
        products[id].product.price=price;
    }

    void startShopping() {
        this->vmState = new InitializeNewTransaction();
        thread t1([&]() {
            this->vmState->process(this);
        });

        
        t1.join();
    }

};

// State class implementations
void InitializeNewTransaction::process(VendineMachineController *vm) {
    vm->currentTransaction = Transaction(vm->currentTransactionId+1); 
    vm->vmState = new SelectProduct();
    vm->vmState->process(vm);
}

void SelectProduct::process(VendineMachineController *vm) {
    cout<<"Selection Product and quantities"<<endl;
    int productId, quantity; 
    cin>>productId>>quantity;
    vm->currentTransaction.productId = productId;
    vm->currentTransaction.quantity = quantity;
    vm->vmState = new CollectMoney();
    vm->vmState->process(vm);
}

void CollectMoney::process(VendineMachineController *vm) {
    cout<<"Please Pay the price for  product "<<endl;
    int paidAmount; 
    cin>>paidAmount;
    vm->currentTransaction.paidAmount = paidAmount;
    vm->vmState = new ValidateAndProcess();
    vm->vmState->process(vm);
}

void ValidateAndProcess::process(VendineMachineController *vm) {
    Transaction currentTx = vm -> currentTransaction; 
    if(vm->products.find(currentTx.productId) == vm->products.end()) {
        cout<<"Product doesn't exist any more"<<endl;
        vm->vmState = new EndTransaction();
        vm->vmState->process(vm);
    } else if (vm->products[currentTx.productId].currentStock < currentTx.quantity) {
        cout<<"Selected quantity is more than existing "<<vm->products[currentTx.productId].currentStock<<endl;
        vm->vmState = new EndTransaction();
        vm->vmState->process(vm);
    } else {
        int totalPrice = vm->products[currentTx.productId].product.price * currentTx.quantity;
        if(currentTx.paidAmount < totalPrice) {
            cout<<"Not enough money to buy this product. Required: "<<totalPrice<<endl;
            vm->vmState = new EndTransaction();
            vm->vmState->process(vm);
        } else {
            vm->products[currentTx.productId].currentStock -= currentTx.quantity;
            if(vm->products[currentTx.productId].currentStock == 0) {
                vm->products.erase(currentTx.productId);
            }
            int changemoney = currentTx.paidAmount - totalPrice;
            cout<<"Please collect your product and change money "<<changemoney<<endl; 
            vm->vmState = new EndTransaction();
            vm->vmState->process(vm);
        }
    }
}

void EndTransaction::process(VendineMachineController *vm) {
    vm->currentTransactionId = vm->currentTransaction.id;
    vm->currentTransaction = Transaction(-1);
    cout<<"Ending Transaction... \n Thanks for visiting our vending machine"<<endl;
}

int main() {
    VendineMachineController *vm = new VendineMachineController();
    vm -> addProduct(1, 10, 5);
    vm -> addProduct(2, 1, 7);
    vm -> addProduct(3, 2, 5);
    vm->updateStock(3,1);
    vm->startShopping();
    return 0; 
}