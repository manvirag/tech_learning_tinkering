#include<vector>
#include<string>

class AtmScreen
{
    private:
      
    public: 
      void ShowOptionsToUser(std::vector<std::string> options);
      std::string GetTheUserSelectedOption(std::vector<std::string> options);
};