// ZigZag Conversion.cpp : 定义控制台应用程序的入口点。
//

#include "stdafx.h"
#include <iostream>
#include <string>
using namespace std;


class Solution {
public:
	string convert(string s, int numRows) {
		if (numRows == 1)
			return s;
		string result;
		int num_group = numRows * 2 - 2;
		for (int i = 0; i < numRows; i++) {
			if (i == 0 || i == (numRows - 1)) {
				for (int j = 0; j < s.size() / num_group; j++)
					result += s[i + j * num_group];
				if (i < s.size() % num_group)
					result += s[i + (s.size() / num_group) * num_group];
			}
			else {
				for (int j = 0; j < s.size() / num_group; j++) {
					result += s[i + j * num_group];
					result += s[(j + 1) * num_group - i];
				}
				if (i < s.size() % num_group)
					result += s[i + (s.size() / num_group) * num_group];
				if ((num_group - i) < s.size() % num_group)
					result += s[(s.size() / num_group + 1) * num_group - i];
			}
		}
		return result;
	}
};

string stringToString(string input) {
	return input.substr(1, input.length() - 1);
}

int stringToInteger(string input) {
	return stoi(input);
}

int main() {
	string line;
	while (getline(cin, line)) {
		string s = stringToString(line);
		getline(cin, line);
		int numRows = stringToInteger(line);

		string ret = Solution().convert(s, numRows);

		string out = (ret);
		cout << out << endl;
	}
	return 0;
}

