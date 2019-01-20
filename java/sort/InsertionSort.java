import java.util.Scanner;
import java.util.Arrays;

/**
 * Insertion sort is a simple sorting algorithm that builds
 * the final sorted array (or list) one item at time. It is
 * much less efficient on large lists than more advanced 
 * algorithms such as quicksort, heapsort, or merge sort.
 * https://en.wikipedia.org/wiki/Insertion_sort
 * Level: Medium
 */
class InsertionSort {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        try {
            // Input
            System.out.print("Enter the amount of data: ");
            int amountData = scanner.nextInt();
            int[] items = new int[amountData];
            System.out.print("Input data (separated by spaces): ");
            for (int a = 0; a < amountData; a++) {
                items[a] = Integer.parseInt(scanner.next());
            }

            // Process
            for (int a = 1; a < amountData; a++) {
                for (int b = a; b > 0; b--) {
                    int itemA = items[b - 1];
                    int itemB = items[b];
                    if (itemA > itemB) {
                        int temp = itemA;
                        itemA = itemB;
                        itemB = temp;
                        items[b - 1] = itemA;
                        items[b] = itemB;
                    } else {
                        break;
                    }
                }
            }

            // Output
            System.out.println("Output: " + Arrays.toString(items));
        } catch (Exception e) {
            e.printStackTrace();
            System.out.println("Invalid data");
        }
    }
}
