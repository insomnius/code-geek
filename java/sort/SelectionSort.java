import java.util.Scanner;
import java.util.Arrays;

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
