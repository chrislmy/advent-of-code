package adventofcode.day4;

import java.util.HashMap;
import java.util.Map;

public class Passport {
  String byr;
  String iyr;
  String eyr;
  String hgt;
  String hcl;
  String ecl;
  String pid;
  String cid;

  public Passport(HashMap<String, String> entries) {
    processDetails(entries);
  }

  private void processDetails(HashMap<String, String> entries) {
    for (Map.Entry<String, String> entry : entries.entrySet()) {
      processField(entry.getKey(), entry.getValue());
    }
  }

  private void processField(String fieldName, String value) {
    switch (fieldName) {
      case "byr":
        byr = value;
        break;
      case "iyr":
        iyr = value;
        break;
      case "eyr":
        eyr = value;
        break;
      case "hgt":
        hgt = value;
        break;
      case "hcl":
        hcl = value;
        break;
      case "ecl":
        ecl = value;
        break;
      case "pid":
        pid = value;
        break;
      case "cid":
        cid = value;
        break;
    }
  }

  @Override
  public String toString() {
    return "Passport{" +
            "byr='" + byr + '\'' +
            ", iyr='" + iyr + '\'' +
            ", eyr='" + eyr + '\'' +
            ", hgt='" + hgt + '\'' +
            ", hcl='" + hcl + '\'' +
            ", ecl='" + ecl + '\'' +
            ", pid='" + pid + '\'' +
            ", cid='" + cid + '\'' +
            '}';
  }
}
