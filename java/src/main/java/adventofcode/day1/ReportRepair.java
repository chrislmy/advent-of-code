package adventofcode.day1;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashSet;
import java.util.List;

// Day 1 of Advent of Code challenge
// Link: https://adventofcode.com/2020/day/1
public class ReportRepair {

  private static final int TARGET = 2020;

  // Time complexity: O(N)
  // Space complexity: O(N)
  public static int[] solutionPart1() {
    List<Integer> input = parseInput();
    HashSet<Integer> set = new HashSet<>();

    for (int num : input) {
      int remainder = TARGET - num;
      if (set.contains(remainder)) {
        return new int[]{remainder, num};
      }
      set.add(num);
    }

    // Can't find pairs that add to 2020
    return new int[]{-1, -1};
  }

  // Time Complexity: O(N)
  // Space Complexity: O(1), assuming merge sorting in place
  public static int[] solutionPart2() {
    List<Integer> input = parseInput();
    Collections.sort(input);

    for(int i  = 0; i < input.size() - 2; i++) {
      int start = i + 1;
      int end = input.size() - 1;

      while (start < end) {
        int result = input.get(i) + input.get(start) + input.get(end);
        if (result == TARGET) {
          return new int[]{input.get(i), input.get(start), input.get(end)};
        } else if (result < TARGET) {
          start++;
        } else {
          end--;
        }
      }
    }

    // Can't find triplet
    return new int[]{-1, -1, -1};
  }

  private static List<Integer> parseInput() {
    List<Integer> result = new ArrayList<>();

    try {
      FileReader fileReader = new FileReader("java/src/main/resources/inputs/day1.txt");
      BufferedReader bufferedReader = new BufferedReader(fileReader);
      String line = bufferedReader.readLine();
      while (line != null) {
        result.add(Integer.parseInt(line));
        line = bufferedReader.readLine();
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

    return result;
  }

  public static void main(String[] args) {
    int[] pair = solutionPart1();
    int[] triplet = solutionPart2();
    int result1 = pair[0] * pair[1];
    int result2 = triplet[0] * triplet[1] * triplet[2];
    System.out.println(String.format("Pair: [%d, %d], Result: %d", pair[0], pair[1], result1));
    System.out.println(String.format("Triplet: [%d, %d, %d], Result: %d", triplet[0], triplet[1], triplet[2], result2));
  }
}
