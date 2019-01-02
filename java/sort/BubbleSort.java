import java.util.Scanner;
import java.util.Arrays;

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
