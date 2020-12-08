package adventofcode.day3;

import javafx.util.Pair;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

// Day 3 of Advent of Code challenge
// Link: https://adventofcode.com/2020/day/3
public class TobogganTrajectory {

  private static final char TREE = '#';

  // Used in part 2
  private static final List<Pair<Integer, Integer>> slopes = new ArrayList<Pair<Integer, Integer>>() {{
    add(new Pair<>(1, 1));
    add(new Pair<>(3, 1));
    add(new Pair<>(5, 1));
    add(new Pair<>(7, 1));
    add(new Pair<>(1, 2));
  }};

  public static int solution(int part, Pair<Integer, Integer> slope) {
    List<List<Character>> grid = parseInput();
    int row = 0;
    int col = 0;
    int result = 0;

    while (row < grid.size()) {
      if (grid.get(row).get(col) == TREE) result++;
      int nextCol = part == 1 ? col + 3 : col + slope.getKey();
      row = part == 1 ? row + 1 : row + slope.getValue();
      col = nextCol % grid.get(0).size();
    }

    return result;
  }

  private static List<List<Character>> parseInput() {
    List<List<Character>> grid = new ArrayList<>();

    try {
      FileReader fileReader = new FileReader("java/src/main/resources/inputs/day3.txt");
      BufferedReader bufferedReader = new BufferedReader(fileReader);
      String line = bufferedReader.readLine();

      while (line != null) {
        List<Character> row = new ArrayList<>();
        for (char c : line.toCharArray()) row.add(c);
        grid.add(row);
        line = bufferedReader.readLine();
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

    return grid;
  }

  public static void main(String[] args) {
    System.out.println("Part 1: Number of trees: " + solution(1, null));
    int part2Result = 1;
    for (Pair<Integer, Integer> slope : slopes) {
      part2Result *= solution(2, slope);
    }
    System.out.println("Part 2: Number of trees: " + part2Result);
  }
}
