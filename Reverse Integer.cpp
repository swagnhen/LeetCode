// Reverse Integer.cpp : �������̨Ӧ�ó������ڵ㡣
//

#include "stdafx.h"
#include <iostream>
#include <string>
using namespace std;

class Solution
{
  public:
	int getpos(__int64 x, __int64 pos)
	{
		__int64 _pos = 1;
		for (int i = 0; i < pos; i++)
			_pos *= 10;
		return abs(x) % _pos * 10 / _pos;
	}

	int reverse(int x)
	{
		int _ref, temp;
		int ref = 1000000000;
		__int64 result = 0;
		for (int i = 10; i > 0; i--)
		{
			if (abs(x) >= ref)
			{
				_ref = 1;
				result = 0;
				for (int j = 1; j <= i; j++)
				{
					result *= 10;
					result += getpos(x, j);
				}
				if (abs(result) >= 2147483648)
					return 0;
				if (x >= 0)
					return result;
				else
					return result * -1;
			}
			ref /= 10;
		}
		return 0;
	}
};

int stringToInteger(string input)
{
	return stoi(input);
}

int main()
{
	string line;
	while (getline(cin, line))
	{
		int x = stringToInteger(line);

		int ret = Solution().reverse(x);

		string out = to_string(ret);
		cout << out << endl;
	}
	return 0;
}
