import java.util.Scanner;

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
