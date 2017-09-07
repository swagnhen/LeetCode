// Median of Two Sorted Arrays.cpp : 定义控制台应用程序的入口点。
//

#include "stdafx.h"
#include <iostream>
#include <vector>
#include <algorithm>
#include <sstream>
using namespace std;

class Solution {
public:
	double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
		int num_size = nums1.size() + nums2.size();
		int j = 0 , k = 0;
		double target1, target2;
		if (num_size % 2 == 1) {
			for (int i = 0; i <= num_size / 2; i++) {
				if (j >= nums1.size()) {
					target1 = nums2[k];
					k++;
				}
				else if (k >= nums2.size()) {
					target1 = nums1[j];
					j++;
				}
				else {
					if (nums1[j] <= nums2[k]) {
						target1 = nums1[j];
						j++;
					}
					else {
						target1 = nums2[k];
						k++;
					}
				}
			}
			return target1;
		}
		else {
			for (int i = 0; i < num_size / 2; i++) {
				if (j >= nums1.size()) {
					target1 = nums2[k];
					k++;
				}
				else if (k >= nums2.size()) {
					target1 = nums1[j];
					j++;
				}
				else {
					if (nums1[j] <= nums2[k]) {
						target1 = nums1[j];
						j++;
					}
					else {
						target1 = nums2[k];
						k++;
					}
				}
			}
			if (j >= nums1.size()) {
				target2 = nums2[k];
				k++;
			}
			else if (k >= nums2.size()) {
				target2 = nums1[j];
				j++;
			}
			else {
				if (nums1[j] <= nums2[k]) {
					target2 = nums1[j];
					j++;
				}
				else {
					target2 = nums2[k];
					k++;
				}
			}
			return (target1 + target2) / 2;
		}
	}
};

void trimLeftTrailingSpaces(string &input) {
	input.erase(input.begin(), find_if(input.begin(), input.end(), [](int ch) {
		return !isspace(ch);
	}));
}

void trimRightTrailingSpaces(string &input) {
	input.erase(find_if(input.rbegin(), input.rend(), [](int ch) {
		return !isspace(ch);
	}).base(), input.end());
}

vector<int> stringToIntegerVector(string input) {
	vector<int> output;
	trimLeftTrailingSpaces(input);
	trimRightTrailingSpaces(input);
	input = input.substr(1, input.length() - 2);
	stringstream ss;
	ss.str(input);
	string item;
	char delim = ',';
	while (getline(ss, item, delim)) {
		output.push_back(stoi(item));
	}
	return output;
}

int main() {
	string line;
	while (getline(cin, line)) {
		vector<int> nums1 = stringToIntegerVector(line);
		getline(cin, line);
		vector<int> nums2 = stringToIntegerVector(line);

		double ret = Solution().findMedianSortedArrays(nums1, nums2);

		string out = to_string(ret);
		cout << out << endl;
	}
	return 0;
}
