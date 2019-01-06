import java.util.Scanner;
import java.util.Arrays;

/**
 * In computer science, selection sort is a sorting algorithm,
 * specially an in-place-comparison sort. It has 0(n2) time complexity
 * making it inefficient on large lists, and generally performs worse
 * than the similar insertion sort. Selection sort is noted for its
 * simplicity, and it has performance advantages over more complicated
 * algorithms in certain situations, particularly where auxiliary memory
 * is limited.
 *
 * The algorithm divides the input list two parts: the sublist of items
 * already sorted, which is built up from left to right at the front (left)
 * of the list, and the sublist of items remaining to be sorted that occupy
 * the rest of the list. Initially, the sorted sublist is empty and the
 * unsorted sublist is the entire input list. The algorithm proceeds by
 * finding the smallest (or largest, depending on sorting order) element
 * in the unsorted sublist, exchanging (swapping) it with the leftmost
 * unsorted element (putting it in sorted order), and moving the sublist
 * boundaries one element to the right.
 * https://en.wikipedia.org/wiki/Selection_sort
 * Level: Easy
 */
class SelectionSort {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        try {
            // Input
            System.out.print("Input jumlah data: ");
            int jumlahData = scanner.nextInt();
            int[] itemData = new int[jumlahData];
            System.out.print("Input item data (setiap item dipisahkan dengan spasi): ");
            for (int a = 0; a < jumlahData; a++) {
                itemData[a] = Integer.parseInt(scanner.next());
            }

            // Process
            for (int a = 0; a < jumlahData - 1; a++) {
                int itemA = itemData[a];
                int indexLowest = a + 1;
                int lowest = itemData[indexLowest];
                for (int b = a + 1; b < jumlahData; b++) {
                    int itemB = itemData[b];
                    if (lowest > itemB) {
                        lowest = itemB;
                        indexLowest = b;
                    }
                }
                if (lowest < itemA) {
                    itemData[a] = lowest;
                    itemData[indexLowest] = itemA;
                }
            }

            // Output
            System.out.println("Output: " + Arrays.toString(itemData));
        } catch (Exception e) {
            e.printStackTrace();
            System.out.println("Invalid data");
        }
    }
}
