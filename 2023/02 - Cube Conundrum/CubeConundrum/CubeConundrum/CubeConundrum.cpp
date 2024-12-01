#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <algorithm>
#include <regex>

using namespace std;

int extractMaxCount(const string& line, const regex& colorPattern) {
    sregex_iterator currentMatch(line.begin(), line.end(), colorPattern);
    sregex_iterator lastMatch;
    int maxCount = 0;
    while (currentMatch != lastMatch) {
        smatch match = *currentMatch;
        maxCount = max(maxCount, stoi(match.str(1)));
        currentMatch++;
    }
    return maxCount;
}

int main() {
    ifstream file("../input.txt");
    string line;
    int sumOfPossibleGameIDs = 0;
    long long sumOfPowers = 0;

    regex gamePattern("Game (\\d+)");
    regex redPattern("(\\d+) red");
    regex greenPattern("(\\d+) green");
    regex bluePattern("(\\d+) blue");

    if (file.is_open()) {
        while (getline(file, line)) {
            smatch gameMatch;
            regex_search(line, gameMatch, gamePattern);
            int gameID = stoi(gameMatch.str(1));

            int maxRed = extractMaxCount(line, redPattern);
            int maxGreen = extractMaxCount(line, greenPattern);
            int maxBlue = extractMaxCount(line, bluePattern);

            if (maxRed <= 12 && maxGreen <= 13 && maxBlue <= 14) {
                sumOfPossibleGameIDs += gameID;
            }

            sumOfPowers += (long long)maxRed * maxGreen * maxBlue;
        }
        file.close();
    }
    else {
        cout << "Unable to open file";
        return 1;
    }

    cout << "Sum of possible game IDs: " << sumOfPossibleGameIDs << endl;
    cout << "Sum of the power of the minimum sets: " << sumOfPowers << endl;
    return 0;
}
