// Longest Substring Without Repeating Characters.cpp : 定义控制台应用程序的入口点。
//

#include "stdafx.h"
#include <iostream>
#include <string>
using namespace std;

class Solution {
public:
	int lengthOfLongestSubstring(string s) {
		int max = 0;
		int cum = 0;
		bool flag[256];
		memset(flag, false, sizeof(bool) * 256);
		int i = 0, j = 0;
		while (i < s.size()) {
			cum = 1;
			flag[s[i]] = true;
			for (j = i + 1; j < s.size(); j++) {
				if (flag[s[j]] == true) {
					i = j - cum + 1;
					break;
				}
				flag[s[j]] = true;
				cum++;
			}
			if (cum > max)
				max = cum;
			if (j >= s.size())
				break;
			memset(flag, false, sizeof(bool) * 256);
		}
		return max;
	}
};

string stringToString(string input) {
	return input.substr(1, input.length() - 1);
}

int main() {
	string line;
	while (getline(cin, line)) {
		string s = stringToString(line);

		int ret = Solution().lengthOfLongestSubstring(s);

		string out = to_string(ret);
		cout << out << endl;
	}
	return 0;
}

