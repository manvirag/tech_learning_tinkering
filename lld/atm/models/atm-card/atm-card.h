#include<string>
#include<vector>

class AtmCard{
    private:
      std::string userName;
      std::string cardNumber;
      std::string bank;

    public:
      AtmCard(std::string userName, std::string cardNumber, std::string bank);
      std::string GetCardNumber(); 
      std::string GetCardBank(); 
      std::string GetCardHolderName(); 
};