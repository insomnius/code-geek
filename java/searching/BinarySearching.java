import java.util.Scanner;
import java.util.Arrays;

/**
 * In computer science, binary search, also known as half-interval search,
 * logarithmic search, or binary chop, is a search algorithm that finds
 * the position of a target value within a sorted array. Binary search
 * compares the target value to the middle element of the array. If
 * they are not equal, the half in which the target cannot lie is eliminated
 * and the search continues on the remaining half, again taking the middle 
 * element to compare to the target value, and repeating this until the 
 * target value is found. If the search ends with the remaining half being
 * empty, the target is not in the array. Even though the idea is simple,
 * implementing binary search correctly requires attention to some subtleties
 * about its exit conditions and midpoint calculation, particularly if the
 * values in the array are not all of the whole numbers in the range.
 *
 * Binary search runs in logarithmic time in the worst case, making O(log n)
 * comparisons, where n is the number of elements in the array, the O is
 * Big O notation, and log is the logarithm. Binary search takes constant
 * (O(1)) space, meaning that the space taken by the algorithm is the same
 * for any number of elements in the array. Binary search is faster than
 * linear search except for small arrays, but the array must be sorted first.
 * Although specialized data structures designed for fast searching, such as
 * hash tables, can be searched more efficiently, binary search applies to
 * a wider range of problems.
 *
 * There are numerous variations of binary search. In particular, fractional
 * cascading speeds up binary searches for the same value in multiple arrays.
 * Fractional cascading efficiently solves a number of search problems in 
 * computational geometry and in numerous other fields. Exponential search
 * extends binary search to unbounded lists. The binary search tree and B-tree
 * data structures are based on binary search.
 * https://en.wikipedia.org/wiki/Binary_search_algorithm
 * Level: Medium
 */
class BinarySearching {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        try {
            // Input
            System.out.print("Input jumlah data: ");
            int amountOfData = scanner.nextInt();
            System.out.print("Input item data (setiap item dipisahkan dengan spasi): ");
            int[] itemData = new int[amountOfData];
            for (int a = 0; a < amountOfData; a++) {
                itemData[a] = Integer.parseInt(scanner.next());
            }
            System.out.print("Input item data yang dicari: ");
            int itemSearch = scanner.nextInt();

            // Process
            // Sorting
            for (int a = 0; a < amountOfData - 1; a++) {
                int itemA = itemData[a];
                int indexLowest = a + 1;
                int lowest = itemData[indexLowest];
                for (int b = a + 1; b < amountOfData; b++) {
                    int itemB = itemData[b];
                    if (itemB < indexLowest) {
                        indexLowest = b;
                        lowest = itemB;
                    }
                }
                if (lowest < itemA) {
                    itemData[a] = lowest;
                    itemData[indexLowest] = itemA;
                }
            }

            // Binary searching
            int count = 0;
            boolean isFound = false;
            int initialIndex = 0;
            int finalIndex = amountOfData - 1;
            int tempValue = 0;
            while (!isFound) {
                int middleIndex = (initialIndex + finalIndex) / 2;
                int value = itemData[middleIndex];
                if (value == itemSearch) {
                    isFound = true;
                    break;
                } else if (value < itemSearch) {
                    initialIndex = middleIndex + 1;
                } else if (value > itemSearch) {
                    finalIndex = middleIndex - 1;
                }
                if (value == tempValue) {
                    count += 1;
                } else {
                    tempValue = value;
                    count = 0;
                }
                if (count == 3) {
                    break;
                }   
            }
            if (isFound) {
                System.out.println("Output: data yang dicari ditemukan");
            } else {
                System.out.println("Output: data yang dicari tidak ditemukan");
            }
        } catch (Exception e) {
            e.printStackTrace();
            System.out.println("Invalid data");
        }
    }
}
