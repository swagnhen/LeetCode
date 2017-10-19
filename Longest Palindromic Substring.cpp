// Longest Palindromic Substring.cpp : �������̨Ӧ�ó������ڵ㡣
//

#include "stdafx.h"
#include <iostream>
#include <string>
using namespace std;

class Solution
{
  public:
	string longestPalindrome(string s)
	{
		int start1 = 0, start2 = 0;
		int max1 = 0, max2 = 0, cum1 = 0, cum2 = 0;
		int r = 0;
		string result;
		for (int i = 0; i < s.size(); i++)
		{
			cum1 = 0;
			r = (i < (s.size() / 2)) ? i : (s.size() - i - 1);
			for (int j = 1; j <= r; j++)
			{
				if (s[i - j] != s[i + j])
					break;
				cum1++;
			}
			if (cum1 > max1)
			{
				start1 = i - cum1;
				max1 = cum1;
			}
			cum2 = 0;
			r = (i < (s.size() / 2)) ? i : (s.size() - i - 2);
			for (int j = 0; j <= r; j++)
			{
				if (s[i - j] != s[i + j + 1])
					break;
				cum2++;
			}
			if (cum2 > max2)
			{
				start2 = i - cum2 + 1;
				max2 = cum2;
			}
		}
		if (max1 >= max2)
			result.assign(s, start1, 2 * max1 + 1);
		else
			result.assign(s, start2, 2 * max2);
		return result;
	}
};

string stringToString(string input)
{
	return input.substr(1, input.length() - 1);
}

int main()
{
	string line;
	while (getline(cin, line))
	{
		string s = stringToString(line);

		string ret = Solution().longestPalindrome(s);

		string out = (ret);
		cout << out << endl;
	}
	return 0;
}
