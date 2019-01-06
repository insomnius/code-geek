import java.util.Scanner;
import java.util.Arrays;

/**
 * Bubble sort, sometimes referred to as sinking sort, is a simple sorting algorithm that repeatedly
 * steps through the list, compares adjacent pairs and swaps them if they are in the wrong order.
 * The pass through the list is repeated until the list is sorted. The algorithm, which is a comparison
 * sort, is named for the way smaller or larger elements "bubble" to the top of the list. Although the 
 * algorithm is simple, it is too slow and impractical for most problems even when compared to insertion sort.
 * Bubble sort can be practical if the input is mostly sorted order with some out-of-order.
 * https://en.wikipedia.org/wiki/Bubble_sort
 * Level: Easy
 */
class BubbleSort {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        try {
            // Input
            System.out.print("Input jumlah data: ");
            int jumlahData = scanner.nextInt();
            System.out.print("Input item data (setiap item dipisahkan dengan spasi): ");
            int[] itemData = new int[jumlahData];
            for (int a = 0; a < jumlahData; a++) {
                itemData[a] = Integer.parseInt(scanner.next());
            }

            // Process
            for (int a = 1; a < jumlahData; a++) {
                for (int b = 0; b < jumlahData - a; b++) {
                    int itemA = itemData[b];
                    int itemB = itemData[b + 1];
                    if (itemA > itemB) {
                        int temp = itemA;
                        itemA = itemB;
                        itemB = temp;
                        itemData[b] = itemA;
                        itemData[b + 1] = itemB;
                    }
                }
            }

            // Output
            System.out.println("output: " + Arrays.toString(itemData));
        } catch (Exception e) {
            e.printStackTrace();
            System.out.println("Invalid data");
        }
    }
}
