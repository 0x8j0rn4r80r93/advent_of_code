using System;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

class Program
{
    static void Main()
    {
        string filePath = @"..\..\..\..\input.txt"; 
        string input = File.ReadAllText(filePath);

        int sumTask1 = Solve(input, @"\d"); // Solving Task 1
        int sumTask2 = Solve(input, @"\d|one|two|three|four|five|six|seven|eight|nine"); // Solving Task 2

        Console.WriteLine("Total sum of calibration values for Task 1: " + sumTask1);
        Console.WriteLine("Total sum of calibration values for Task 2: " + sumTask2);
    }

    static int Solve(string input, string pattern)
    {
        return input
            .Split('\n')
            .Select(line => new
            {
                First = Regex.Match(line, pattern),
                Last = Regex.Match(line, pattern, RegexOptions.RightToLeft)
            })
            .Where(match => match.First.Success && match.Last.Success)
            .Select(match => ParseMatch(match.First.Value) * 10 + ParseMatch(match.Last.Value))
            .Sum();
    }

    static int ParseMatch(string match)
    {
        return match switch
        {
            "one" => 1,
            "two" => 2,
            "three" => 3,
            "four" => 4,
            "five" => 5,
            "six" => 6,
            "seven" => 7,
            "eight" => 8,
            "nine" => 9,
            _ => int.Parse(match)
        };
    }
}
