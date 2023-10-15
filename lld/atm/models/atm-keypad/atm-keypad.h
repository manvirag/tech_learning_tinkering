#include<vector>
#include<string>

 class AtmKeypad
    {
        private:
            std::vector<std::string> keys{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"};

        public:
           void ShowKeyOnScreen();
           std::string GetSelectedKey();

    };
