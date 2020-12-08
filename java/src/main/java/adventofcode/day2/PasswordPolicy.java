package adventofcode.day2;

public class PasswordPolicy {

  char value;
  int minimum;
  int maximum;
  String password;

  public PasswordPolicy(char value, int minimum, int maximum, String password) {
    this.value = value;
    this.minimum = minimum;
    this.maximum = maximum;
    this.password = password;
  }

  @Override
  public String toString() {
    return String.format("Value: %c, min: %d, max: %d, password: %s", value, minimum, maximum, password);
  }
}
