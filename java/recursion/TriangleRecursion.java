import java.util.Scanner;

/**
 * Recursion occurs when a thing is defined in terms of itself or of its type. 
 * Recursion is used in variety of disciplines ranging from linguistics to login.
 * The most common application of recursion is in mathematics and computer science,
 * where a function being defined is applied within its own definition. While
 * this apparently defines an infinite number of instances (function values),
 * it is often done in such a way that no loop or infinite chain of references
 * can occur.
 * https://en.wikipedia.org/wiki/Recursion
 * Level: Medium
 */
class TriangleRecursion {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Input number: ");
        try {
            int inputNumber = scanner.nextInt();
            printTriangle(inputNumber, inputNumber);
        } catch (Exception e) {
            e.printStackTrace();
            System.out.println("Invalid data");
        }
    }

    static void printTriangle(int initialValue, int finalValue) {
        if (initialValue < 1) {
            return;
        } else {
            for (int a = initialValue; a <= finalValue; a++) {
                System.out.print("*");
            }
            System.out.println();
            printTriangle(--initialValue, finalValue );
        }
    }
}
