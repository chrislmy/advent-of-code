package adventofcode.day2;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

// Day 2 of Advent of Code challenge
// Link: https://adventofcode.com/2020/day/2
public class PasswordPhilosophy {

  public static int solution(int part) {
    List<String> lines = parseInput();
    int result = 0;

    for (String line : lines) {
      PasswordPolicy passwordPolicy = getPasswordPolicy(line);
      boolean isValidPassword = part == 1
              ? isValidPasswordPart1(passwordPolicy) : isValidPasswordPart2(passwordPolicy);
      result += isValidPassword ? 1 : 0;
    }

    return result;
  }

  private static boolean isValidPasswordPart1(PasswordPolicy passwordPolicy) {
    int count = 0;

    for (char c : passwordPolicy.password.toCharArray()) {
      if (c == passwordPolicy.value) count++;
    }

    return count <= passwordPolicy.maximum && count >= passwordPolicy.minimum;
  }

  private static boolean isValidPasswordPart2(PasswordPolicy passwordPolicy) {
    boolean result = false;
    if (passwordPolicy.password.charAt(passwordPolicy.minimum - 1) == passwordPolicy.value) result = true;
    if (passwordPolicy.password.charAt(passwordPolicy.maximum - 1) == passwordPolicy.value) result = !result;
    return result;
  }

  private static PasswordPolicy getPasswordPolicy(String input) {
    String[] segments = input.split(":");
    String[] passwordPolicySegment = segments[0].split(" ");
    String[] minMax = passwordPolicySegment[0].split("-");

    return new PasswordPolicy(passwordPolicySegment[1].charAt(0),
            Integer.parseInt(minMax[0]),Integer.parseInt(minMax[1]), segments[1].trim());
  }

  private static List<String> parseInput() {
    List<String> result = new ArrayList<>();

    try {
      FileReader fileReader = new FileReader("java/src/main/resources/inputs/day2.txt");
      BufferedReader bufferedReader = new BufferedReader(fileReader);
      String line = bufferedReader.readLine();
      while (line != null) {
        result.add(line);
        line = bufferedReader.readLine();
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

    return result;
  }

  public static void main(String[] args) {
    System.out.println("Part 1: Number of valid passwords: " + solution(1));
    System.out.println("Part 2: Number of valid passwords: " + solution(2));
  }
}
