package adventofcode.day4;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;

public class PassportProcessing {

  private List<Passport> passports;
  private static final String[] requiredFields = new String[]{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"};

  public int solution() {
    List<HashMap<String, String>> passportMaps = parseInput();
    int result = 0;

    for(HashMap<String, String> passportMap : passportMaps) {
      HashSet<String> requiredFieldsSet = new HashSet<>();
      for (String requiredField : requiredFields) requiredFieldsSet.add(requiredField);
      for (Map.Entry<String, String> entry : passportMap.entrySet()) {
        requiredFieldsSet.remove(entry.getKey());
      }
      result += requiredFieldsSet.size() == 0 ? 1 : 0;
    }

    return result;
  }

  private List<HashMap<String, String>> parseInput() {
    passports = new ArrayList<>();
    List<HashMap<String, String>> passportMap = new ArrayList<>();

    try {
      FileReader fileReader = new FileReader("java/src/main/resources/inputs/day4.txt");
      BufferedReader bufferedReader = new BufferedReader(fileReader);
      String line = bufferedReader.readLine();
      HashMap<String, String> fieldMap = new HashMap<>();

      while (line != null) {
        if (line.length() == 0) {
          line = bufferedReader.readLine();
          passports.add(new Passport(fieldMap));
          passportMap.add(fieldMap);
          fieldMap = new HashMap<>();
          continue;
        }

        String[] fields = line.split(" ");

        for (String field : fields) {
          String[] fieldEntry = field.split(":");
          String fieldName = fieldEntry[0];
          String fieldValue = fieldEntry[1];

          fieldMap.put(fieldName, fieldValue);
        }

        line = bufferedReader.readLine();
      }

      passportMap.add(fieldMap);
      passports.add(new Passport(fieldMap));
    } catch (IOException e) {
      e.printStackTrace();
    }

    return passportMap;
  }

  public static void main(String[] args) {
    System.out.println("Part 1: Number of valid passports: " + new PassportProcessing().solution());
  }
}
