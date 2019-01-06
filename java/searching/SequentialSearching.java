import java.util.Scanner;

/**
 * In computer science, linear search or sequential search is a method
 * for finding an element within a list. It sequentially checks each
 * element of the list until a match is found or the whole list has
 * been searched.
 *
 * Linear search runs in at worst linear time and makes at most
 * n comparisons, where n is the length of the list. If each element
 * is equally likely to be searched, then linear search has an average
 * case of n/2 comparisons, but the average case can be affected if 
 * the search probabilities for each element vary. Linear search is
 * rarely practical because other search algorithms and schemes,
 * such as the binary search algorithm and hash tables, allow
 * significantly faster searching for all but short lists.
 * https://en.wikipedia.org/wiki/linear_search
 * Level: Easy
 */
class SequentialSearching {
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
            System.out.print("Input item data yang dicari: ");
            int dataCari = scanner.nextInt();

            // Process
            boolean isFound = false;
            for (int item : itemData) {
                if (item == dataCari) {
                    System.out.println("Output: data yang dicari ditemukan");
                    isFound = true;
                    break;
                }
            }
            if (!isFound) {
                System.out.println("Output: data yang dicari tidak ditemukan");
            }
        } catch (Exception e) {
            e.printStackTrace();
            System.out.println("Invalid data");
        }
    }
}
