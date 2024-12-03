import java.nio.file.Files;
import java.nio.file.Paths;
import java.io.IOException;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Main {
    public static void main(String[] args) {
        task1();
        task2();
    }

    public static void task1() {
        String filePath = "../input.txt";
        try {
            String content = new String(Files.readAllBytes(Paths.get(filePath)));
            Pattern pattern = Pattern.compile("mul\\((\\d+),(\\d+)\\)");
            Matcher matcher = pattern.matcher(content);
            int totalSum = 0;
            while (matcher.find()) {
                int num1 = Integer.parseInt(matcher.group(1));
                int num2 = Integer.parseInt(matcher.group(2));
                totalSum += num1 * num2;
            }
            System.out.println("Task 1 | Total sum: " + totalSum);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static void task2() {
        String filePath = "../input.txt";
        try {
            String content = new String(Files.readAllBytes(Paths.get(filePath)));
            Pattern pattern = Pattern.compile("mul\\((\\d+),(\\d+)\\)|don't\\(\\)|do\\(\\)");
            Matcher matcher = pattern.matcher(content);
            int totalSum = 0;
            boolean isDisabled = false;
            while (matcher.find()) {
                String match = matcher.group();
                if (match.equals("don't()")) {
                    isDisabled = true;
                } else if (match.equals("do()")) {
                    isDisabled = false;
                } else if (!isDisabled) {
                    int num1 = Integer.parseInt(matcher.group(1));
                    int num2 = Integer.parseInt(matcher.group(2));
                    totalSum += num1 * num2;
                }
            }
            System.out.println("Task 2 | Total sum: " + totalSum);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}