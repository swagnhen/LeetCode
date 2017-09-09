// String to Integer (atoi).cpp : 定义控制台应用程序的入口点。
//

#include "stdafx.h"
#include <iostream>
#include <string>
using namespace std;

class Solution {
public:
	int myAtoi(string str) {
		__int64 result = 0;
		int flag = 1;
		for (int i = 0; i < str.size(); i++) {
			if (str[i] != '+' && str[i] != '-' && str[i] != 32 && ('0' > str[i] || str[i] > '9'))
				return 0;
			if (str[i] == '+') {
				flag = 1;
				if ('0' > str[i + 1] || str[i + 1] > '9')
					return 0;
			}
			if (str[i] == '-') {
				flag = -1;
				if ('0' > str[i + 1] || str[i + 1] > '9')
					return 0;
			}
			if (str[i] >= '0'&&str[i] <= '9') {
				result *= 10;
				result += str[i] - '0';
				if (flag == 1 && result > 2147483647)
					return 2147483647;
				if (flag == -1 && result > 2147483648)
					return (__int64)(2147483648) * -1;
				if ('0' > str[i + 1] || str[i + 1] > '9')
					break;
			}
		}
		return result * flag;
	}
};

string stringToString(string input) {
	return input.substr(1, input.length() - 1);
}

int main() {
	string line;
	while (getline(cin, line)) {
		string str = stringToString(line);

		int ret = Solution().myAtoi(str);

		string out = to_string(ret);
		cout << out << endl;
	}
	return 0;
}

